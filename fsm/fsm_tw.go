package fsm

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type PayoutFsmTw struct {
	*PaymentFsm
}

func NewFsmTw(ctx context.Context, configFileUrl string) *PayoutFsmTw {
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

type PayoutOperationHandlerTw struct {
}

func (handler *PayoutOperationHandlerTw) LockAndGetFsm(state PaymentPayoutFsmStatus) *PayoutFsmTw {
	fsm := NewFsmTw(context.Background(), "./conf/tw_fsm.json")
	fsm.FSM.SetState(state.String())
	return fsm
}

func (handler *PayoutOperationHandlerTw) SysCreateInvoice(model *PayoutModel) error {
	fsmTw := handler.LockAndGetFsm(model.PayoutStatus)
	timeStart := time.Now()
	// 记录发票生成时间日志
	fmt.Printf("[PayoutFsmTw]%v:----1.发票生成完成----\n", time.Now())

	// 给运营发送发票生成邮件
	// fmt.Printf("[PayoutFsmTw]%v:----生成发票通知完成----\n", time.Now())

	// 统计耗时
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Printf("[PayoutFsmTw]%v:----2.生成发票耗时%vs----\n", time.Now(), time.Since(timeStart).Seconds())
	return fsmTw.FSM.Event(PaymentPayoutFsmEventInvoiceGenerated)
}

func (handler *PayoutOperationHandlerTw) Pay(model *PayoutModel) error {
	fsmTw := handler.LockAndGetFsm(model.PayoutStatus)
	fmt.Printf("[PayoutFsmTw]%v:----发起银行卡支付----\n", time.Now())
	return fsmTw.FSM.Event(PaymentPayoutFsmEventPayByBankTransferChannel)
}

func (handler *PayoutOperationHandlerTw) PayComplete(model *PayoutModel) error {
	fsmTw := handler.LockAndGetFsm(model.PayoutStatus)
	fmt.Printf("[PayoutFsmTw]%v:----支付完成----\n", time.Now())
	return fsmTw.FSM.Event(PaymentPayoutFsmEventPayoutSucceed)
}
