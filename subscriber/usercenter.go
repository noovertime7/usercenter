package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	usercenter "usercenter/proto/usercenter"
)

type Usercenter struct{}

func (e *Usercenter) Handle(ctx context.Context, msg *usercenter.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *usercenter.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
