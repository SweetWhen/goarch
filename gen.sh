#!/bin/bash

rm -rf custompb/proto/gen/*

#pbDir="custompb/proto"
protoc --go_out=./ custompb/proto/interceptor/*.proto
protoc --go_out=./ --go-grpc_out=./  custompb/proto/helloworld/*.proto


