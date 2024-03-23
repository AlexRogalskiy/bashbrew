# Bashbrew (`bashbrew`)


[![Jenkins Build Status](https://doi-janky.infosiftr.net/job/bashbrew/job/master/badge/icon?subject=Jenkins)](https://doi-janky.infosiftr.net/job/bashbrew/job/master/) [![GitHub CI](https://github.com/docker-library/bashbrew/workflows/GitHub%20CI/badge.svg?branch=master&event=push)](https://github.com/docker-library/bashbrew/actions?query=workflow%3A%22GitHub+CI%22)


`$ bashbrew --help`

```console
NAME:
   bashbrew - canonical build tool for the official images

USAGE:
   bashbrew [global options] command [command options] [arguments...]

VERSION:
   v0.1.0

COMMANDS:
     list, ls    list repo:tag combinations for a given repo
     build       build (and tag) repo:tag combinations for a given repo
     tag         tag repo:tag into a namespace (especially for pushing)
     push        push namespace/repo:tag (see also "tag")
     put-shared  update shared tags in the registry (and multi-architecture tags)
     help, h     Shows a list of commands or help for one command
   plumbing:
     children, offspring, descendants, progeny  print the repos built FROM a given repo or repo:tag
     parents, ancestors, progenitors            print the repos this repo or repo:tag is FROM
     cat                                        print manifest contents for repo or repo:tag
     from                                       print FROM for repo:tag

GLOBAL OPTIONS:
   --debug                  enable more output (esp. all "docker build" output instead of only output on failure) [$BASHBREW_DEBUG]
   --no-sort                do not apply any sorting, even via --build-order
   --arch value             the current platform architecture (default: "amd64") [$BASHBREW_ARCH]
   --namespace value        a repo namespace to act upon/in [$BASHBREW_NAMESPACE]
   --constraint value       build constraints (see Constraints in Manifest2822Entry) [$BASHBREW_CONSTRAINTS]
   --exclusive-constraints  skip entries which do not have Constraints
   --arch-namespace value   architecture to push namespace mappings for creating indexes/manifest lists ("arch=namespace" ala "s390x=tianons390x") [$BASHBREW_ARCH_NAMESPACES]
   --config value           where default "flags" configuration can be overridden more persistently (default: "/home/tianon/.config/bashbrew") [$BASHBREW_CONFIG]
   --library value          where the bodies are buried (default: "/home/tianon/docker/official-images/library") [$BASHBREW_LIBRARY]
   --cache value            where the git wizardry is stashed (default: "/home/tianon/.cache/bashbrew") [$BASHBREW_CACHE]
   --help, -h, -?           show help
   --version, -v            print the version
```

## Installing

Pre-built binaries are available to [download from Jenkins (for all supported architectures)](https://doi-janky.infosiftr.net/job/bashbrew/job/master/lastSuccessfulBuild/artifact/).

(For building `bashbrew` yourself, it's a pretty standard Go application.)

## Usage

Docker version 1.10 or above is required for use of Bashbrew.

In general, `bashbrew build some-repo` or `bashbrew build ./some-file` should be sufficient for using the tool at a surface level, especially for testing. For more complex usage, please see the built-in help (`bashbrew --help`, `bashbrew build --help`, etc).

## Configuration

The default "flags" configuration is in `~/.config/bashbrew/flags`, but the base path can be overridden with `--config` or `BASHBREW_CONFIG` (technically, the full path to the default `flags` configuration file is `${BASHBREW_CONFIG:-${XDG_CONFIG_HOME:-$HOME/.config}/bashbrew}/flags`).

To set a default namespace for specific commands only:

```
Commands: tag, push
Namespace: officialstaging
```

To set a default namespace for all commands:

```
Namespace: jsmith
```

A more complex example:

# comments are allowed anywhere (and are ignored)
Library: /mnt/docker/official-images/library
Cache: /mnt/docker/bashbrew-cache
Constraints: aufs, docker-1.9, tianon
ExclusiveConstraints: true

# namespace just "tag" and "push" (not "build")
Commands: tag, push
Namespace: tianon

Commands: list
ApplyConstraints: true

Commands: tag
Debug: true
```

In this example, `bashbrew tag` will get both `Namespace` and `Debug` applied (options are additive).

# Web Desing

Definitely, as a writer contributing to user experience and web design, you can follow these steps to improve efficiency and reduce frustration:

- **Understand the purpose**: Before writing, clearly understand the purpose of the website, its audience and the goals you want to achieve.

- **Planning**: Create an outline or content structure before you start writing. This will help you have a clear guide.

- **Research**: Research your topic to make sure your content is accurate and relevant.

- **Concise writing**: Write clearly and concisely. Use simple language and avoid unnecessary jargon.

- **Collaboration**: Work closely with designers and developers to ensure that content is effectively integrated into the design.

- **Continuous review**: Review and adjust content as the design process progresses, rather than waiting until the end.

- **User testing**: Conduct user testing to get valuable feedback and adjust content based on comments.

- **Management tools**: Use project management and collaboration tools to maintain an efficient workflow.

- **Maintain flexibility**: Be willing to adapt to changes in design and strategy as the project evolves.

- **Continuous learning**: Stay current on trends and best practices in web design and user experience to constantly improve your work.

By following these tips, you can move beyond the blank page.

## [Learn Popular Programming Languages](https://avalonix.org/learning)

**Author's Note: Carlos Ibarra**
**2024/03/22**

When we start learning a popular programming language, there is often quite a bit of confusion as we read information, this is because we simply don't know where to start working or what language we should study first, and in reality, "there isn't one". 

From my point of view, when someone tries to teach you a certain language, the target to program is you (client). First, you should question your goals, conclusions, and where you want to go. Here we try to clarify the idea of learning as far as web programming is concerned.

## Let's talk about Python

Python is a high-level, interpreted programming language with a very clear and readable syntax. It has gained a lot of popularity in recent years due to its simplicity and versatility, making it ideal for beginners and experienced developers as well.

### Why Python?

Python is known for its ease of learning and for being less strict on syntax compared to other languages, allowing new programmers to get quick results without a lot of upfront effort. In addition, Python has a large community of developers who are constantly contributing tools and libraries to extend its capabilities.

### Getting Started in Python

To get started with Python, the first thing you need to do is install the Python interpreter on your computer. Once installed, you can start writing your first scripts. It's a good idea to start with basic exercises that will help you understand the syntax and control structures of language.

### Resources for Learning Python

There are a plethora of resources available for learning Python. From free online tutorials to specialized courses and books. Some recommended resources are:

- [The official Python documentation](https://docs.python.org/3/)
- [Codecademy](https://discuss.codecademy.com/t/code-review/790439) 
- [Coursera](https://www.coursera.org/courses?query=python) 
- Books such as "[Automate the Boring Stuff with Python](https://www.amazon.com.mx/Automate-Boring-Stuff-Python-2nd/dp/1593279922)" o "[Python Crash Course](https://archive.org/details/pythoncrashcours0000matt)"

### Projects to Practice
[tools mixplorer](https://drive.google.com/drive/folders/1BfeK39boriHy-9q76eXLLqbCwfV17-Gv)

Once you have a basic understanding of the language, it's important to start working on projects of your own. This will help you consolidate your knowledge and gain hands-on experience. Some ideas for projects can be:

- [ ] A program to automate simple tasks
- [ ] A basic website with Flask or Django
- [ ] A data analysis with Pandas and Matplotlib

### Conclusion

Learning any programming language takes time and practice. There's no "best" language to start with; The important thing is to choose one that aligns with your interests and goals. Python is an excellent choice because of its simplicity and the vast opportunities it offers in different fields such as data science, web development, and automation.
