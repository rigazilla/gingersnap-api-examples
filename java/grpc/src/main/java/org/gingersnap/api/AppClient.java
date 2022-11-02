package org.gingersnap.api;

import java.util.Arrays;
import java.util.List;

import com.google.protobuf.Duration;
import com.google.protobuf.util.JsonFormat;

import gingersnap.api.service.cache.v1alpha2.GetResponse;
import gingersnap.config.cache.v1alpha1.LazyCachingRuleSpec;
import gingersnap.config.cache.v1alpha1.NamespacedRef;
import gingersnap.config.cache.v1alpha1.Value;
import io.grpc.Channel;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import rulestore.v1alpha1.CreateLazyRuleRequest;
import rulestore.v1alpha1.GetLazyRuleRequest;
import rulestore.v1alpha1.RuleStoreGrpc;

/**
 * Sample client code that makes gRPC calls to the server.
 */
public class AppClient {
	private final RuleStoreGrpc.RuleStoreBlockingStub blockingStub;

	/**
	 * Construct client for accessing RouteGuide server using the existing channel.
	 */
	public AppClient(Channel channel) {
		blockingStub = RuleStoreGrpc.newBlockingStub(channel);
	}

	public static void main(String[] args) {
		ManagedChannel channel = ManagedChannelBuilder.forTarget("localhost:8980").usePlaintext().build();
		AppClient client = new AppClient(channel);
		LazyCachingRuleSpec.Builder lazyRule1Builder = LazyCachingRuleSpec.newBuilder();
		// Creating fields for Lazy Rule
		{
			// Populating value
			List<String> valueColumns = Arrays.asList("name", "surname", "address");
			Value.Builder valueBuilder = Value.newBuilder()
					.addAllValueColumns(valueColumns);
			NamespacedRef.Builder ns = NamespacedRef.newBuilder().setName("myCache").setNamespace("myNamespace");
			// Assembling Lazy Rule
			lazyRule1Builder.setQuery("select name,surname,address,age from myTable where name='?' and value='?'");
			// Adding Value
			lazyRule1Builder.setValue(valueBuilder);
			// Adding ref to the cache
			lazyRule1Builder.setCacheRef(ns);
		}

		// Creating the request message
		CreateLazyRuleRequest.Builder clrrBuilder = CreateLazyRuleRequest.newBuilder().setRule(lazyRule1Builder);
		LazyCachingRuleSpec response = client.blockingStub.createLazyRule(clrrBuilder.build());
		// Print old rule for the same key (namespace.name)
		try {
			System.out.println("createLazyRule response");
			System.out.println(JsonFormat.printer().print(response));
		} catch (Exception e) {
		}

		GetLazyRuleRequest.Builder glrrBuilder = GetLazyRuleRequest.newBuilder().setName("myNamespace.myCache");

		LazyCachingRuleSpec getResponse = client.blockingStub.getLazyRule(glrrBuilder.build());
		// Print rule
		try {
			System.out.println("getLazyRule response");
			System.out.println(JsonFormat.printer().print(getResponse));
		} catch (Exception e) {
		}
		// Updating same rule
		LazyCachingRuleSpec.Builder rule2Builder = LazyCachingRuleSpec.newBuilder(getResponse);
		rule2Builder.setQuery("select * from myTable where name='?' and value='?'");
		CreateLazyRuleRequest.Builder clrrBuilder2 = CreateLazyRuleRequest.newBuilder().setRule(rule2Builder);
		LazyCachingRuleSpec response2 = client.blockingStub.createLazyRule(clrrBuilder2.build());
		// Print old rule for the same key (namespace.name)
		try {
			System.out.println("createLazyRule response");
			System.out.println(JsonFormat.printer().print(response2));
		} catch (Exception e) {
		}
		channel.shutdown();
	}
}
