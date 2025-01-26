package fsm

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"
)

func NewFsm(ctx context.Context, config *FsmConfig) (*PaymentFsm, error) {
	// 校验配置是否有效
	if config == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}
	if config.InitialState == "" {
		return nil, fmt.Errorf("initial state cannot be empty")
	}

	// 定义事件
	events := fsm.Events{}
	for eventName, eventConfig := range config.Events {
		events = append(events, fsm.EventDesc{
			Name: eventName,
			Src:  eventConfig.From, // 来源状态
			Dst:  eventConfig.To,   // 目标状态
		})
	}

	// 定义回调
	callbacks := fsm.Callbacks{}
	for eventName, callbackFunc := range config.Callbacks {
		callback := getCallbackByName(callbackFunc) // 根据名称获取回调
		if callback == nil {
			fmt.Printf("Callback function not found, callback: %s\n", callbackFunc)
			continue
		}
		callbacks[eventName] = callback
	}

	// 创建状态机实例
	fmsInstance := fsm.NewFSM(
		config.InitialState.String(), // 初始状态
		events,                       // 状态机事件
		callbacks,                    // 回调
	)

	// 返回包装的状态机
	return &PaymentFsm{
		FSM:     fmsInstance,
		Config:  config,
		Context: ctx,
	}, nil
}

var callbackRegistry = map[string]fsm.Callback{
	CallbackPaymentPayoutFsmEventAffiliateCancel: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventAffiliateCancel 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventAutoRetryPayout: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventAutoRetryPayout 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventClose: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventClose 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventDismissInvoiceOrReceipt: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventDismissInvoiceOrReceipt 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventFinishPostback: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventFinishPostback 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventInvoiceGenerated: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventInvoiceGenerated 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventL1FullyReview: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventL1FullyReview 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventL2FullyReview: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventL2FullyReview 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventManualRetryPayout: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventManualRetryPayout 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventOpsCancel: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventOpsCancel 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventPaid: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventPaid 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventPartitionPostback: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventPartitionPostback 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventPassInvoiceOrReceipt: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventPassInvoiceOrReceipt 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventPayByShopeePayChannel: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventPayByShopeePayChannel 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventPayByBankTransferChannel: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventPayByBankTransferChannel 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventPayoutFailedToOfflinePaid: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventPayoutFailedToOfflinePaid 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventPayoutFailedTransferToAutoRetry: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventPayoutFailedTransferToAutoRetry 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventPayoutFailedTransferToManualRetry: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventPayoutFailedTransferToManualRetry 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventPayoutSucceed: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventPayoutSucceed 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventSubmitInvoiceOrReceipt: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventSubmitInvoiceOrReceipt 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventSystemCancel: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventSystemCancel 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventTransferToOffline: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventTransferToOffline 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
	CallbackPaymentPayoutFsmEventUserDeletedToOfflinePaid: func(e *fsm.Event) {
		// 示例回调：处理 PaymentPayoutFsmEventUserDeletedToOfflinePaid 事件
		fmt.Printf("Event: %s, From: %s, To: %s\n", e.Event, e.Src, e.Dst)
	},
}

func getCallbackByName(eventCallback string) fsm.Callback {
	if callback, ok := callbackRegistry[eventCallback]; ok {
		return callback
	}
	fmt.Printf("Callback not found, eventCallback:%s\n", eventCallback)
	return nil
}

type PayoutOperationHandler interface {
	SysCreateInvoice(*PayoutModel) error
	Pay(*PayoutModel) error
	PayComplete(*PayoutModel) error
}

var payoutHandler PayoutOperationHandler

var PaymentPayoutHandlerRegionMap = map[string]PayoutOperationHandler{
	"br": &PayoutOperationHandlerBr{},
	"tw": &PayoutOperationHandlerTw{},
}
