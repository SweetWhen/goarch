package main

import (
	"context"
	"fmt"
	svrPb "github.com/SweetWhen/goarch/custompb/proto/gen/helloworld"
	ciOpt "github.com/SweetWhen/goarch/custompb/proto/gen/interceptor"
	"log"
	"net"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

var handlers = &struct {
	Methods map[string]*ciOpt.MethodHandler
	Services map[string]*ciOpt.ServiceHandler
}{
	Methods:make(map[string]*ciOpt.MethodHandler),
	Services: make(map[string]*ciOpt.ServiceHandler),
}

func main()  {

	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		services := fd.Services()
		for i := 0; i < services.Len(); i++ {
			service := services.Get(i)
			if serviceHandler, _ := proto.GetExtension(service.Options(), ciOpt.E_ServiceHandler).(*ciOpt.ServiceHandler); serviceHandler != nil {
				fmt.Println()
				fmt.Println("--------service--------")
				fmt.Println("service name:" + string(service.FullName()))
				handlers.Services[string(service.FullName())] = serviceHandler

				if serviceHandler.Authorization != nil && *serviceHandler.Authorization != "" {
					fmt.Println("use interceptor authorization: " + *serviceHandler.Authorization)
				}
				fmt.Println("------- service -------")
			}
			methods := service.Methods()
			for k := 0; k < methods.Len(); k++ {
				method := methods.Get(k)
				if methodHandler, _ := proto.GetExtension(method.Options(), ciOpt.E_MethodHandler).(*ciOpt.MethodHandler); methodHandler != nil{
					fmt.Println()
					fmt.Println("--------method--------")
					fmt.Println("method name:" + string(method.FullName()))

					if methodHandler.Whitelist != nil && *methodHandler.Whitelist != "" {
						fmt.Println("use interceptor whitelist: " + *methodHandler.Whitelist )
					}
					if methodHandler.Logger != nil {
						fmt.Println("use interceptor logger: " + strconv.FormatBool(*methodHandler.Logger))
					}
					handlers.Methods[fmt.Sprintf("/%s/%s", string(service.FullName()), string(method.Name()))] = methodHandler
					fmt.Println("------- method -------")
				}

			}

		}

		return true
	})

	svrOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			fullMethod := strings.Split(info.FullMethod, "/")
			svrname := fullMethod[1]
			fmt.Printf("info.FullMethod:%s, fullMethod:%v, svrname:%s\n", info.FullMethod, fullMethod, svrname)
			serviceHandler := handlers.Services[svrname]
			fmt.Println("svrHandler[", svrname, "]:", serviceHandler.String())
			methodHandler := handlers.Methods[info.FullMethod]
			fmt.Println("methodHandler[", info.FullMethod, "]:", methodHandler.String())
			return handler(ctx, req)
		}),
	}
	
	lis, err := net.Listen("tcp", ":8877")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer(svrOpts...)
	svrPb.RegisterGreeterServer(s, &Service{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal("serve err", err)
		}
	}()

	conn, err := grpc.Dial("localhost:8877", grpc.WithInsecure())
	if err != nil {
		log.Fatal("dial err: ",err)
	}
	defer conn.Close()
	c := svrPb.NewGreeterClient(conn)
	_, err = c.SayHello(context.TODO(), &svrPb.HelloRequest{Name: "whensweet"})
	if err != nil {
		log.Fatal("sayHello err", err)
	}
	_, err = c.SayHello2(context.Background(), &svrPb.HelloRequest{Name: "whensweet hello2"})
	if err != nil {
		log.Fatal("sayHello2 err: ", err)
	}

}

type Service struct {
	svrPb.UnimplementedGreeterServer
}

func (s *Service) SayHello(ctx context.Context, req *svrPb.HelloRequest) (resp *svrPb.HelloReply, err error) {
	log.Println("SayHello...", req)
	return &svrPb.HelloReply{}, nil
}

func (s *Service) SayHello2(ctx context.Context, req *svrPb.HelloRequest) (resp *svrPb.HelloReply, err error) {
	log.Println("SayHello2...", req)
	return &svrPb.HelloReply{}, nil
}

 