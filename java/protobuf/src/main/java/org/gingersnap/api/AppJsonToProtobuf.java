package org.gingersnap.api;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;
import com.google.protobuf.TextFormat;
import com.google.protobuf.util.JsonFormat;

import gingersnap.config.cache.v1alpha1.*;  

public class AppJsonToProtobuf {
    // Conversion yaml->json->protobuf->json
    public static void main(String[] args) {
        CacheConf.Builder cache = CacheConf.newBuilder();
        try {
            String json = convertYamlToJson(yaml);
            JsonFormat.parser().merge(json, cache);
            System.out.println(TextFormat.printer().escapingNonAscii(false).shortDebugString(cache.build()));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    private static String convertYamlToJson(String yaml) throws JsonMappingException, JsonProcessingException {
        ObjectMapper yamlReader = new ObjectMapper(new YAMLFactory());
        Object obj = yamlReader.readValue(yaml, Object.class);

        ObjectMapper jsonWriter = new ObjectMapper();
        return jsonWriter.writeValueAsString(obj);
    }
    public static String json = 
    "{"+
"  \"eagerCachingRuleSpecs\": {"+
"    \"myEagerCacheRule1\": {"+
"      \"cacheRef\": {"+
"        \"name\": \"myCache\","+
"        \"namespace\": \"myNamespace\""+
"      },"+
"      \"resources\": {"+
"        \"requests\": {"+
"          \"memory\": {"+
"            \"string\": \"1Gi\""+
"          },"+
"          \"cpu\": {"+
"            \"string\": \"500m\""+
"          }"+
"        },"+
"        \"limits\": {"+
"          \"memory\": {"+
"            \"string\": \"2Gi\""+
"          },"+
"          \"cpu\": {"+
"            \"string\": \"1\""+
"          }"+
"        }"+
"      },"+
"      \"tableName\": \"TABLE_EAGER_RULE_1\","+
"      \"key\": {"+
"        \"format\": \"JSON\","+
"        \"keySeparator\": \",\","+
"        \"keyColumns\": ["+
"          \"col1\","+
"          \"col3\","+
"          \"col4\""+
"        ]"+
"      },"+
"      \"value\": {"+
"        \"valueColumns\": ["+
"          \"col6\","+
"          \"col7\","+
"          \"col8\""+
"        ]"+
"      }"+
"    }"+
"  },"+
"  \"LazyCachingRuleSpecs\": {"+
"    \"myLazyCacheRule1\": {"+
"      \"cacheRef\": {"+
"        \"name\": \"myCache\","+
"        \"namespace\": \"myNamespace\""+
"      },"+
"      \"query\": \"select name,surname,address,age from myTable where name='?' and value='?'\","+
"      \"value\": {"+
"        \"valueColumns\": ["+
"          \"name\","+
"          \"surname\","+
"          \"address\""+
"        ]"+
"      }"+
"    }"+
"  }"+
"}";
public static String yaml = 
"eagerCachingRuleSpecs:\n"+
"  myEagerCacheRule1:\n"+
"    cacheRef:\n"+
"      name: myCache\n"+
"      namespace: myNamespace\n"+
"    resources:\n"+
"      requests:\n"+
"        memory:\n"+
"          string: 1Gi\n"+
"        cpu:\n"+
"          string: 500m\n"+
"      limits:\n"+
"        memory:\n"+
"          string: 2Gi\n"+
"        cpu:\n"+
"          string: \"1\"\n"+
"    tableName: TABLE_EAGER_RULE_1\n"+
"    key:\n"+
"      format: JSON\n"+
"      keySeparator: ','\n"+
"      keyColumns:\n"+
"        - col1\n"+
"        - col3\n"+
"        - col4\n"+
"    value:\n"+
"      valueColumns:\n"+
"        - col6\n"+
"        - col7\n"+
"        - col8\n"+
"LazyCachingRuleSpecs:\n"+
"  myLazyCacheRule1:\n"+
"    cacheRef:\n"+
"      name: myCache\n"+
"      namespace: myNamespace\n"+
"    query: select name,surname,address,age from myTable where name='?' and value='?'\n"+
"    value:\n"+
"      valueColumns:\n"+
"        - name\n"+
"        - surname\n"+
"        - address\n"+
"\n";
}