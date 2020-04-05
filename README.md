# Debugging with Telepresence!
[Telepresence](https://www.telepresence.io/) is fast, local development fro kubernetes and OpenShift micro-services.
This repository is sample repository for debugging with Telepresence for kubernetes pod.
This sample project include two application, gateway and backend.
These applications deploy with envoy as a sidecar proxy, and each application communicate through envoy.


## install

### Telepresence

https://www.telepresence.io/reference/install

### Usave

1. Create GKE cluster and deploy services.
```shell script
$ make init
... wait for while

## build images
$ make build

## replace tag name 

## apply
$ make apply

## connect to service & send request
$ kubectl port-forward service/gateway-service 8080:8080
Forwarding from 127.0.0.1:8080 -> 10000
Forwarding from [::1]:8080 -> 10000

... open another tab
$ curl localhost:8080/greeting?name=tom
message":"Hey! tom, Nice to meet you!!","datetime":"2020-04-05 07:26:03 +0000 UTC"

```

## LICENSE
MIT



