
gobin:
	mkdir gobin

gobin/protoc-gen-go: gobin
	go build -o ./gobin/protoc-gen-go ./vendor/github.com/golang/protobuf/protoc-gen-go

gobin/example-client: gobin
	go build -o ./gobin/example-client ./cmd/example-client
gobin/log-server: gobin
	go build -o ./gobin/log-server ./cmd/log-server

.PHONY: binaries
binaries: gobin/protoc-gen-go gobin/example-client gobin/log-server

# github.com/golang/protobuf/protoc-gen-go
CMD := protoc --plugin=protoc-gen-go=./gobin/protoc-gen-go --go_out=paths=source_relative,plugins=grpc:./genproto --proto_path=./proto

.PHONY: rmgenerate
rmgenerate: 
	rm -rf genproto

.PHONY: regenerate
regenerate: rmgenerate generate

.PHONY: generate
generate: gobin/protoc-gen-go
	mkdir -p genproto
	find ./proto -name *.proto | xargs -n1 $(CMD)

.PHONY: test
test:
	go test ./pkg/...

.PHONY: format
format:
	find ./pkg ./cmd -name *.go | xargs -n1 gofmt -s -w 

