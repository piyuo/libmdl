package regional

import (
	"github.com/piyuo/libsrv/data"
)

// Customer represent Customer in store
//
type Customer struct {
	data.DomainObject
}

// CustomerTable return Customer table
//
//	table := regional.CustomerTable()
//
func (c *Regional) CustomerTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Customer",
		Factory: func() data.Object {
			return &Customer{}
		},
	}
}
