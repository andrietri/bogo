# bogo
Go CLI generate project.

```
├── cmd/                Initial stage of the application will run.
├── internal/           Core module of the application and contains the implementation of various business logic.
│   ├──  controller/    This module is only to gather input (REST/gRPC/console/etc) and pass input as request to service.
│   │   ├──  http/
│   │   ├──  grpc/
│   ├──  service/       This module contain business logic, this module get input request from controller, this module use repository for things related to persistence.
│   ├──  repository/    This module only for things related to persistence (CRUD database/redis/etc).
├── go.mod
└── go.sum
```

## Installation

You can install by using go binary.

```shell
go install github.com/andrietri/bogo@latest
```

or you can define your prefered version.

```shell
go install github.com/andrietri/bogo@v1.0.0
```

You can check your version by running.

```shell
bogo version
```

## Get Started

1. Follow installation.

2. Create new project.

```shell
bogo create-project foo-project github.com/andrietri/foo-project
```

New project will be created in directory `foo-project`.

3. Create new module.

```shell
cd foo-project && bogo create-module profile
```

New module will be created in directory `internal/controller/http/`.
New module will be created in directory `internal/service/`.
New module will be created in directory `internal/repository/`.

## Docs

```shell
bogo --help
Golang CLI generate project.

Usage:
  bogo [command]

Available Commands:
    create-module  Create new module with the provided module name
    create-project Create new project with the provided module name
    help           Help about any command
    structure      Explain code structure
    version        Show bogo version

Flags:
  -h, --help   help for bogo

Use "bogo [command] --help" for more information about a command.
```