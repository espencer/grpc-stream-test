# GRPC Stream Test

PoC and/or testing of GRPC streaming, performance, memory impact, etc.

## Getting started

_Optional_ Set GOPATH to `pwd`:

```sh
export GOPATH=`pwd`
```

Install Protocol Buffers for your platform: https://github.com/google/protobuf/releases

Download _go_ dependencies:

```sh
go get google.golang.org/grpc
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

```

Generate `.go` from `.proto`:

```
protoc streaming_service.proto --go_out=plugins=grpc:.
```

Running the server:

```
go run server.go streaming_service.pb.go
```

Running the client, in other terminal:

```
export GOPATH=`pwd`
go run client.go streaming_service.pb.go
```
