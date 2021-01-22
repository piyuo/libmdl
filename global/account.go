package global

import (
	"time"

	"github.com/piyuo/libsrv/data"
)

// Account represent account in piyuo.com, account can have many user and many store
//
type Account struct {
	data.BaseObject

	// Suspend set to true mean account not receve for long time. all store close and only can use renew service
	//
	Suspend bool

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

	//OwerID is owner user id
	//
	OwnerID string

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
