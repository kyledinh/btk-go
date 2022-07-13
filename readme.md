**Table of Contents** 

- [BTK Project](#bin-tool-kit-project)
  - [About the project](#about-the-project)
    - [API docs](#api-docs)
    - [Design](#design)
    - [Status](#status)
    - [See also](#see-also)
  - [Getting started](#getting-started)
    - [Layout](#layout)
  - [Notes](#notes)


# BIN Tool Kit Project

## About the project

A developer's tool kit. 

### API docs

The template doesn't have API docs. For web service, please include API docs here, whether it's
auto-generated or hand-written. For auto-generated API docs, you can also give instructions on the
build process.

### Design

The template follows project convention doc.

* [Repository Conventions](https://github.com/caicloud/engineering/blob/master/guidelines/repo_conventions.md)

### Status

This is starter project.

### See also

* [Golang template project](https://github.com/caicloud/golang-template-project)

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

### Layout
```
.
├── Makefile
├── cmd
│   └── cli
│       └── main.go
├── config
│   └── sample.launch.json
├── dist
│   ├── btk-cli-linux
│   └── btk-cli-macos
├── docs
│   └── resources.md
├── go.mod
├── go.sum
├── pkg
│   └── hash
│       └── hash_test.go
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
* `docs` for project documentations.
* `hack` contains scripts used to manage this repository, e.g. codegen, installation, verification, etc.
* `pkg` places most of project business logic and locate `api` package.
* `release` [chart](https://github.com/caicloud/charts) for production deployment.
* `test` holds all tests (except unit tests), e.g. integration, e2e tests.
* `third_party` for all third party libraries and tools, e.g. swagger ui, protocol buf, etc.
* `vendor` contains all vendored code. Gitignored.

## Notes

* Makefile **MUST NOT** change well-defined command semantics, see Makefile for details.
* Every project **MUST** use `dep` for vendor management and **MUST** checkin `vendor` direcotry.
* `cmd` and `build` **MUST** have the same set of subdirectories for main targets
  * For example, `cmd/admin,cmd/controller` and `build/admin,build/controller`.
  * Dockerfile **MUST** be put under `build` directory even if you have only one Dockerfile.