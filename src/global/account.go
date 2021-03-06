package global

import (
	"context"
	"time"

	"github.com/piyuo/libsrv/src/data"
	"github.com/pkg/errors"
)

// Account represent account in piyuo.com, account can have many user and many store
//
type Account struct {
	data.DomainObject

	// Region datacenter used by this account
	//
	Region string

	// Locale is owner locale
	//
	Locale string

	// Timezone is store defult locale
	//
	Timezone string

	// TimezoneOffset is store defult locale
	//
	TimezoneOffset int

	// Plan is account servie plan
	//
	Plan AccountPlan

	// Currency is plan fee currency
	//
	Currency string

	// Plan is account servie plan
	//
	PlanServiceFee int64

	// BillDate of an existing contract is the date bill must be renewed
	//
	BillDate time.Time

	// RenewalDate of an existing contract is the date on which it must be renewed. every created account must have renewal date
	// RenewalDate will not update if account owner didn't pay
	// if RenewalDate is less than 6 month from now. the account will be remove
	//
	RenewalDate time.Time

	// PaymentMethod is how user pay for the service
	//
	PaymentMethod AccountPaymentMethod

	// Policy is Casbin Policy
	//
	Policy string

	// Roles keep custom roles
	//
	Roles map[string]string
}

// Status return account status base on renewal date
//
//	status := account.Status()
//
func (c *Account) Status() AccountStatus {
	safetyLine := time.Now().AddDate(0, 0, -1).UTC()
	if c.RenewalDate.After(safetyLine) {
		return AccountStatusOK
	}

	suspendDeadline := time.Now().AddDate(0, 0, -60).UTC()
	if c.RenewalDate.Before(suspendDeadline) {
		return AccountStatusSuspend
	}

	return AccountStatusNonRenewal
}

// SuspendDate return account suspend date base on renewal date
//
//	suspendDate := account.SuspendDate()
//
func (c *Account) SuspendDate() time.Time {
	return c.RenewalDate.AddDate(0, 0, 60)
}

// AccountTable return account table
//
//	table := db.AccountTable()
//
func (c *Global) AccountTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Account",
		Factory: func() data.Object {
			return &Account{}
		},
	}
}

// RemoveAllAccount remove all account
//
//	err := RemoveAllAccount(ctx)
//
func (c *Global) RemoveAllAccount(ctx context.Context) error {
	return c.AccountTable().Clear(ctx)
}

// GetAccountByID get store by account id
//
func (c *Global) GetAccountByID(ctx context.Context, accountID string) (*Account, error) {
	iAccount, err := c.AccountTable().Get(ctx, accountID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get store by accountID: "+accountID)
	}
	if iAccount == nil {
		return nil, nil // possible account already removed
	}
	return iAccount.(*Account), nil
}
