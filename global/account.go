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

	// TimezoneName is store defult locale
	//
	TimezoneName string

	// TimezoneOffset is store defult locale
	//
	TimezoneOffset int

	// Email is owner email, indicate who own this account
	//
	Email string

	// FirstName is user first name
	//
	FirstName string

	// LastName is user last name
	//
	LastName string

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
