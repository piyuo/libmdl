package mdl

// Status is account/user status
//
type Status uint8

const (
	// StatusNone is no status
	//
	StatusNone Status = 0

	// StatusAccountSuspend mean account not renew in time and wait for recycle
	//
	StatusAccountSuspend Status = 1

	// StatusActive mean account/user is active and work normally
	//
	StatusActive Status = 2

	// StatusUserResign mean user has been leave this account permanently
	//
	StatusUserResign Status = 3

	// StatusAccountCanceled mean accont has problem and has been canceled manually. this account will not recycle and close permanently
	//
	StatusAccountCanceled Status = 4

	// StatusUserCanceled mean user has problem and has been canceled manually. this user will not recycle and close permanently
	//
	StatusUserCanceled Status = 5
)
