package global

import (
	"time"

	"github.com/piyuo/libmdl/mdl"
	"github.com/piyuo/libsrv/data"
)

// Account represent single store, ID is serial id to keep it short
//
type Account struct {
	data.BaseObject

	// Status account status
	//
	Status mdl.AccountStatus

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

	// UserID is owner user id, indicate which user own this account
	//
	UserID string

	// Plan is account servie plan
	//
	Plan mdl.Plan

	// Renewal is piyuo service renew date
	//
	Renewal time.Time

	// PaymentMethod is how user pay for the service
	//
	PaymentMethod mdl.PaymentMethod

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
