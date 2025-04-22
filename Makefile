LOCAL_BIN:=$(CURDIR)/bin

init:
	ls go.mod || go mod init gitlab.com/martyn.andrw/microlink
	go mod tidy

proto-gen:
	rm -rf pkg/*
	PATH=$(PATH):$(LOCAL_BIN) protoc -I=proto/ \
		--go_out=pkg/ \
		--go-grpc_out=pkg/ \
		proto/link.proto \
		--experimental_allow_proto3_optional
	mv pkg/gitlab.com/martyn.andrw/microlink/pkg/* pkg/
	rm -rf pkg/gitlab.com

run:
	go run main.go