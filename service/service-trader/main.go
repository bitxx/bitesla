package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/jason-wj/bitesla/service/service-trader/handler"
	"github.com/jason-wj/bitesla/service/service-trader/subscriber"

	example "github.com/jason-wj/bitesla/service/service-trader/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("bitesla.srv.trader"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("bitesla.srv.trader", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("bitesla.srv.trader", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
