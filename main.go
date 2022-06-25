package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"usercenter/handler"
	"usercenter/subscriber"

	userservice "gitee.com/noovertime/usercenter/proto/usercenter"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.usercenter"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	usercenter.RegisterUsercenterHandler(service.Server(), new(handler.Usercenter))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.usercenter", service.Server(), new(subscriber.Usercenter))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
