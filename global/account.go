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

	// Region datacenter used by this account
	//
	Region string

	// Locale is owner locale
	//
	Locale string

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

	// Status account/user status
	//
	Status mdl.Status

	// Renewal is piyuo service renew date
	//
	Renewal time.Time

	// PaymentMethod is how user pay for the service
	//
	PaymentMethod mdl.PaymentMethod
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
