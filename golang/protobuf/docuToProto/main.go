package main

import (
	"fmt"

	"encoding/json"

	"github.com/rigazilla/gingersnap-api-examples/golang/protobuf/docuToProto/gingersnap-api/config/cache/v1alpha1"
	"google.golang.org/protobuf/types/known/structpb"
	yamme "sigs.k8s.io/yaml"
)

func example1(x *structpb.Struct) string {
	return x.String()
}

// Command below generates the set of .pb.go files. .proto comes for the gingersnap-api project
// imported as submodule of this repo.
// The --go_opt=module=.. strips out the default module for the generated files, so files are generated
// in the `gingersnap-api/config/cache/v1alpha` folder in the go module root and can be imported as
// `import "your-module-name/gingersnap-api/config/cache/v1alpha`
//go:generate protoc --proto_path=../../.. --proto_path=../../../gingersnap-api --go_out=. --go_opt=Mgingersnap-api/config/cache/v1alpha1/cache.proto=github.com/rigazilla/gingersnap-api-examples/golang/grpc/example/client/api/config/cache/v1alpha1 --go_opt=paths=source_relative gingersnap-api/config/cache/v1alpha1/cache.proto

func main() {
	yaml := `
    cacheSpec:
      resources:
        requests:
          memory: "4Gi"
          cpu: "2"
        limits:
          memory: "8Gi"
          cpu: "4"
      dataSource:
        connectionProperties:
          prop1: value1
          prop2: value2    
    eagerCachingRuleSpecs:
      myEagerCacheRule:
        cacheRef:
          name: myCache
          namespace: myNamespace
        resources:
          requests:
            memory: "2Gi"
            cpu: "500m"
          limits:
            memory: "2Gi"
            cpu: "1"
        tableName: TABLE_EAGER_RULE_2
        key:
          format: 1
          keySeparator: ','
          keyColumns:
            - col2
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
		fmt.Printf("\n\nproto err: %v\n", err)
	}
	fmt.Println("============ Readable Protobuf Output =============")

	readableCache, _ := json.Marshal(cache)
	fmt.Printf("Readable Cache:\n %s\n", readableCache)
}
