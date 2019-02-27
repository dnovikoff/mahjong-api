
gobin:
	mkdir gobin

gobin/protoc-gen-go: gobin
	go build -o ./gobin/protoc-gen-go ./vendor/github.com/golang/protobuf/protoc-gen-go

gobin/example-client: gobin
	go build -o ./gobin/example-client ./cmd/example-client/client.go

.PHONY: binaries
binaries: gobin/protoc-gen-go gobin/example-client

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
	$(CMD) ./proto/base/*.proto
	$(CMD) ./proto/rules/*.proto
	$(CMD) ./proto/log/*.proto
	$(CMD) ./proto/api/*.proto

.PHONY: test
test:
	go test ./pkg/...
