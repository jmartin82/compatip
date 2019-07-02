protoc -I{GOPATH}/src/github.com/jmartin82/compatip --go_out=plugins=grpc:. pkg/rpc/version.proto
