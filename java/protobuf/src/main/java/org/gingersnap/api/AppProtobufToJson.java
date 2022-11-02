package org.gingersnap.api;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import com.google.protobuf.util.JsonFormat;
import com.google.protobuf.Duration;
import gingersnap.config.cache.v1alpha1.*;
import k8s.io.apimachinery.pkg.api.resource.*;

/**
 * Hello world!
 *
 */
public class AppProtobufToJson {
        // Conversion protobuf -> json
        public static void main(String[] args) {
                EagerCachingRuleSpec.Builder eagerRule1Builder = EagerCachingRuleSpec.newBuilder();
                // Creating fields for Eager Rule
                // Populating name ref
                NamespacedRef.Builder ns = NamespacedRef.newBuilder().setName("myCache").setNamespace("myNamespace");

                {
                        // Populating resources
                        Resources.Builder eagerRule1ResourcesBuilder = Resources.newBuilder();
                        ResourceQuantity.Builder eagerRule1LimitsBuilder = ResourceQuantity.newBuilder();
                        ResourceQuantity.Builder eagerRule1RequestsBuilder = ResourceQuantity.newBuilder();
                        QuantityOuterClass.Quantity.Builder cpuLimits = QuantityOuterClass.Quantity.newBuilder();
                        QuantityOuterClass.Quantity.Builder cpuRequests = QuantityOuterClass.Quantity.newBuilder();
                        QuantityOuterClass.Quantity.Builder memoryLimits = QuantityOuterClass.Quantity.newBuilder();
                        QuantityOuterClass.Quantity.Builder memoryRequests = QuantityOuterClass.Quantity.newBuilder();
                        memoryLimits.setString("2Gi");
                        memoryRequests.setString("1Gi");
                        cpuLimits.setString("1");
                        cpuRequests.setString("500m");
                        eagerRule1LimitsBuilder.setCpu(cpuLimits).setMemory(memoryLimits);
                        eagerRule1RequestsBuilder.setCpu(cpuRequests).setMemory(memoryRequests);
                        eagerRule1ResourcesBuilder.setLimits(eagerRule1LimitsBuilder)
                                        .setRequests(eagerRule1RequestsBuilder);
                        // Populating key
                        List<String> keyColumns = Arrays.asList("col1", "col3", "col4");
                        Key.Builder keyBuilder = Key.newBuilder()
                                        .setFormat(KeyFormat.JSON)
                                        .setKeySeparator(",")
                                        .addAllKeyColumns(keyColumns);
                        // Populating value
                        List<String> valueColumns = Arrays.asList("col6", "col7", "col8");
                        Value.Builder valueBuilder = Value.newBuilder()
                                        .addAllValueColumns(valueColumns);
                        // Assembling Eager Rule
                        eagerRule1Builder.setTableName("TABLE_EAGER_RULE_1");
                        // Adding resources
                        eagerRule1Builder.setResources(eagerRule1ResourcesBuilder);
                        // Adding key
                        eagerRule1Builder.setKey(keyBuilder);
                        // Adding value
                        eagerRule1Builder.setValue(valueBuilder);
                        // Adding ref to the cache
                        eagerRule1Builder.setCacheRef(ns);
                }


                LazyCachingRuleSpec.Builder lazyRule1Builder = LazyCachingRuleSpec.newBuilder();
                // Creating fields for Eager Rule
                {
                        // Populating value
                        List<String> valueColumns = Arrays.asList("name", "surname", "address");
                        Value.Builder valueBuilder = Value.newBuilder()
                                        .addAllValueColumns(valueColumns);
                        // Assembling Lazy Rule
                        lazyRule1Builder.setQuery("select name,surname,address,age from myTable where name='?' and value='?'");
                        // Adding Value
                        lazyRule1Builder.setValue(valueBuilder);
                        // Adding ref to the cache
                        lazyRule1Builder.setCacheRef(ns);
                }

                // Build cache
                CacheConf.Builder cacheConfBuilder = CacheConf.newBuilder()
                        .putEagerCachingRuleSpecs("myEagerCacheRule1", eagerRule1Builder.build())
                        .putLazyCachingRuleSpecs("myLazyCacheRule1", lazyRule1Builder.build());
                CacheConf cacheConf = cacheConfBuilder.build();
                try {
                        System.out.println(JsonFormat.printer().print(cacheConf));
                } catch (Exception e) {
                }
        }
}
