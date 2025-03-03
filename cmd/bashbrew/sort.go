package main

import (
	"github.com/docker-library/bashbrew/manifest"
	"pault.ag/go/topsort"
)

func sortRepos(repos []string, applyConstraints bool) ([]string, error) {
	rs := []*Repo{}
	rsMap := map[*Repo]string{}
	for _, repo := range repos {
		r, err := fetch(repo)
		if err != nil {
			return nil, err
		}
		if _, ok := rsMap[r]; ok {
			// if we have a duplicate, let's prefer the first
			continue
		}
		rs = append(rs, r)
		rsMap[r] = repo
	}

	// short circuit if we don't have to go further
	if noSortFlag || len(repos) <= 1 {
		return repos, nil
	}

	rs, err := sortRepoObjects(rs, applyConstraints)
	if err != nil {
		return nil, err
	}

	ret := []string{}
	for _, r := range rs {
		ret = append(ret, rsMap[r])
	}
	return ret, nil
}

func (r Repo) SortedEntries(applyConstraints bool) ([]*manifest.Manifest2822Entry, error) {
	entries := r.Entries()

	// short circuit if we don't have to go further
	if noSortFlag || len(entries) <= 1 {
		return entries, nil
	}

	// create individual "Repo" objects for each entry in "r" so they can be sorted by the same "sortRepoObjects" function
	rs := []*Repo{}
	for i := range entries {
		rs = append(rs, r.EntryRepo(entries[i]))
	}

	rs, err := sortRepoObjects(rs, applyConstraints)
	if err != nil {
		return nil, err
	}

	ret := []*manifest.Manifest2822Entry{}
	for _, entryR := range rs {
		ret = append(ret, entryR.TagEntries...)
	}
	return ret, nil
}

func sortRepoObjects(rs []*Repo, applyConstraints bool) ([]*Repo, error) {
	// short circuit if we don't have to go further
	if noSortFlag || len(rs) <= 1 {
		return rs, nil
	}

	network := topsort.NewNetwork()

	// a map of alternate tag names to the canonical "node name" for topsort purposes
	canonicalNodes := map[string]string{}
	canonicalRepos := map[string]*Repo{}

	for _, r := range rs {
		node := r.Identifier()
		for _, entry := range r.Entries() {
			// add edges both with and without namespace so sorting still works properly for official images ("bashbrew --namespace amd64 list --build-order wordpress php")
			// this should be reasonably harmless for other use cases of --namespace but catches things like "tianon/foo -> tianon/bar" and things like "php -> wordpress" equally even if we're building to target a different namespace
			tags := []string{}
			tags = append(tags, r.Tags("", false, entry)...)
			tags = append(tags, r.Tags(namespace, false, entry)...)
			for _, tag := range tags {
				if canonicalRepo, ok := canonicalRepos[tag]; ok && canonicalRepo.TagName != "" {
					// if we run into a duplicate, we want to prefer a specific tag over a full repo
					continue
				}

				canonicalNodes[tag] = node
				canonicalRepos[tag] = r
			}
		}
		network.AddNode(node, r)
	}

	for _, r := range rs {
		for _, entry := range r.Entries() {
			if applyConstraints && r.SkipConstraints(entry) {
				continue
			}
			/*
			// TODO need archFilter here :(
			if archFilter && !entry.HasArchitecture(arch) {
				continue
			}
			*/

			entryArches := []string{arch}
			if !applyConstraints { // TODO && !archFilter
				entryArches = entry.Architectures
			}

			froms := []string{}
			for _, entryArch := range entryArches {
				archFroms, err := r.ArchDockerFroms(entryArch, entry)
				if err != nil {
					return nil, err
				}
			ArchFroms:
				for _, archFrom := range archFroms {
					for _, from := range froms {
						if from == archFrom {
							continue ArchFroms
						}
					}
					froms = append(froms, archFrom)
				}
			}

			for _, from := range froms {
				fromNode, ok := canonicalNodes[from]
				if !ok {
					// if our FROM isn't in the list of things we're sorting, it isn't relevant in this context
					continue
				}

				// TODO somehow reconcile/avoid "a:a -> b:b, b:b -> a:c" (which will exhibit here as cyclic)
				for _, tag := range r.Tags(namespace, false, entry) {
					if tagNode, ok := canonicalNodes[tag]; ok {
						if tagNode == fromNode {
							// don't be cyclic
							continue
						}
						if err := network.AddEdge(fromNode, tagNode); err != nil {
							return nil, err
						}
					}
				}
			}
		}
	}

	nodes, err := network.Sort()
	if err != nil {
		return nil, err
	}

	ret := []*Repo{}
	for _, node := range nodes {
		ret = append(ret, node.Value.(*Repo))
	}

	return ret, nil
}
