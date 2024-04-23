### Todo List:
##### Implement a proper healthcheck for the gRPC service. 
```
healthcheck:
            test: ["CMD", "bin/grpc_health_probe-linux-amd64", "-addr=localhost:50051"]
            interval: 30s
            timeout: 30s
            retries: 3
```
Source: https://pkg.go.dev/google.golang.org/grpc/health/grpc_health_v1

#### Migrate elasticsearch and gRPC clients to credential usage 

#### Unit-testing for all the components

#### Add documentation

#### Handle elasticsearch connection error

#### Dynamically picking up the `service` ip for the gRCP client



### Resources:
* ElasticSearch go-client docs: https://pkg.go.dev/github.com/elastic/go-elasticsearch/v8/esapi@v8.13.1#Get

