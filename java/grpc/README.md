The protoc code generation is embedded in the pom.xml, two execution of the protobuf plugin are needed to keep Engytita API separated from the application coded gRPC. Different output directories are configured since the plugin wants to cleanup the target before generation.

### AppServer
Provides two services for create and get a region.
### AppClient
Example of client that consumes the services.
