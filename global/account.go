package global

import (
	"time"

	"github.com/piyuo/libmdl/def"
	"github.com/piyuo/libsrv/data"
)

// Account represent single store, ID is serial id to keep it short
//
type Account struct {
	data.BaseObject

	// Email is owner email, indicate who own this account
	//
	Email string

	// FirstName is user first name
	//
	FirstName string

	// LastName is user last name
	//
	LastName string

	// StoreName name is user store name
	//
	StoreName string

	// Domain is domain in piyuo.com, eg. example.piyuo.com, example is domain
	//
	Domain string

	// CustomDomain is custom domain name user defined, eg. cacake.com
	//
	CustomDomain string

	// Region datacenter this account
	//
	Region string

	// Locale is owner locale
	//
	Locale string

	// Plan is account servie plan
	//
	Plan def.Plan

	// State account status
	//
	State def.State

	// Renewal is piyuo service renew date
	//
	Renewal time.Time

	// Payments how to pay bill
	//
	Payments def.Payments
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
