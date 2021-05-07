package global

import (
	"context"
	"time"

	"github.com/piyuo/libsrv/db"
	"github.com/pkg/errors"
)

// Account can have many user and many store
//
type Account struct {
	db.Model

	// Status user status
	//
	Suspend bool `firestore:"Suspend,omitempty"`

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

	// RenewalDate of an existing contract is the date on which it must be renewed. every created account must have renewal date
	// it will not update if account owner didn't pay
	// if less than 6 month from now. the account will be remove
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

// GetAccountByID get store by account id
//
func GetAccountByID(ctx context.Context, accountID string) (*Account, error) {
	client, err := Client(ctx)
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
