package fsm

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type PayoutFsmBr struct {
	*PaymentFsm
}

func NewFsmBr(ctx context.Context, configFileUrl string) *PayoutFsmBr {
	config, err := LoadConfigFromFile(configFileUrl)
	if err != nil {
		panic(err)
	}
	fsm, err := NewFsm(ctx, config)
	if err != nil {
		panic(err)
	}
	return &PayoutFsmBr{
		PaymentFsm: fsm,
	}
}

type PayoutOperationHandlerBr struct {
}

func (handler *PayoutOperationHandlerBr) LockAndGetFsm(state PaymentPayoutFsmStatus) *PayoutFsmBr {
	fsm := NewFsmBr(context.Background(), "./conf/br_fsm.json")
	fsm.FSM.SetState(state.String())
	return fsm
}

func (handler *PayoutOperationHandlerBr) SysCreateInvoice(model *PayoutModel) error {
	fsmBr := handler.LockAndGetFsm(model.PayoutStatus)
	timeStart := time.Now()
	// 记录发票生成时间日志
	fmt.Printf("[PayoutFsmBr]%v:----1.发票生成完成----\n", time.Now())

	// 给运营发送发票生成邮件
	fmt.Printf("[PayoutFsmBr]%v:----2.生成发票通知完成----\n", time.Now())

	// 统计耗时
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Printf("[PayoutFsmBr]%v:----3.生成发票耗时%vs----\n", time.Now(), time.Since(timeStart).Seconds())
	return fsmBr.FSM.Event(PaymentPayoutFsmEventInvoiceGenerated)
}

func (handler *PayoutOperationHandlerBr) Pay(model *PayoutModel) error {
	fsmBr := handler.LockAndGetFsm(model.PayoutStatus)
	fmt.Printf("[PayoutFsmBr]%v:----发起Shopee Pay支付----\n", time.Now())
	return fsmBr.FSM.Event(PaymentPayoutFsmEventPayByShopeePayChannel)
}

func (handler *PayoutOperationHandlerBr) PayComplete(model *PayoutModel) error {
	fsmBr := handler.LockAndGetFsm(model.PayoutStatus)
	fmt.Printf("[PayoutFsmBr]%v:----支付完成----\n", time.Now())
	return fsmBr.FSM.Event(PaymentPayoutFsmEventPayoutSucceed)
}
