
.PHONY: idl
idl:
	@bash gen.sh

.PHONY: pb-debug
pb-debug:
	@rm -rf custompb/proto/gen/debug && mkdir -p custompb/proto/gen/debug && \
	protoc --debug_out="./custompb/proto/gen/debug:./custompb/proto/gen/debug" custompb/proto/interceptor/*.proto custompb/proto/helloworld/*.proto