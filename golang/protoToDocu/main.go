package main

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/rigazilla/engytita-api-examples/golang/config/cache/v1alpha"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Command below generates the set of .pb.go files. .proto comes for the engytita-api project
// imported as submodule of this repo.
// The --go_opt=module=.. strips out the default module for the generated files, so files are generated
// in the `config/cache/v1alpha` folder in the go module root and can be imported as
// `import "your-module-name/config/cache/v1alpha`
//go:generate protoc --proto_path=../../engytita-api  --go_out=.. --go_opt=module=github.com/engytita/engytita-api/example/golang config/cache/v1alpha/region.proto config/cache/v1alpha/cache.proto config/cache/v1alpha/datasource.proto
func main() {
	var cache = v1alpha.Cache{
		Name:      "cacheExample",
		Namespace: "nsExample",
		Regions: []*v1alpha.Region{{
			Name:       "Region1",
			Datasource: "Datasource1",
			Rule: &v1alpha.Rule{
				RuleType: &v1alpha.Rule_Jsonpath{
					Jsonpath: &v1alpha.Jsonpath{
						Value: "some.domain.stores",
					},
				},
			},
			Expiration: &v1alpha.Expiration{
				ExpirationType: &v1alpha.Expiration_Schedule{
					Schedule: "0 0 1 * *",
				},
			},
		},
			{
				Name:       "Region2",
				Datasource: "Datasource2",
				Rule: &v1alpha.Rule{
					RuleType: &v1alpha.Rule_Wildcard{
						Wildcard: &v1alpha.Wildcard{
							Value: "/pets/(.*)",
						},
					},
				},
				Expiration: &v1alpha.Expiration{
					ExpirationType: &v1alpha.Expiration_Lifespan{
						Lifespan: &durationpb.Duration{
							Seconds: 24 * 3600,
						},
					},
				},
				Preload: &v1alpha.Preload{
					PreloadType: &v1alpha.Preload_Http{
						Http: &v1alpha.Http{
							Url: "value",
						},
					},
					Schedule: "0 0 1 * *",
				},
				Bound: &v1alpha.Bound{
					BoundType: &v1alpha.Bound_Count{
						Count: &v1alpha.Count{
							Value: 1000,
						},
					},
				},
			},
		},
	}
	printer := proto.TextMarshaler{}
	fmt.Println("============ Readable Protobuf Output =============")
	printer.Marshal(os.Stdout, &cache)
	jb, _ := protojson.Marshal(&cache)
	fmt.Println("============ Json Output =============")
	fmt.Println(string(jb))
}

func create(x int64) *int64 {
	return &x
}
