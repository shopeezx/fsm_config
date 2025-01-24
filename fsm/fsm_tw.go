/**
 * @Copyright: 2025 Shopee
 * @Author: xiang.zhong
 * @Email: xiang.zhong@shopee.com
 * @Date: 2025/01/24 15:42
 */
package fsm

import (
	"context"

	"go.uber.org/zap"
)

type PayoutFsmTw struct {
	*PaymentFsm
}

func NewFsmTw(ctx context.Context, configFileUrl string, model PayoutModel) *PayoutFsmTw {
	logger := (&zap.Logger{})
	// .With(zap.Any("model", model))
	config, err := LoadConfigFromFile(configFileUrl)
	if err != nil {
		panic(err)
	}
	fsm, err := NewFsm(ctx, config, logger)
	if err != nil {
		panic(err)
	}

	fsm.FSM.SetState(model.PayoutStatus.String())
	return &PayoutFsmTw{
		PaymentFsm: fsm,
	}
}
