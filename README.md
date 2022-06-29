# engytita-api-examples
Engytita entities needed for administration and configuration are described as a set of .proto files. User can generate protobuf code to handle configuration entities as document (i.e json files) or add grpc IDL to generate client/server API. 

This is a collection of examples that consume engytita-api .proto definitions in both ways in differents languages. Each example has its own readme with more info.
Engytita API spec are include as git submodule, after cloned you may want to run:
- `git submodule init`
- `git submodule update`

## Golang
### Protobuf document
protoc command is embedded as `go:generate tag` and explained in the .go files.
To run the examples:
- `cd golang/protobuf`
- `go generate ./...`
- `go run protoToDocu/main.go`
- `go run docuToProto/main.go`

### grpc API
protoc command is embedded as `go:generate tag` and explained in the .go files.
- `cd golang/grpc`
- `go generate ./...`
- `go run server/main.go`
- in a different terminal `go run client/main.go`


## Java
### Protobuf document
protoc command is embedded in pom.xml
To run the examples:
- `cd java/protobuf`
- `mvn compile`
- `mvn exec:java -q -Dexec.mainClass=org.engytita.api.AppJsonToProtobuf`
- `mvn exec:java -q -Dexec.mainClass=org.engytita.api.AppProtobufToJson`
  

