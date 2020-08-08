package regional

import (
	"github.com/piyuo/libmdl/def"
	"github.com/piyuo/libsrv/data"
)

// Store represent single store, ID is global account id
//
type Store struct {
	data.BaseObject

	// Name is store name
	//
	Name string

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
}

// StoreTable return store table
//
func (c *Regional) StoreTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Store",
		Factory: func() data.Object {
			return &Store{}
		},
	}
}
