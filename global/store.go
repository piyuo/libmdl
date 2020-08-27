package global

import (
	"github.com/piyuo/libsrv/data"
)

// Store represent single store, ID is global account id
//
type Store struct {
	data.BaseObject

	// AccountID show this store belong to which account
	//
	AccountID string

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

	// Name is store name
	//
	Name string

	// Domain is domain in piyuo.com, eg. example.piyuo.com, example is domain
	//
	Domain string

	// CustomDomain is custom domain name user defined, eg. cacake.com
	//
	CustomDomain string

	// Locations keep all locations
	//
	Locations map[string]string

	// Policy is Casbin Policy
	//
	Policy string

	// Roles keep custom roles
	//
	Roles map[string]string
}

// StoreTable return store table
//
//	table := db.StoreTable()
//
func (c *Global) StoreTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Store",
		Factory: func() data.Object {
			return &Store{}
		},
	}
}
