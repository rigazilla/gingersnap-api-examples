Protoc code generation is embedded in the main.go files as a _go:generate_ tag.

### docuToProto
Takes a YAML document as an input string, turns it into json, turns it into a protobuf object
### protoToDocu
Takes a Golang struct, turns it into a json document
