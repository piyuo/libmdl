package regional

import (
	"github.com/piyuo/libsrv/src/data"
)

// Product represent product in store
//
type Product struct {
	data.DomainObject
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
