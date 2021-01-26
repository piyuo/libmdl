package global

// AccountStatus is account status
//
type AccountStatus int8

const (
	// AccountStatusOK account is active and work normally
	//
	AccountStatusOK AccountStatus = 1

	// AccountStatusNonRenewal account didn't renew the plan
	//
	AccountStatusNonRenewal AccountStatus = -1

	// AccountStatusSuspend account didn't renew the plan for 60 days
	//
	AccountStatusSuspend AccountStatus = -2
)

// ErrorAccountNonRenewal account didn't renew the plan
//
const ErrorAccountNonRenewal = "ACCOUNT_NON_RENEWAL"

// ErrorAccountSuspend account didn't renew the plan for 60 days
//
const ErrorAccountSuspend = "ACCOUNT_SUSPEND"

// AccountStatusToError convert status to error code. return empty if nothing wrong
//
func AccountStatusToError(status AccountStatus) string {
	switch status {
	case AccountStatusNonRenewal:
		return ErrorAccountNonRenewal
	case AccountStatusSuspend:
		return ErrorAccountSuspend
	}
	return ""
}
