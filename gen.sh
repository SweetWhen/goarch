#!/bin/bash

rm -rf custompb/proto/gen/*

#--go_opt=module=$prefix gen pb.go to another path
#--go_opt=paths=source_relative gen pb.go to the path the same with .proto file
protoc --go_out=./custompb --go_opt=module=github.com/SweetWhen/goarch/custompb custompb/proto/interceptor/*.proto
protoc --go_out=./custompb --go_opt=module=github.com/SweetWhen/goarch/custompb \
        --go-grpc_out=./custompb --go-grpc_opt=module=github.com/SweetWhen/goarch/custompb custompb/proto/helloworld/*.proto


