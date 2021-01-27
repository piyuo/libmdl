package regional

import (
	"context"

	"github.com/piyuo/libsrv/data"
)

// Store represent store in a location, ID is serial id to keep it short
//
type Store struct {
	data.DomainObject

	// Name is store name
	//
	Name string

	// Status show store is open or closed
	//
	Status StoreStatus

	// DomainName is full domain name, eg.  example.piyuo.com
	//
	DomainName string

	// CustomDomain is custom domain name user defined, eg. www.cacake.com
	//
	CustomDomainName string
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

// RemoveAllStore remove all store
//
//	err := RemoveAllStore(ctx)
//
func (c *Regional) RemoveAllStore(ctx context.Context) error {
	return c.StoreTable().Clear(ctx)
}
