package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"io/ioutil"
	"testing"
)

func TestGenGin(t *testing.T) {
	data, err := ioutil.ReadFile("./proto/gen/debug/code_generator_request.pb.bin")
	if err != nil {
		t.Fatal(err)
	}
	req := &pluginpb.CodeGeneratorRequest{}
	if err = proto.Unmarshal(data, req); err != nil {
		t.Fatal(err)
	}

	opts := protogen.Options{}
	plugin, err := opts.New(req)
	if err = myGenPlugin(plugin); err != nil {
		plugin.Error(err)
	}
	response := plugin.Response()
	out, err := proto.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(out))
}
