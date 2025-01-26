package fsm

import (
	"context"
)

type PayoutFsmBr struct {
	*PaymentFsm
}

func NewFsmBr(ctx context.Context, configFileUrl string, model PayoutModel) *PayoutFsmBr {
	config, err := LoadConfigFromFile(configFileUrl)
	if err != nil {
		panic(err)
	}
	fsm, err := NewFsm(ctx, config)
	if err != nil {
		panic(err)
	}

	fsm.FSM.SetState(model.PayoutStatus.String())
	return &PayoutFsmBr{
		PaymentFsm: fsm,
	}
}
