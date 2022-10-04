package main

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/rigazilla/gingersnap-api-examples/golang/protobuf/gingersnap-api/config/cache/v1alpha"
	"google.golang.org/protobuf/encoding/protojson"
	yamme "sigs.k8s.io/yaml"
)

// Command below generates the set of .pb.go files. .proto comes for the gingersnap-api project
// imported as submodule of this repo.
// The --go_opt=module=.. strips out the default module for the generated files, so files are generated
// in the `config/cache/v1alpha` folder in the go module root and can be imported as
// `import "your-module-name/config/cache/v1alpha`
//go:generate protoc --proto_path=../../../gingersnap-api  --go_out=.. config/cache/v1alpha/region.proto config/cache/v1alpha/cache.proto config/cache/v1alpha/datasource.proto

func main() {
	yaml := `
name: cacheExample
namespace: nsExample
regions:
  - name: Region1
    datasource: Datasource1
    rule:
        jsonpath:
          value: some.domain.stores
    expiration:
        schedule: 0 0 1 * *
  - name: Region2
    datasource: Datasource2
    rule:
        wildcard:
          value: /pets/(.*)
    preload:
        http:
          url: value
        schedule: 0 0 1 * *
    expiration:
        lifespan: 86400s
    bound:
        count:
          value: 1000
`
	json, err := yamme.YAMLToJSON([]byte(yaml))
	if err != nil {
		fmt.Printf("json err: %v\n", err)
	}
	cache := &v1alpha.Cache{}
	err = protojson.Unmarshal(json, cache)
	if err != nil {
		fmt.Printf("proto err: %v\n", err)
	}
	printer := proto.TextMarshaler{}
	fmt.Println("============ Readable Protobuf Output =============")
	printer.Marshal(os.Stdout, cache)
}
