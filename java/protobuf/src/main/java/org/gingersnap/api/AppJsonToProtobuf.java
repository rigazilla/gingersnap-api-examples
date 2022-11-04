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
public static String yaml = 
"eagerCachingRuleSpecs:\n"+
"  myEagerCacheRule1:\n"+
"    cacheRef:\n"+
"      name: myCache\n"+
"      namespace: myNamespace\n"+
"    resources:\n"+
"      requests:\n"+
"        memory: \"1Gi\"\n"+
"        cpu: \"500m\"\n"+
"      limits:\n"+
"        memory: \"2Gi\"\n"+
"        cpu:  \"1\"\n"+
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