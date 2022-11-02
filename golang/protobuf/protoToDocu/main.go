package main

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/rigazilla/gingersnap-api-examples/golang/protobuf/protoToDocu/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// Command below generates the set of .pb.go files. .proto comes for the gingersnap-api project
// imported as submodule of this repo.
// The --go_opt=module=.. strips out the default module for the generated files, so files are generated
// in the `config/cache/v1alpha` folder in the go module root and can be imported as
// `import "your-module-name/config/cache/v1alpha`
//go:generate protoc --proto_path=../../../gingersnap-api  --go_out=. config/cache/v1alpha1/cache.proto apimachinery/pkg/api/resource/quantity.proto
func main() {
	cpuRequest := resource.MustParse("1")
	memoryRequest := resource.MustParse("1Gi")
	cpuLimit := resource.MustParse("2")
	memoryLimit := resource.MustParse("2Gi")
	cache := v1alpha1.CacheSpec{
		Resources: &v1alpha1.Resources{
			Requests: &v1alpha1.ResourceQuantity{
				Cpu:    &cpuRequest,
				Memory: &memoryRequest,
			},
			Limits: &v1alpha1.ResourceQuantity{
				Cpu:    &cpuLimit,
				Memory: &memoryLimit,
			},
		},
	}
	lazyRule := v1alpha1.LazyCachingRuleSpec{
		Query: "select * from MY_TABLE where name='?' and surname='?'",
		Value: &v1alpha1.Value{
			ValueColumns: []string{"col1", "col3", "col4"},
		},
	}

	printer := proto.TextMarshaler{}
	fmt.Println("============ Readable Protobuf Output =============")
	printer.Marshal(os.Stdout, &lazyRule)
	printer.Marshal(os.Stdout, &cache)
	fmt.Println("============ Json Output =============")
	jb, _ := json.Marshal(&lazyRule)
	fmt.Println(string(jb))
	jb, _ = json.Marshal(&cache)
	fmt.Println(string(jb))
}
