package mdl

// UserStatus is user status
//
type UserStatus int8

const (
	// UserStatusActive user is active and work normally
	//
	UserStatusActive UserStatus = 1

	// UserStatusLeave user has been leave this account permanently
	//
	UserStatusLeave = 0

	// UserStatusCanceled store mean user has problem and has been canceled manually
	//
	UserStatusCanceled = -1
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
	case UserStatusLeave:
		return ErrorUserLeave
	case UserStatusCanceled:
		return ErrorUserCanceled
	}
	return ""
}
