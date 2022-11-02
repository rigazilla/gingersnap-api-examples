package org.gingersnap.api;

import java.io.IOException;

import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.TimeUnit;
import java.util.logging.Logger;

import gingersnap.config.cache.v1alpha1.LazyCachingRuleSpec;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;
import rulestore.v1alpha1.CreateLazyRuleRequest;
import rulestore.v1alpha1.GetLazyRuleRequest;
import rulestore.v1alpha1.RuleStoreGrpc.RuleStoreImplBase;

public class AppServer {
	  private static final Logger logger = Logger.getLogger(AppServer.class.getName());

	  private final int port;
	  private final Server server;

	  /** Create a RouteGuide server listening on {@code port} using {@code featureFile} database. */
	  public AppServer(int port) throws IOException {
	    this(ServerBuilder.forPort(port), port);
	  }

	  /** Create a RouteGuide server using serverBuilder as a base and features as data. */
	  public AppServer(ServerBuilder<?> serverBuilder, int port) {
	    this.port = port;
	    server = serverBuilder.addService(new RulestoreService())
	        .build();
	  }

	  /** Start serving requests. */
	  public void start() throws IOException {
	    server.start();
	    logger.info("Server started, listening on " + port);
	    Runtime.getRuntime().addShutdownHook(new Thread() {
	      @Override
	      public void run() {
	        // Use stderr here since the logger may have been reset by its JVM shutdown hook.
	        System.err.println("*** shutting down gRPC server since JVM is shutting down");
	        try {
	          AppServer.this.stop();
	        } catch (InterruptedException e) {
	          e.printStackTrace(System.err);
	        }
	        System.err.println("*** server shut down");
	      }
	    });
	  }

	  /** Stop serving requests and shutdown resources. */
	  public void stop() throws InterruptedException {
	    if (server != null) {
	      server.shutdown().awaitTermination(30, TimeUnit.SECONDS);
	    }
	  }

	  /**
	   * Await termination on the main thread since the grpc library uses daemon threads.
	   */
	  private void blockUntilShutdown() throws InterruptedException {
	    if (server != null) {
	      server.awaitTermination();
	    }
	  }

	  /**
	   * Main method.  This comment makes the linter happy.
	   */
	  public static void main(String[] args) throws Exception {
	    AppServer server = new AppServer(8980);
	    server.start();
	    server.blockUntilShutdown();
	  }
private static class RulestoreService extends   RuleStoreImplBase {
	
	private Map<String, LazyCachingRuleSpec> mapOfRules = new HashMap<>();
	@Override
	public void createLazyRule(CreateLazyRuleRequest request, StreamObserver<LazyCachingRuleSpec> responseObserver) {
		LazyCachingRuleSpec newR = request.getRule();
		// Using namespace.name as key
		String name = newR.getCacheRef().getNamespace()+"."+newR.getCacheRef().getName();
		LazyCachingRuleSpec oldR = mapOfRules.get(name);
		mapOfRules.put(name, newR);
		responseObserver.onNext(oldR);
		responseObserver.onCompleted();
	}

	@Override
	public void getLazyRule(GetLazyRuleRequest request, StreamObserver<LazyCachingRuleSpec> responseObserver) {
		String name = request.getName();
		LazyCachingRuleSpec region = mapOfRules.get(name);
		responseObserver.onNext(region);
		responseObserver.onCompleted();
	}
}
}
