package mdl

// PaymentMethod is how user pay for the service
//
type PaymentMethod int

const (
	// PaymentMethodSubscription is pay by in-App subscription
	//
	PaymentMethodSubscription PaymentMethod = 1

	// PaymentMethodBill is send bill to user
	//
	PaymentMethodBill PaymentMethod = 2
)
