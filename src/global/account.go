package global

import (
	"context"
	"time"

	"github.com/piyuo/libsrv/src/db"
	"github.com/pkg/errors"
)

// Account can have many user and many store
//
type Account struct {
	db.Model

	// Region datacenter used by this account
	//
	Region string `firestore:"Region,omitempty"`

	// Locale is owner locale
	//
	Locale string `firestore:"Locale,omitempty"`

	// Timezone is store defult locale
	//
	Timezone string `firestore:"Timezone,omitempty"`

	// TimezoneOffset is store defult locale
	//
	TimezoneOffset int `firestore:"TimezoneOffset,omitempty"`

	// Plan is account servie plan
	//
	Plan AccountPlan `firestore:"Plan,omitempty"`

	// Currency is plan fee currency
	//
	Currency string `firestore:"Currency,omitempty"`

	// Plan is account servie plan
	//
	PlanServiceFee int64 `firestore:"PlanServiceFee,omitempty"`

	// BillDate of an existing contract is the date bill must be renewed
	//
	BillDate time.Time `firestore:"BillDate,omitempty"`

	// RenewalDate of an existing contract is the date on which it must be renewed. every created account must have renewal date
	// RenewalDate will not update if account owner didn't pay
	// if RenewalDate is less than 6 month from now. the account will be remove
	//
	RenewalDate time.Time `firestore:"RenewalDate,omitempty"`

	// PaymentMethod is how user pay for the service
	//
	PaymentMethod AccountPaymentMethod `firestore:"PaymentMethod,omitempty"`

	// Policy is Casbin Policy
	//
	Policy string `firestore:"Policy,omitempty"`

	// Roles keep custom roles
	//
	Roles map[string]string `firestore:"Roles,omitempty"`
}

// Factory create a empty object, return object must be nil safe, no nil in any field
//
func (c *Account) Factory() db.Object {
	return &Account{
		Roles: map[string]string{},
	}
}

// Collection return the name in database
//
func (c *Account) Collection() string {
	return "Account"
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

// GetAccountByID get store by account id
//
func GetAccountByID(ctx context.Context, accountID string) (*Account, error) {
	client, err := GlobalClient(ctx)
	if err != nil {
		return nil, err
	}

	iAccount, err := client.Get(ctx, &Account{}, accountID)
	if err != nil {
		return nil, errors.Wrap(err, "get account "+accountID)
	}
	if iAccount == nil {
		return nil, nil // possible account already removed
	}
	return iAccount.(*Account), nil
}
