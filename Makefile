
.PHONY: idl
idl:
	@bash gen.sh

.PHONY: pb-debug
pb-debug:
	@rm -rf custompb/proto/gen/debug && mkdir -p custompb/proto/gen/debug && \
	protoc --debug_out="./custompb/proto/gen/debug:./custompb/proto/gen/debug" custompb/proto/interceptor/*.proto custompb/proto/helloworld/*.proto

.PHONY: sweet-gin
sweet-gin:
	@cd custompb && go build -o protoc-gen-sweet-gin && mv protoc-gen-sweet-gin /home/agan/code/go/bin/

.PHONY: urlrouter
urlrouter:
	@docker run -p 10000:10000  --rm -d --name envoy \
	-v `pwd`/sweetim/cmd/urlrouter/envoy.yaml:/etc/envoy/envoy.yaml \
	-v `pwd`/sweetim/cmd/urlrouter/share:/share \
	envoyproxy/envoy-alpine:v1.20.2