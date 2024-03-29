# gingersnap-api-examples
Gingersnap Cloud entities needed for administration and configuration are described as a set of .proto files. User can generate protobuf code to handle configuration entities as document (i.e json files) or add grpc IDL to generate client/server API. 

This is a collection of examples that consume gingersnap-project/api .proto in both ways in different languages. Each example has its own readme with more info. Key step of the process is the _protoc_ generation of the API, this is probably the starting point to integrate the gingersnap API in your application.
Gingersnap Cloud API spec are include as git submodule, after cloned you may want to run:
- `git submodule init`
- `git submodule update`

## Golang
### Protobuf document
protoc command is embedded as `go:generate tag` and explained in the .go files.
To run the examples:
- `cd golang/protobuf/docuToProto` or `cd golang/protobuf/protoToDocu`
- `go generate`
- `go run main.go`

### gRPC API
protoc command is embedded as `go:generate tag` and explained in the .go files.
- `cd golang/grpc/example/server`
- `go generate`
- `go run main.go`

in a different terminal:
- `cd golang/grpc/example/client`
- `go generate`
- `go run main.go`


## Java
### Protobuf document
protoc command is embedded in pom.xml
To run the examples:
- `cd java/protobuf`
- `mvn compile`
- `mvn exec:java -q -Dexec.mainClass=org.gingersnap.api.AppJsonToProtobuf`
- `mvn exec:java -q -Dexec.mainClass=org.gingersnap.api.AppProtobufToJson`
  
### gRPC API
protoc command is embedded in pom.xml
To run the examples:
- `cd java/grpc`
- `mvn compile`
- `mvn exec:java -q -Dexec.mainClass=org.gingersnap.api.AppServer`
- `mvn exec:java -q -Dexec.mainClass=org.gingersnap.api.AppClient`

