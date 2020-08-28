package mdl

// UserStatus is user status
//
type UserStatus int8

const (
	// UserActive user is active and work normally
	//
	UserActive UserStatus = 1

	// UserLeave user has been leave this account permanently
	//
	UserLeave = 0

	// UserCanceled store mean user has problem and has been canceled manually
	//
	UserCanceled = -1
)

// ErrorUserLeave is account suspend error code to client
//
const ErrorUserLeave = "USER_LEAVE"

// ErrorUserCanceled is user canceled error code to client
//
const ErrorUserCanceled = "USER_CANCELED"

// UserStatusToError convert status to error code. return empty if nothing wrong
//
func UserStatusToError(status UserStatus) string {
	switch status {
	case UserLeave:
		return ErrorUserLeave
	case UserCanceled:
		return ErrorUserCanceled
	}
	return ""
}
