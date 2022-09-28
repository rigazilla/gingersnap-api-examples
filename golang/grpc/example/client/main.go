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
	"flag"
	"log"
	"time"

	pb "github.com/rigazilla/engytita-api-examples/golang/grpc/client/gingersnap-cloud-api/config/cache/v1alpha"
	gr "github.com/rigazilla/engytita-api-examples/golang/grpc/client/regionstore/v1alpha"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

//go:generate protoc --proto_path=../protos/ --proto_path=../../../../gingersnap-cloud-api/ --go_out=. --go-grpc_out=. config/cache/v1alpha/region.proto config/cache/v1alpha/cache.proto config/cache/v1alpha/datasource.proto
//go:generate protoc --proto_path=../protos/ --proto_path=../../../../gingersnap-cloud-api/ --go-grpc_opt=Mconfig/cache/v1alpha/region.proto=github.com/rigazilla/engytita-api-examples/golang/grpc/client/gingersnap-cloud-api/config/cache/v1alpha --go_opt=Mconfig/cache/v1alpha/region.proto=github.com/rigazilla/engytita-api-examples/golang/grpc/client/gingersnap-cloud-api/config/cache/v1alpha --go_out=. --go-grpc_out=.  server.proto
func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := gr.NewRegionStoreClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var regReq = gr.CreateRegionRequest{
		Region: &pb.Region{
			Name:       "Region1",
			Datasource: "Datasource1",
			Rule: &pb.Rule{
				RuleType: &pb.Rule_Jsonpath{
					Jsonpath: &pb.Jsonpath{
						Value: "some.domain.stores",
					},
				},
			},
			Expiration: &pb.Expiration{
				ExpirationType: &pb.Expiration_Schedule{
					Schedule: "0 0 1 * *",
				},
			},
		},
	}

	r, err := c.CreateRegion(ctx, &regReq)
	if err != nil {
		log.Printf("could not create region: %v", err)
	}
	log.Printf("CreateRegion response: %v", r)
	getRegReq := gr.GetRegionRequest{
		Name: "Region1",
	}
	r, err = c.GetRegion(ctx, &getRegReq)
	if err != nil {
		log.Printf("could not get region: %v", err)
	}
	log.Printf("GetRegion response: %v", r)
	regReq.Region.Datasource = "NewDataSource"
	r, err = c.CreateRegion(ctx, &regReq)
	if err != nil {
		log.Printf("could not create region: %v", err)
	}
	log.Printf("CreateRegion response: %v", r)

	r, err = c.GetRegion(ctx, &getRegReq)
	if err != nil {
		log.Printf("could not get region: %v", err)
	}
	log.Printf("GetRegion response: %v", r)
}
