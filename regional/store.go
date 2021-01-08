package regional

import (
	"github.com/piyuo/libsrv/data"
)

// Store represent store in a location, ID is serial id to keep it short
//
type Store struct {
	data.BaseObject

	// AccountID show this store belong to which account
	//
	AccountID string

	// Name is store name
	//
	Name string

	// Status show store is open or closed
	//
	Status BusinessStatus

	// Domain is domain in piyuo.com, eg. example.piyuo.com, example is domain
	//
	Domain string

	// CustomDomain is custom domain name user defined, eg. cacake.com
	//
	CustomDomain string
}

// StoreTable return store table
//
//	table := regional.StoreTable()
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
