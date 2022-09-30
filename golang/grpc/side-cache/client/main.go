/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	cachepb "github.com/rigazilla/gingersnap-cloud-api-examples/golang/grpc/side-cache/client/gingersnap-cloud-api/service/cache/v1alpha"
	cachepbv1alpha2 "github.com/rigazilla/gingersnap-cloud-api-examples/golang/grpc/side-cache/client/gingersnap-cloud-api/service/cache/v1alpha2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

//go:generate protoc --proto_path=../../../../gingersnap-cloud-api service/cache/v1alpha/cache.proto --go-grpc_out=. --go_out=.
//go:generate protoc --proto_path=../../../../gingersnap-cloud-api --grpc-gateway_out=logtostderr=true:. service/cache/v1alpha/cache.proto

//go:generate protoc --proto_path=../../../../gingersnap-cloud-api service/cache/v1alpha2/cache.proto --go-grpc_out=. --go_out=.
//go:generate protoc --proto_path=../../../../gingersnap-cloud-api --grpc-gateway_out=logtostderr=true:. service/cache/v1alpha2/cache.proto

func main() {
	if os.Getenv("GRPC_SERVER_PORT") == "" {
		e := godotenv.Load() //Load .env file for local environment
		if e != nil {
			fmt.Println(e)
		}
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:"+os.Getenv("GRPC_SERVER_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := cachepb.NewCacheServiceClient(conn)

	cv1alpha2 := cachepbv1alpha2.NewCacheServiceClient(conn)

	if retVal, err := c.Get(context.Background(), &cachepb.Key{Key: []byte{'g', 'R', 'P', 'C'}}); err == nil {
		fmt.Printf(("Result %v\n"), retVal)
	} else {
		fmt.Printf("Error in invocation: %v\n", err)
	}

	if retVal, err := cv1alpha2.Get(context.Background(), &cachepbv1alpha2.Key{Key: []byte{'g', 'R', 'P', 'C'}}); err == nil {
		fmt.Printf(("Result %v\n"), retVal)
	} else {
		fmt.Printf("Error in invocation: %v\n", err)
	}
}
