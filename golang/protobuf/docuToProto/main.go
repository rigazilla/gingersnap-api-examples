package main

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/rigazilla/gingersnap-api-examples/golang/protobuf/docuToProto/api/v1alpha1"
	yamme "sigs.k8s.io/yaml"
)

// Command below generates the set of .pb.go files. .proto comes for the gingersnap-api project
// imported as submodule of this repo.
// The --go_opt=module=.. strips out the default module for the generated files, so files are generated
// in the `config/cache/v1alpha` folder in the go module root and can be imported as
// `import "your-module-name/config/cache/v1alpha`
//go:generate protoc --proto_path=gingersnap-api  --go_out=. config/cache/v1alpha1/cache.proto apimachinery/pkg/api/resource/quantity.proto

func main() {
	yaml := `
  eagerCachingRuleSpecs:
    myEagerCacheRule1:
      cacheRef:
        name: myCache
        namespace: myNamespace
      resources:
        requests:
          memory: "1Gi"
          cpu: "500m"
        limits:
          memory: "2Gi"
          cpu: "1"
      tableName: TABLE_EAGER_RULE_1
      key:
        format: JSON
        keySeparator: ','
        keyColumns:
          - col1
          - col3
          - col4
      value:
        valueColumns:
          - col6
          - col7
          - col8
  lazyCachingRuleSpecs:
    myLazyCacheRule1:
      cacheRef:
        name: myCache
        namespace: myNamespace
      query: select name,surname,address,age from myTable where name='?' and value='?'
      value:
        valueColumns:
          - name
          - surname
          - address
`
	jsonString, err := yamme.YAMLToJSON([]byte(yaml))
	if err != nil {
		fmt.Printf("json err: %v\n", err)
	}
	cache := &v1alpha1.CacheConf{}
	err = json.Unmarshal(jsonString, cache)
	if err != nil {
		fmt.Printf("proto err: %v\n", err)
	}
	printer := proto.TextMarshaler{}
	fmt.Println("============ Readable Protobuf Output =============")
	printer.Marshal(os.Stdout, cache)
	fmt.Printf("Resources\n")
	fmt.Printf("    Requests: memory=%d, cpu=%d\n",
		cache.EagerCachingRuleSpecs["myEagerCacheRule1"].Resources.Requests.Memory.Value(),
		cache.EagerCachingRuleSpecs["myEagerCacheRule1"].Resources.Requests.Cpu.MilliValue())
	fmt.Printf("    Limits: memory=%d, cpu=%d\n",
		cache.EagerCachingRuleSpecs["myEagerCacheRule1"].Resources.Limits.Memory.Value(),
		cache.EagerCachingRuleSpecs["myEagerCacheRule1"].Resources.Limits.Cpu.MilliValue())

}
