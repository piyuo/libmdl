package mdl

// PaymentMethod is how user pay for the service
//
type PaymentMethod int

const (
	// PaymentMethodNotSet user not set the payment method
	//
	PaymentMethodNotSet PaymentMethod = 0

	// PaymentMethodInAppSubscription is pay by in-App subscription
	//
	PaymentMethodInAppSubscription = 1

	// PaymentMethodSendBill is send bill to user
	//
	PaymentMethodSendBill PaymentMethod = 2
)
