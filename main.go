package main

import (
	"context"
	"fmt"
	"time"

	"fsm_config/fsm"
)

func main() {
	payoutModel := fsm.PayoutModel{
		Id:           100,
		Operator:     "system",
		CreateTime:   time.Now().AddDate(0, 0, -1).Unix(),
		UpdateTime:   time.Now().Unix(),
		PayoutStatus: fsm.PaymentPayoutFsmStatusPendingInvoiceGeneration,
	}
	// br fsm
	fmt.Printf("------------Test Br FSM------------ \n")
	brFsm := fsm.NewFsmBr(context.Background(), "./conf/br_fsm.json", payoutModel)
	brFsm.FSM.SetState(fsm.PaymentPayoutFsmStatusPendingInvoiceGeneration.String())
	if err := brFsm.FSM.Event(fsm.PaymentPayoutFsmEventInvoiceGenerated); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("fms current state is %s\n", brFsm.FSM.Current())

	brFsm = fsm.NewFsmBr(context.Background(), "./conf/br_fsm.json", payoutModel)
	brFsm.FSM.SetState(fsm.PaymentPayoutFsmStatusPendingL2Pay.String())
	if err := brFsm.FSM.Event(fsm.PaymentPayoutFsmEventPayByShopeePayChannel); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("fms current state is %s\n", brFsm.FSM.Current())

	brFsm = fsm.NewFsmBr(context.Background(), "./conf/br_fsm.json", payoutModel)
	brFsm.FSM.SetState(fsm.PaymentPayoutFsmStatusPendingInvoiceOrReceiptUpload.String())
	if err := brFsm.FSM.Event(fsm.PaymentPayoutFsmEventPayByShopeePayChannel); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("fms current state is %s\n", brFsm.FSM.Current())

	brFsm = fsm.NewFsmBr(context.Background(), "./conf/br_fsm.json", payoutModel)
	brFsm.FSM.SetState(fsm.PaymentPayoutFsmStatusPendingShopeePay.String())
	if err := brFsm.FSM.Event(fsm.PaymentPayoutFsmEventPayoutSucceed); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("fms current state is %s\n", brFsm.FSM.Current())

	// tw fsm
	fmt.Printf("------------Test Br FSM------------ \n")
	twFsm := fsm.NewFsmTw(context.Background(), "./conf/tw_fsm.json", payoutModel)
	twFsm.FSM.SetState(fsm.PaymentPayoutFsmStatusPendingInvoiceGeneration.String())
	if err := twFsm.FSM.Event(fsm.PaymentPayoutFsmEventInvoiceGenerated); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("fms current state is %s\n", twFsm.FSM.Current())

	twFsm = fsm.NewFsmTw(context.Background(), "./conf/tw_fsm.json", payoutModel)
	twFsm.FSM.SetState(fsm.PaymentPayoutFsmStatusPendingL2Pay.String())
	if err := twFsm.FSM.Event(fsm.PaymentPayoutFsmEventPayByBankTransferChannel); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("fms current state is %s\n", twFsm.FSM.Current())

	twFsm = fsm.NewFsmTw(context.Background(), "./conf/tw_fsm.json", payoutModel)
	twFsm.FSM.SetState(fsm.PaymentPayoutFsmStatusPendingL2Pay.String())
	if err := twFsm.FSM.Event(fsm.PaymentPayoutFsmEventPayByShopeePayChannel); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("fms current state is %s\n", twFsm.FSM.Current())

	twFsm = fsm.NewFsmTw(context.Background(), "./conf/tw_fsm.json", payoutModel)
	twFsm.FSM.SetState(fsm.PaymentPayoutFsmStatusPendingBank.String())
	if err := twFsm.FSM.Event(fsm.PaymentPayoutFsmEventPayoutSucceed); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("fms current state is %s\n", twFsm.FSM.Current())
}
