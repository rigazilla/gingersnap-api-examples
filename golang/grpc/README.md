This example generate a client/server API starting from the _protos/server.proto_ file. This file defines a set of rpc operations that depend from the Engytita API entities.
grpc and protobuf are generated in different steps and in different golang packages.

Protoc generation is embedded in .go files, two different execution are necessaries to keep separation between entities and operations.

### server
Implements a server with two services _CreateRegion_, _GetRegion_.

### client
An example of client that consumes the services.