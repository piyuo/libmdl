package comm

// Plan is piyuo service plan, used in account
//
type Plan int

const (
	// Free Plan
	//
	Free Plan = 0

	// Basic Plan
	//
	Basic Plan = 1

	// Premium Plan
	//
	Premium Plan = 2

	// Business Plan
	//
	Business Plan = 3

	// APP Plan
	//
	APP Plan = 4

	// Chain Plan
	//
	Chain Plan = 5

	// Source Plan
	//
	Source Plan = 6
)

// State is piyuo service status
//
type State int

const (
	// Active mean account is active and work normally
	//
	Active State = 1

	// Pending mean account not renew in time and wait for recycle
	//
	Pending State = 0

	// Canceled mean accont has problem and has been canceled manually. this account will not recycle and close permanently
	//
	Canceled State = -1
)

// Payments is how user pay for service
//
type Payments int

const (
	// Subscription is pay by in-App subscription
	//
	Subscription Payments = 1

	// Bill is send bill to user
	//
	Bill Payments = 2
)
