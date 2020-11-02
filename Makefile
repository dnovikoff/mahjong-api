include protoc.mk

gobin:
	mkdir gobin

gobin/example-client: gobin
	go build -mod vendor -o ./gobin/example-client ./cmd/example-client

gobin/log-server: gobin
	go build -mod vendor -o ./gobin/log-server ./cmd/log-server

.PHONY: binaries
binaries: gobin/protoc-gen-go gobin/example-client gobin/log-server

# github.com/golang/protobuf/protoc-gen-go
CMD := $(protoc_go_cmd) --go_out=paths=source_relative,plugins=grpc:./genproto --proto_path=./proto

.PHONY: rmgenerate
rmgenerate: 
	rm -rf genproto

.PHONY: regenerate
regenerate: rmgenerate generate

.PHONY: generate
generate: $(protoc_gen_go)
	mkdir -p genproto
	find ./proto -name *.proto | xargs -n1 $(CMD)

.PHONY: test
test:
	go test -mod vendor ./pkg/...

.PHONY: format
format:
	find ./pkg ./cmd -name *.go | xargs -n1 gofmt -s -w 

