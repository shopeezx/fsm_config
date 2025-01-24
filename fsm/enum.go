/**
 * @Copyright: 2025 Shopee
 * @Author: xiang.zhong
 * @Email: xiang.zhong@shopee.com
 * @Date: 2025/01/24 15:38
 */
package fsm

// PaymentPayoutFsmStatus FSM中的状态值
type PaymentPayoutFsmStatus string

const (
	PaymentPayoutFsmStatusPendingReceiptGeneration                PaymentPayoutFsmStatus = "PendingBillReceiptGeneration"
	PaymentPayoutFsmStatusPendingInvoiceGeneration                PaymentPayoutFsmStatus = "PendingBillInvoiceGeneration"
	PaymentPayoutFsmStatusDuringPostback                          PaymentPayoutFsmStatus = "DuringPostback"
	PaymentPayoutFsmStatusPendingL2Pay                            PaymentPayoutFsmStatus = "PendingL2Pay"
	PaymentPayoutFsmStatusPendingInvoiceOrReceiptUpload           PaymentPayoutFsmStatus = "PendingInvoiceOrReceiptUpload"
	PaymentPayoutFsmStatusPendingInvoiceOrReceiptReview           PaymentPayoutFsmStatus = "PendingInvoiceOrReceiptReview"
	PaymentPayoutFsmStatusCompleted                               PaymentPayoutFsmStatus = "Completed"
	PaymentPayoutFsmStatusPendingL1Review                         PaymentPayoutFsmStatus = "PendingL1Review"
	PaymentPayoutFsmStatusPendingOfflinePaid                      PaymentPayoutFsmStatus = "PendingOfflinePaid"
	PaymentPayoutFsmStatusPendingAffiliateConfirmInvoiceOrReceipt PaymentPayoutFsmStatus = "PendingAffiliateConfirm"
	PaymentPayoutFsmStatusPendingRetry                            PaymentPayoutFsmStatus = "PendingRetry"
	PaymentPayoutFsmStatusTransferredToOfflinePendingPaid         PaymentPayoutFsmStatus = "TransferredToOfflinePendingPaid"
	PaymentPayoutFsmStatusPendingBank                             PaymentPayoutFsmStatus = "PendingBank"
	PaymentPayoutFsmStatusPendingShopeePay                        PaymentPayoutFsmStatus = "PendingShopeePay"
	PaymentPayoutFsmStatusPendingL2Review                         PaymentPayoutFsmStatus = "PendingL2Review"
	PaymentPayoutFsmStatusCanceled                                PaymentPayoutFsmStatus = "Canceled"
	PaymentPayoutFsmStatusClosed                                  PaymentPayoutFsmStatus = "Closed"
)

func (s PaymentPayoutFsmStatus) String() string {
	return string(s)
}

const (
	PaymentPayoutFsmEventInvoiceGenerated         = "PaymentPayoutFsmEventInvoiceGenerated"
	PaymentPayoutFsmEventSubmitInvoiceOrReceipt   = "PaymentPayoutFsmEventSubmitInvoiceOrReceipt"
	PaymentPayoutFsmEventPassInvoiceOrReceipt     = "PaymentPayoutFsmEventPassInvoiceOrReceipt"
	PaymentPayoutFsmEventDismissInvoiceOrReceipt  = "PaymentPayoutFsmEventDismissInvoiceOrReceipt"
	PaymentPayoutFsmEventPartitionPostback        = "PaymentPayoutFsmEventPartitionPostback"
	PaymentPayoutFsmEventFinishPostback           = "PaymentPayoutFsmEventFinishPostback"
	PaymentPayoutFsmEventPaid                     = "PaymentPayoutFsmEventPaid"
	PaymentPayoutFsmEventPayByShopeePayChannel    = "PaymentPayoutFsmEventPayByShopeePayChannel"
	PaymentPayoutFsmEventPayByBankTransferChannel = "PaymentPayoutFsmEventPayByBankTransferChannel"
	PaymentPayoutFsmEventPayoutSucceed            = "PaymentPayoutFsmEventPayoutSucceed"
	PaymentPayoutFsmEventTransferToOffline        = "PaymentPayoutFsmEventTransferToOffline"
)

const (
	CallbackPaymentPayoutFsmEventAffiliateCancel                   = "callbackPaymentPayoutFsmEventAffiliateCancel"
	CallbackPaymentPayoutFsmEventAutoRetryPayout                   = "callbackPaymentPayoutFsmEventAutoRetryPayout"
	CallbackPaymentPayoutFsmEventClose                             = "callbackPaymentPayoutFsmEventClose"
	CallbackPaymentPayoutFsmEventDismissInvoiceOrReceipt           = "callbackPaymentPayoutFsmEventDismissInvoiceOrReceipt"
	CallbackPaymentPayoutFsmEventFinishPostback                    = "callbackPaymentPayoutFsmEventFinishPostback"
	CallbackPaymentPayoutFsmEventInvoiceGenerated                  = "callbackPaymentPayoutFsmEventInvoiceGenerated"
	CallbackPaymentPayoutFsmEventL1FullyReview                     = "callbackPaymentPayoutFsmEventL1FullyReview"
	CallbackPaymentPayoutFsmEventL2FullyReview                     = "callbackPaymentPayoutFsmEventL2FullyReview"
	CallbackPaymentPayoutFsmEventManualRetryPayout                 = "callbackPaymentPayoutFsmEventManualRetryPayout"
	CallbackPaymentPayoutFsmEventOpsCancel                         = "callbackPaymentPayoutFsmEventOpsCancel"
	CallbackPaymentPayoutFsmEventPaid                              = "callbackPaymentPayoutFsmEventPaid"
	CallbackPaymentPayoutFsmEventPartitionPostback                 = "callbackPaymentPayoutFsmEventPartitionPostback"
	CallbackPaymentPayoutFsmEventPassInvoiceOrReceipt              = "callbackPaymentPayoutFsmEventPassInvoiceOrReceipt"
	CallbackPaymentPayoutFsmEventPayByShopeePayChannel             = "callbackPaymentPayoutFsmEventPayByShopeePayChannel"
	CallbackPaymentPayoutFsmEventPayByBankTransferChannel          = "callbackPaymentPayoutFsmEventPayByBankTransferChannel"
	CallbackPaymentPayoutFsmEventPayoutFailedToOfflinePaid         = "callbackPaymentPayoutFsmEventPayoutFailedToOfflinePaid"
	CallbackPaymentPayoutFsmEventPayoutFailedTransferToAutoRetry   = "callbackPaymentPayoutFsmEventPayoutFailedTransferToAutoRetry"
	CallbackPaymentPayoutFsmEventPayoutFailedTransferToManualRetry = "callbackPaymentPayoutFsmEventPayoutFailedTransferToManualRetry"
	CallbackPaymentPayoutFsmEventPayoutSucceed                     = "callbackPaymentPayoutFsmEventPayoutSucceed"
	CallbackPaymentPayoutFsmEventSubmitInvoiceOrReceipt            = "callbackPaymentPayoutFsmEventSubmitInvoiceOrReceipt"
	CallbackPaymentPayoutFsmEventSystemCancel                      = "callbackPaymentPayoutFsmEventSystemCancel"
	CallbackPaymentPayoutFsmEventTransferToOffline                 = "callbackPaymentPayoutFsmEventTransferToOffline"
	CallbackPaymentPayoutFsmEventUserDeletedToOfflinePaid          = "callbackPaymentPayoutFsmEventUserDeletedToOfflinePaid"
)
