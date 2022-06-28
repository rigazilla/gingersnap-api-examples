# engytita-api-examples
This is an examples collection of application that consumes engytita-api .proto definitions. A couple of examples are available for each language: a proto->json converter and the other way round.
### Golang
protoc command is embedded as `go:generate tag` and explained in the .go files.
To run the examples:
- `go generate ./...`
- `go run protoToDocu/main.go`
- `go run docuToProto/main.go`
### Java
protoc command is embedded in pom.xml
To run the examples:
- `mvn compile`
- `mvn exec:java -q -Dexec.mainClass=org.engytita.api.AppJsonToProtobuf`
- `mvn exec:java -q -Dexec.mainClass=org.engytita.api.AppProtobufToJson`
  

