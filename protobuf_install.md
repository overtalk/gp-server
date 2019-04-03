# Documents

- [Introduction](./docs/introduction.md)

# Development Guide

### Preparation
```
git clone https://github.com/googleapis/googleapis ../
rsync -av --delete ../googleapis/google/ ./google

brew install -y protobuf grpc

go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
go get -u github.com/ckaznocha/protoc-gen-lint
go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

```

### Generators and Scripts

- `lint.sh` to check .proto file syntax, etc.
- `docgen.sh` to generating documents
- `codegen_go.sh` generating source code for golang
