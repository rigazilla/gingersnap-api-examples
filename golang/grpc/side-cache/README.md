## Run the server
cd service
go generate && go build
./service

## Test it with HTTP
curl -X GET http://localhost:5252/api/v1alpha/SFRUUAo=

Values are base64 encoded

## Test it with gRPC
cd client
go generate && go build
./client
