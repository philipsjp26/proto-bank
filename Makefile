#!/bin/bash
run: 
	@go run main.go

gen-proto:
	@protoc --proto_path=api/${DIR} api/${DIR}/proto/*.proto --go_out=api/${DIR} --go-grpc_out=api/${DIR}