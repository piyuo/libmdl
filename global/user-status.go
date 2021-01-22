package global

// UserStatus is user status
//
type UserStatus int8

const (
	// UserStatusActive user is active and work normally
	//
	UserStatusActive UserStatus = 1

	// UserStatusSuspend user has been suspend
	//
	UserStatusSuspend = -1
)

// ErrorUserSuspend is user suspend error code to client
//
const ErrorUserSuspend = "USER_SUSPEND"

// UserStatusToError convert status to error code. return empty if nothing wrong
//
func UserStatusToError(status UserStatus) string {
	switch status {
	case UserStatusSuspend:
		return ErrorUserSuspend
	}
	return ""
}
