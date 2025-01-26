package fsm

import (
	"context"
)

type PayoutFsmTw struct {
	*PaymentFsm
}

func NewFsmTw(ctx context.Context, configFileUrl string, model PayoutModel) *PayoutFsmTw {
	config, err := LoadConfigFromFile(configFileUrl)
	if err != nil {
		panic(err)
	}
	fsm, err := NewFsm(ctx, config)
	if err != nil {
		panic(err)
	}
	return &PayoutFsmTw{
		PaymentFsm: fsm,
	}
}
