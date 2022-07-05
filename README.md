# Setup

Install gRPC tooling

```
brew install protobuf
brew install protoc-gen-go

go install google.golang.org/grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

# Generate gRPC code

`protoc -I=./pb --go_out=./src/server ./pb/messages.proto`