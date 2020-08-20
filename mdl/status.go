package mdl

// Status is account/user status
//
type Status uint8

// ErrorAccountSuspend is account suspend error code to client
//
const ErrorAccountSuspend = "ACCOUNT_SUSPEND"

// ErrorAccountCanceled is account canceled error code to client
//
const ErrorAccountCanceled = "ACCOUNT_CANCELED"

// ErrorUserLeave is account suspend error code to client
//
const ErrorUserLeave = "USER_LEAVE"

// ErrorUserCanceled is user canceled error code to client
//
const ErrorUserCanceled = "USER_CANCELED"

const (
	// StatusNone is no status
	//
	StatusNone Status = 0

	// StatusActive mean account/user is active and work normally
	//
	StatusActive Status = 1

	// StatusAccountSuspend mean account not renew in time and wait for recycle
	//
	StatusAccountSuspend Status = 2

	// StatusAccountCanceled mean accont has problem and has been canceled manually. this account will not recycle and close permanently
	//
	StatusAccountCanceled Status = 3

	// StatusUserLeave mean user has been leave this account permanently
	//
	StatusUserLeave Status = 4

	// StatusUserCanceled mean user has problem and has been canceled manually. this user will not recycle and close permanently
	//
	StatusUserCanceled Status = 5
)

// StatusToError convert status to error code. return empty if nothing wrong
//
func StatusToError(status Status) string {
	switch status {
	case StatusAccountSuspend:
		return ErrorAccountSuspend
	case StatusAccountCanceled:
		return ErrorAccountCanceled
	case StatusUserLeave:
		return ErrorUserLeave
	case StatusUserCanceled:
		return ErrorUserCanceled
	}
	return ""
}
