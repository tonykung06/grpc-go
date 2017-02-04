## Generating proto go files
- `go get -u google.golang.org/grpc  github.com/golang/protobuf/proto github.com/golang/protobuf/protoc-gen-go`
- `protoc -I ./pb ./pb/messages.proto --go_out=plugins=grpc:./pb`

## Compiling
- `go install github.com/grpc-go/server`
- `go install github.com/grpc-go/client`