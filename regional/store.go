package regional

import (
	data "github.com/piyuo/libsrv/data"
)

// Store represent single store, ID is global account id
//
type Store struct {
	data.BaseObject

	// Name is store name
	//
	Name string

	// SubDomain is sub domain in piyuo.com, eg. example.piyuo.com, example is sub domain
	//
	SubDomain string

	// CustomDomain is custom domain name user defined, eg. cacake.com
	//
	CustomDomain string

	// Plan is account servie plan
	//
	Plan int

	// Status account status
	//
	Status int

	// Locale is store main locale
	//
	Locale string
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
