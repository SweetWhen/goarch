package main

import (
	"github.com/SweetWhen/goarch/custompb/proto/gen/helloworld"
	"github.com/SweetWhen/goarch/sweetim/internal/logic/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	r := gin.Default()
	// use middleware...
	helloworld.RegisterGreeterHTTPServer(r, service.New())

	if err := r.Run(":9551"); err != nil {
		log.Fatalf("gin.Run err:%s\n", err.Error())
	}
}
