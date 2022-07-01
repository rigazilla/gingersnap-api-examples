package org.engytita.api;

import com.google.protobuf.Duration;
import com.google.protobuf.util.JsonFormat;

import config.cache.v1alpha.Bound;
import config.cache.v1alpha.Count;
import config.cache.v1alpha.Expiration;
import config.cache.v1alpha.Http;
import config.cache.v1alpha.Preload;
import config.cache.v1alpha.Region;
import config.cache.v1alpha.Rule;
import config.cache.v1alpha.Wildcard;
import io.grpc.Channel;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import regionstore.v1alpha.CreateRegionRequest;
import regionstore.v1alpha.GetRegionRequest;
import regionstore.v1alpha.RegionStoreGrpc;
import regionstore.v1alpha.RegionStoreGrpc.RegionStoreBlockingStub;

/**
 * Sample client code that makes gRPC calls to the server.
 */
public class AppClient {
	private final RegionStoreBlockingStub blockingStub;

	/**
	 * Construct client for accessing RouteGuide server using the existing channel.
	 */
	public AppClient(Channel channel) {
		blockingStub = RegionStoreGrpc.newBlockingStub(channel);
	}

	public static void main(String[] args) {
	    ManagedChannel channel = ManagedChannelBuilder.forTarget("localhost:8980").usePlaintext().build();
		AppClient client = new AppClient(channel);
		// Build region2
		// {"name": "region2", "datasource": "DataSource2"
		// , "rule": {"wildcard": {"value": "/pets/(.*)"}}
		// , "preload": {"http": {"url": "value"},"schedule": "0 0 1 * *"}
		// ,"expiration": {"lifespan": "86400s"}
		// ,"bound": {"count": {"value": "1000"}}}
		Wildcard.Builder wc = Wildcard.newBuilder().setValue("/pets/(.*)");
		Region.Builder r = Region.newBuilder().setName("region2").setDatasource("DataSource2");
		Http.Builder ht = Http.newBuilder().setUrl("value");
		Count.Builder co = Count.newBuilder().setValue(1000);
		Bound.Builder bl = Bound.newBuilder().setCount(co);
		Preload.Builder pr = Preload.newBuilder().setSchedule("0 0 1 * *").setHttp(ht);
		Rule.Builder r1 = Rule.newBuilder().setWildcard(wc);
		Duration.Builder ls = Duration.newBuilder().setSeconds(24 * 3600);
		Expiration.Builder ex = Expiration.newBuilder().setLifespan(ls);
		r.setRule(r1).setPreload(pr).setBound(bl).setExpiration(ex);
		Region regIn = r.build();
		CreateRegionRequest.Builder crr = CreateRegionRequest.newBuilder().setRegion(regIn);
		Region crResp = client.blockingStub.createRegion(crr.build());
		try {
			System.out.println(JsonFormat.printer().print(crResp));
		} catch (Exception e) {
		}
		GetRegionRequest.Builder grr = GetRegionRequest.newBuilder().setName(regIn.getName());
		crResp = client.blockingStub.getRegion(grr.build());
		try {
			System.out.println(JsonFormat.printer().print(crResp));
		} catch (Exception e) {
		}
		Region.Builder newR = Region.newBuilder(crResp);
		newR.setDatasource(("newDataSource2"));
		crr = CreateRegionRequest.newBuilder().setRegion(newR.build());
		crResp = client.blockingStub.createRegion(crr.build());
		try {
			System.out.println(JsonFormat.printer().print(crResp));
		} catch (Exception e) {
		}
	}
}