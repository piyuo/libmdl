package mdl

// AccountStatus is account status
//
type AccountStatus int8

const (
	// AccountStatusActive account is active
	//
	AccountStatusActive AccountStatus = 1

	// AccountStatusSuspend account not renew in time and wait for recycle
	//
	AccountStatusSuspend = -1

	// AccountStatusCanceled accont has problem and has been canceled manually, this account will not recycle and close permanently
	//
	AccountStatusCanceled = -2
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
	case AccountStatusSuspend:
		return ErrorAccountSuspend
	case AccountStatusCanceled:
		return ErrorAccountCanceled
	}
	return ""
}
