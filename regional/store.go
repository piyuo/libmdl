package rmdl

import (
	data "github.com/piyuo/libsrv/data"
)

// Store represent single store, store id is the same with global account id
//
type Store struct {
	data.Object `firestore:"-"`

	// StoreName name is user store name
	//
	StoreName string

	// SubDomain is sub domain in piyuo.com, eg. example.piyuo.com, example is sub domain
	//
	SubDomain string

	// CustomDomain is custom domain name user defined, eg. cacake.com
	//
	CustomDomain string

	// CustomDomain1 is custom domain name user defined, eg. cacake.com
	//
	CustomDomain1 string

	// CustomDomain2 is custom domain name user defined, eg. cacake.com
	//
	CustomDomain2 string
}

// StoreTable return store table
//
func (db *DB) StoreTable() *data.Table {
	return &data.Table{
		Connection: db.Connection,
		TableName:  "store",
		Factory: func() data.ObjectRef {
			return &Store{}
		},
	}
}
