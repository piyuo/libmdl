package regional

import (
	"github.com/piyuo/libsrv/data"
)

// Product represent product in store
//
type Product struct {
	data.BaseObject

	// AccountID show this record belong to which account
	//
	AccountID string
}

// ProductTable return product table
//
//	table := regional.productTable()
//
func (c *Regional) ProductTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Product",
		Factory: func() data.Object {
			return &Product{}
		},
	}
}
