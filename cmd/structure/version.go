package structure

import "fmt"

func Structure() {
	fmt.Println(structureText)
}

var structureText = `
├── cmd/                Initial stage of the application will run.
├── internal/           Core module of the application and contains the implementation of various business logic.
│   ├──  controller/    This module is only to gather input (REST/gRPC/console/etc) and pass input as request to service.
│   │   ├──  http/
│   │   ├──  grpc/
│   ├──  service/       This module contain business logic, this module get input request from controller, this module use repository for things related to persistence.
│   ├──  repository/    This module only for things related to persistence (CRUD database/redis/etc).
├── go.mod
└── go.sum
`
