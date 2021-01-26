package regional

import (
	"context"

	"github.com/piyuo/libsrv/data"
	"github.com/pkg/errors"
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

// RemoveAllStore remove all store
//
//	err := RemoveAllStore(ctx)
//
func (c *Regional) RemoveAllStore(ctx context.Context) error {
	return c.StoreTable().Clear(ctx)
}

// IsDomainTaken return true if domain already taken
//
//	taken, err := IsDomainTaken(ctx, "a@b.c")
//
func (c *Regional) IsDomainTaken(ctx context.Context, domain string) (bool, error) {
	storeTable := c.StoreTable()

	isExist, err := storeTable.Query().Where("Domain", "==", domain).IsExist(ctx)
	if err != nil {
		return false, errors.Wrap(err, "failed to execute query")
	}
	return isExist, nil
}
