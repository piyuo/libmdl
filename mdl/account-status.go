package mdl

// AccountStatus is account status
//
type AccountStatus int8

const (
	// AccountActive account is active
	//
	AccountActive AccountStatus = 1

	// AccountSuspend account not renew in time and wait for recycle
	//
	AccountSuspend = -1

	// AccountCanceled accont has problem and has been canceled manually, this account will not recycle and close permanently
	//
	AccountCanceled = -2
)

// ErrorAccountSuspend is account suspend error code to client
//
const ErrorAccountSuspend = "ACCOUNT_SUSPEND"

// ErrorAccountCanceled is account canceled error code to client
//
const ErrorAccountCanceled = "ACCOUNT_CANCELED"

// AccountStatusToError convert status to error code. return empty if nothing wrong
//
func AccountStatusToError(status AccountStatus) string {
	switch status {
	case AccountSuspend:
		return ErrorAccountSuspend
	case AccountCanceled:
		return ErrorAccountCanceled
	}
	return ""
}
