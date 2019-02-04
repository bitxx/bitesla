package main

import (
	"github.com/jason-wj/bitesla/service/service-algorithm/handler"
	"github.com/jason-wj/bitesla/service/service-algorithm/subscriber"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	example "github.com/jason-wj/bitesla/service/service-algorithm/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("bitesla.srv.algorithm"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("bitesla.srv.algorithm", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("bitesla.srv.algorithm", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
