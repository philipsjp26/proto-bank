#!/bin/bash
run: 
	@go run main.go

gen-proto:
	@protoc  -I. api/meta/proto/meta.proto --proto_path=api/${DIR} api/${DIR}/proto/*.proto --go_out=api/${DIR} --go-grpc_out=api/${DIR}