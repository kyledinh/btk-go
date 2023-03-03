**Table of Contents** 

- [BTK Project](#bin-tool-kit-project)
  - [About the project](#about-the-project)
    - [Usage](#usage)
    - [Design](#design)
    - [Status](#status)
    - [See also](#see-also)
  - [Getting started](#getting-started)
    - [Layout](#layout)
  - [Notes](#notes)

<br><hr><br>

# BIN Tool Kit Project

<br><hr><br>
## About the project

A developer's tool kit. 

### Usage 

- Use the [make](#makefile) to build the CLI binary.
- Run `btk -help` to see the options
- Markdown documentation files in `pkg/codex/` can now be served up to stdout with `btk -docs [target file]`
- Or you can run the webserver with `btk -web` and browse at `http://localhost:8001/docs/`, .md files will be served as html

Flags          | Description
---------------|--------------------------------------------------------------------
-d             | Specify a directory to write to instead of ./,  '-d=output'
-docs          | Output a documentation file
-gen           | Generate models '-gen=model i=specs/project.yaml -d=internal/model'
-gentest       | Generate a unit test scaffolding '-gentest -i file.go'
-i             | Specify a spec yaml file  '-i=spec.yaml'
-j2y           | Convert the default json to yaml.
-jsontoyaml    | Convert default json to yaml.
-o             | Specify a file to write to instead of STDOUT,  '-o=filename.ext'
-snip          | Output a snip/snippet
-v             | for version
-web           | to launch http server
-y2j           | Convert yaml to json instead of the default json to yaml.
-yaml2goschema | Convert spec.yaml/your-spec.yaml to go schema.
-yamltojson    | Convert yaml to json instead of the default json to yaml.


<br><br>
### Example Usage:

- Convert a yaml file to json, output to a subdirectory : `btk -y2j -i=some.yaml -d=output_dir`
- List snippets : `btk -snip`
- Export a snippet to your pastebin (MacOS) : `btk -snip openapi.yaml | pbcopy`, (Cmd +v) to paste.
- List docs : `btk -docs`
- Launch docs in webserver : `btk -web`, Open browser to [http://localhost:8001/docs/index.md](http://localhost:8001/docs/index.md)

<br><br>

### Design

> The template follows project convention doc.

* [Repository Conventions](https://github.com/caicloud/engineering/blob/master/guidelines/repo_conventions.md)

### Status

This is starter project.

### See also

* [Golang template project](https://github.com/caicloud/golang-template-project)


<br><hr><br>

## Getting started

Below we describe the conventions or tools specific to golang project.

Start with the Makefile and see most common operations for this repository.
### Makefile
* make check
* make build 
* make test 


### Software Versions

| Software       | Version | Install                                        |
|----------------|---------|------------------------------------------------|
| Go             | 1.18    | https://go.dev/doc/install                     |
| Docker Desktop | 4.3.x   | https://www.docker.com/products/docker-desktop |

<br><br>

### Layout
```
.
├── Makefile
├── cmd
│   ├── cli
│   │   └── main.go
│   └── http-server 
│       └── main.go
├── config
│   └── sample.launch.json
├── dist
│   ├── btk-cli-linux
│   └── btk-cli-macos
├── go.mod
├── go.sum
├── pkg
│   └── codex        // CONTENT directory to store documentation, in .md or text formats.  
│       ├── docs  
│       ├── snippets
│       └── templates 
├── readme.md
├── scripts
│   └── dev
│       ├── check.sh
│       ├── lint.sh
│       └── setup.sh
└── semvar
```

A brief description of the layout:

* `.github` has two template files for creating PR and issue. Please see the files for more details.
* `.gitignore` varies per project, but all projects need to ignore `bin` directory.
* `.golangci.yml` is the golangci-lint config file.
* `Makefile` is used to build the project. **You need to tweak the variables based on your project**.
* `CHANGELOG.md` contains auto-generated changelog information.
* `OWNERS` contains owners of the project.
* `readme.md` is a detailed description of the project.
* `dist` is to hold build outputs. Folder is gitignored.
* `cmd` contains main packages, each subdirecoty of `cmd` is a main package.
* `build` contains scripts, yaml files, dockerfiles, etc, to build and package the project.
* `hack` contains scripts used to manage this repository, e.g. codegen, installation, verification, etc.
* `pkg` places most of project business logic and locate `api` package. SHARED code for cmd executables 
* `release` [chart](https://github.com/caicloud/charts) for production deployment.
* `test` holds all tests (except unit tests), e.g. integration, e2e tests.
* `third_party` for all third party libraries and tools, e.g. swagger ui, protocol buf, etc.
* `vendor` contains all vendored code. Gitignored.


## Notes

* Makefile **MUST** change well-defined command semantics, see Makefile for details.
* Every project **MUST** use `dep` for vendor management and **MUST** checkin `vendor` direcotry.
* `cmd` and `build` **MUST** have the same set of subdirectories for main targets
  * For example, `cmd/admin,cmd/controller` and `build/admin,build/controller`.
  * Dockerfile **MUST** be put under `build` directory even if you have only one Dockerfile.