package mdl

// Payment is how user pay for service
//
type Payment int

const (
	// PaymentSubscription is pay by in-App subscription
	//
	PaymentSubscription Payment = 1

	// PaymentBill is send bill to user
	//
	PaymentBill Payment = 2
)
