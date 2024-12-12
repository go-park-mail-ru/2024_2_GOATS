package dto

// CreatePaymentData for create_payment action
type CreatePaymentData struct {
	SubscriptionID uint64
	Amount         uint64
}
