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

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/rigazilla/gingersnap-api-examples/golang/grpc/example/server/gingersnap-api/config/cache/v1alpha1"
	gr "github.com/rigazilla/gingersnap-api-examples/golang/grpc/example/server/rulestore/v1alpha1"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	gr.UnimplementedRuleStoreServer
}

var mapOfRegion = make(map[string]*pb.LazyCachingRuleSpec)

// SayHello implements helloworld.GreeterServer
func (s *server) CreateLazyRule(ctx context.Context, in *gr.CreateLazyRuleRequest) (*pb.LazyCachingRuleSpec, error) {
	var rule = in.GetRule()
	log.Printf("Received: %v", rule)
	key := rule.CacheRef.Namespace + "." + rule.CacheRef.Name
	old := mapOfRegion[key]
	mapOfRegion[key] = rule
	return old, nil
}

func (s *server) GetLazyRule(ctx context.Context, in *gr.GetLazyRuleRequest) (*pb.LazyCachingRuleSpec, error) {
	log.Printf("Received: %v", in.Name)
	return mapOfRegion[in.Name], nil
}

//go:generate protoc --proto_path=../../../.. --go_out=. --go_opt=Mgingersnap-api/config/cache/v1alpha1/cache.proto=github.com/rigazilla/gingersnap-api-examples/golang/grpc/example/server/gingersnap-api/config/cache/v1alpha1 --go_opt=paths=source_relative gingersnap-api/config/cache/v1alpha1/cache.proto
//go:generate protoc --proto_path=../../../../grpc-proto/ --proto_path=../../../.. --go_opt=Mgingersnap-api/config/cache/v1alpha1/cache.proto=github.com/rigazilla/gingersnap-api-examples/golang/grpc/example/server/gingersnap-api/config/cache/v1alpha1 --go-grpc_opt=Mgingersnap-api/config/cache/v1alpha1/cache.proto=github.com/rigazilla/gingersnap-api-examples/golang/grpc/example/server/gingersnap-api/config/cache/v1alpha1 --go_out=. --go-grpc_out=.  rulestoreServer.proto
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	gr.RegisterRuleStoreServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
