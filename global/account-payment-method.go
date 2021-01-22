package global

// AccountPaymentMethod is how user pay for the service
//
type AccountPaymentMethod int

const (
	// AccountPaymentMethodNotSet is user not set the payment method
	//
	AccountPaymentMethodNotSet AccountPaymentMethod = 0

	// AccountPaymentMethodInAppSubscription is pay by in-App subscription
	//
	AccountPaymentMethodInAppSubscription = 1

	// AccountPaymentMethodWebSubscription is pay by web subscription
	//
	AccountPaymentMethodWebSubscription = 2
)
