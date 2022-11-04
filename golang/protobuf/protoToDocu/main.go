package main

import (
	"encoding/json"
	"fmt"

	"github.com/rigazilla/gingersnap-api-examples/golang/protobuf/protoToDocu/gingersnap-api/config/cache/v1alpha1"
)

// Command below generates the set of .pb.go files. .proto comes for the gingersnap-api project
// imported as submodule of this repo.
// The --go_opt=module=.. strips out the default module for the generated files, so files are generated
// in the `gingersnap-api/config/cache/v1alpha` folder in the go module root and can be imported as
// `import "your-module-name/gingersnap-api/config/cache/v1alpha`
//go:generate protoc --proto_path=../../.. --include_source_info --go_out=. --go_opt=Mgingersnap-api/config/cache/v1alpha1/cache.proto=github.com/rigazilla/gingersnap-api-examples/golang/grpc/example/client/api/config/cache/v1alpha1 --go_opt=paths=source_relative gingersnap-api/config/cache/v1alpha1/cache.proto
func main() {
	// TODO: use k8s type for quantity. Java side needs some work for this
	// cpuRequest := resource.MustParse("1")
	// memoryRequest := resource.MustParse("1Gi")
	// cpuLimit := resource.MustParse("2")
	// memoryLimit := resource.MustParse("2Gi")
	cpuRequest := "1"
	memoryRequest := "1Gi"
	cpuLimit := "2"
	memoryLimit := "2Gi"
	cache := v1alpha1.CacheSpec{
		Resources: &v1alpha1.Resources{
			Requests: &v1alpha1.ResourceQuantity{
				Cpu:    cpuRequest,
				Memory: memoryRequest,
			},
			Limits: &v1alpha1.ResourceQuantity{
				Cpu:    cpuLimit,
				Memory: memoryLimit,
			},
		},
	}

	eagerRule := v1alpha1.EagerCachingRuleSpec{
		CacheRef: &v1alpha1.NamespacedRef{
			Name:      "cacheName",
			Namespace: "myNamespace",
		},
		TableName: "MY_TABLE",
		Key: &v1alpha1.Key{
			KeyColumns: []string{"kc1", "kc3", "kc5"},
			Format:     v1alpha1.KeyFormat_JSON,
		},
	}

	lazyRule := v1alpha1.LazyCachingRuleSpec{
		Query: "select * from MY_TABLE where name='?' and surname='?'",
		Value: &v1alpha1.Value{
			ValueColumns: []string{"col1", "col3", "col4"},
		},
	}

	cacheConf := v1alpha1.CacheConf{
		CacheSpec:             &cache,
		EagerCachingRuleSpecs: map[string]*v1alpha1.EagerCachingRuleSpec{"myEagerRule": &eagerRule},
		LazyCachingRuleSpecs:  map[string]*v1alpha1.LazyCachingRuleSpec{"myLazyRule": &lazyRule},
	}

	cacheConfJson, _ := json.Marshal(&cacheConf)

	fmt.Printf("CacheConfig to JSON:\n %s", string(cacheConfJson))
}
