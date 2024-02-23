package service

import (
	"context"
	"fmt"
	"github.com/SweetWhen/goarch/custompb/proto/gen/helloworld"
	"google.golang.org/grpc/metadata"
	"log"
)

type IMSvr struct {
	helloworld.UnimplementedGreeterServer
}

func New() *IMSvr {
	s := &IMSvr{}
	return s
}

func (svr *IMSvr) SayHello(ctx context.Context, req *helloworld.HelloRequest) (resp *helloworld.HelloReply, err error) {
	metadata.FromIncomingContext(ctx)
	log.Printf("SayHello get req: %s\n", req.String())
	resp = &helloworld.HelloReply{Message: "hai, " + req.Name}
	return resp, nil
}

func (svr *IMSvr) SayHello2(ctx context.Context, req *helloworld.HelloRequest) (resp *helloworld.HelloReply, err error) {
	return nil, fmt.Errorf("hard code 500")
}
