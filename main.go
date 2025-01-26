package main

import (
	"fmt"
	"time"

	"fsm_config/fsm"
)

func main() {
	payoutModel := &fsm.PayoutModel{
		Id:           100,
		Operator:     "system",
		CreateTime:   time.Now().AddDate(0, 0, -1).Unix(),
		UpdateTime:   time.Now().Unix(),
		PayoutStatus: fsm.PaymentPayoutFsmStatusPendingInvoiceGeneration,
	}
	// br fsm
	fmt.Printf("------------Test Br FSM------------ \n")
	handler := fsm.PaymentPayoutHandlerRegionMap["br"]
	// 测试验证发票生成
	payoutModel.PayoutStatus = fsm.PaymentPayoutFsmStatusPendingInvoiceGeneration
	if err := handler.SysCreateInvoice(payoutModel); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	// 测试验证出款
	payoutModel.PayoutStatus = fsm.PaymentPayoutFsmStatusPendingL2Pay
	if err := handler.Pay(payoutModel); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	// 测试验证完成
	payoutModel.PayoutStatus = fsm.PaymentPayoutFsmStatusPendingShopeePay
	if err := handler.PayComplete(payoutModel); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// tw fsm
	fmt.Printf("------------Test tw FSM------------ \n")
	handler = fsm.PaymentPayoutHandlerRegionMap["tw"]
	// 测试验证发票生成
	payoutModel.PayoutStatus = fsm.PaymentPayoutFsmStatusPendingInvoiceGeneration
	if err := handler.SysCreateInvoice(payoutModel); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	// 测试验证出款
	payoutModel.PayoutStatus = fsm.PaymentPayoutFsmStatusPendingL2Pay
	if err := handler.Pay(payoutModel); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	// 测试验证完成
	payoutModel.PayoutStatus = fsm.PaymentPayoutFsmStatusPendingBank
	if err := handler.PayComplete(payoutModel); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

}
