package fsm

import (
	"context"

	"go.uber.org/zap"
)

type PayoutFsmBr struct {
	*PaymentFsm
}

func NewFsmBr(ctx context.Context, configFileUrl string, model PayoutModel) *PayoutFsmBr {
	logger := &zap.Logger{}
	config, err := LoadConfigFromFile(configFileUrl)
	if err != nil {
		panic(err)
	}
	fsm, err := NewFsm(ctx, config, logger)
	if err != nil {
		panic(err)
	}

	fsm.FSM.SetState(model.PayoutStatus.String())
	return &PayoutFsmBr{
		PaymentFsm: fsm,
	}
}
