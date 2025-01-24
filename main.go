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
	if err := brFsm.FSM.Event(fsm.PaymentPayoutFsmEventInvoiceGenerated); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	payoutModel.PayoutStatus = fsm.PaymentPayoutFsmStatusPendingL2Pay
	brFsm = fsm.NewFsmBr(context.Background(), "./conf/br_fsm.json", payoutModel)
	if err := brFsm.FSM.Event(fsm.PaymentPayoutFsmEventPayByShopeePayChannel); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	payoutModel.PayoutStatus = fsm.PaymentPayoutFsmStatusPendingInvoiceOrReceiptUpload
	brFsm = fsm.NewFsmBr(context.Background(), "./conf/br_fsm.json", payoutModel)
	if err := brFsm.FSM.Event(fsm.PaymentPayoutFsmEventPayByShopeePayChannel); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	payoutModel.PayoutStatus = fsm.PaymentPayoutFsmStatusPendingShopeePay
	brFsm = fsm.NewFsmBr(context.Background(), "./conf/br_fsm.json", payoutModel)
	if err := brFsm.FSM.Event(fsm.PaymentPayoutFsmEventPayoutSucceed); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// tw fsm
	fmt.Printf("------------Test Br FSM------------ \n")
	payoutModel.PayoutStatus = fsm.PaymentPayoutFsmStatusPendingInvoiceGeneration
	twFsm := fsm.NewFsmTw(context.Background(), "./conf/tw_fsm.json", payoutModel)
	if err := twFsm.FSM.Event(fsm.PaymentPayoutFsmEventInvoiceGenerated); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	payoutModel.PayoutStatus = fsm.PaymentPayoutFsmStatusPendingL2Pay
	twFsm = fsm.NewFsmTw(context.Background(), "./conf/tw_fsm.json", payoutModel)
	if err := twFsm.FSM.Event(fsm.PaymentPayoutFsmEventPayByBankTransferChannel); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	payoutModel.PayoutStatus = fsm.PaymentPayoutFsmStatusPendingL2Pay
	twFsm = fsm.NewFsmTw(context.Background(), "./conf/tw_fsm.json", payoutModel)
	if err := twFsm.FSM.Event(fsm.PaymentPayoutFsmEventPayByShopeePayChannel); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	payoutModel.PayoutStatus = fsm.PaymentPayoutFsmStatusPendingBank
	twFsm = fsm.NewFsmTw(context.Background(), "./conf/tw_fsm.json", payoutModel)
	if err := twFsm.FSM.Event(fsm.PaymentPayoutFsmEventPayoutSucceed); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
