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

	"github.com/rigazilla/gingersnap-api-examples/golang/grpc/example/client/gingersnap-api/config/cache/v1alpha1"
	gr "github.com/rigazilla/gingersnap-api-examples/golang/grpc/example/client/rulestore/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

//go:generate protoc --proto_path=../../../.. --include_source_info --go_out=. --go_opt=Mgingersnap-api/config/cache/v1alpha1/cache.proto=github.com/rigazilla/gingersnap-api-examples/golang/grpc/example/client/gingersnap-api/config/cache/v1alpha1 --go_opt=paths=source_relative gingersnap-api/config/cache/v1alpha1/cache.proto
//go:generate protoc --proto_path=../../../../grpc-proto/ --proto_path=../../../.. --go_opt=Mgingersnap-api/config/cache/v1alpha1/cache.proto=github.com/rigazilla/gingersnap-api-examples/golang/grpc/example/client/gingersnap-api/config/cache/v1alpha1 --go-grpc_opt=Mgingersnap-api/config/cache/v1alpha1/cache.proto=github.com/rigazilla/gingersnap-api-examples/golang/grpc/example/client/gingersnap-api/config/cache/v1alpha1 --go_out=. --go-grpc_out=.  rulestoreServer.proto
func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := gr.NewRuleStoreClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var ruleReq = gr.CreateLazyRuleRequest{
		Rule: &v1alpha1.LazyCachingRuleSpec{
			CacheRef: &v1alpha1.NamespacedRef{
				Name:      "ruleName",
				Namespace: "ruleNamespace",
			},
			Query: "select * from MY_TABLE where name='?' and surname='?'",
			Value: &v1alpha1.Value{
				ValueColumns: []string{"col1", "col3", "col4"},
			},
		},
	}

	r, err := c.CreateLazyRule(ctx, &ruleReq)
	if err != nil {
		log.Printf("could not create rule: %v", err)
	}
	log.Printf("CreateLazyRule response: %v", r)
	getRuleReq := gr.GetLazyRuleRequest{
		Name: "ruleNamespace.ruleName",
	}
	r, err = c.GetLazyRule(ctx, &getRuleReq)
	if err != nil {
		log.Printf("could not get rule: %v", err)
	}
	log.Printf("GetLazyRule response: %v", r)

	ruleReq.Rule.Query = "select name,age from MY_TABLE where name='?' and surname='?'"
	r, err = c.CreateLazyRule(ctx, &ruleReq)
	if err != nil {
		log.Printf("could not create rule: %v", err)
	}
	log.Printf("CreateLazyRule response: %v", r)

	r, err = c.GetLazyRule(ctx, &getRuleReq)
	if err != nil {
		log.Printf("could not get rule: %v", err)
	}
	log.Printf("GetLazyRule response: %v", r)
}
