架构设计如下：
![image](https://github.com/user-attachments/assets/493743d5-2b30-4705-babf-840f1e8e4805)

支持灵活配置动态初始化FSM有限状态机基于事件驱动账单状态在各个审核阶段的处理操作，确保状态的转换符合预定义的业务逻辑。
```json
{
  "callbacks": {
    "PaymentPayoutFsmEventInvoiceGenerated": "callbackPaymentPayoutFsmEventInvoiceGenerated",
    "PaymentPayoutFsmEventPayByShopeePayChannel": "callbackPaymentPayoutFsmEventPayByShopeePayChannel",
    "PaymentPayoutFsmEventPayoutSucceed": "callbackPaymentPayoutFsmEventPayoutSucceed"
  },
  "events": {
    "PaymentPayoutFsmEventInvoiceGenerated": {
      "from": [
        "PendingReceiptGeneration",
        "PendingBillInvoiceGeneration"
      ],
      "to": "DuringPostback"
    },
    "PaymentPayoutFsmEventPayByShopeePayChannel": {
      "from": [
        "PendingL2Pay"
      ],
      "to": "PendingShopeePay"
    },
    "PaymentPayoutFsmEventPayoutSucceed": {
      "from": [
        "PendingShopeePay"
      ],
      "to": "Completed"
    }
  },
  "initial_state": "PendingBillInvoiceGeneration"
}

如配置如下状态机：
```
------------Test Br FSM------------ 
Event: PaymentPayoutFsmEventInvoiceGenerated, From: PendingBillInvoiceGeneration, To: DuringPostback
Event: PaymentPayoutFsmEventPayByShopeePayChannel, From: PendingL2Pay, To: PendingShopeePay
Error: event PaymentPayoutFsmEventPayByShopeePayChannel inappropriate in current state PendingInvoiceOrReceiptUpload
Event: PaymentPayoutFsmEventPayoutSucceed, From: PendingShopeePay, To: Completed

访问测试结果：
