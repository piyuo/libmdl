package regional

import (
	"github.com/piyuo/libsrv/data"
)

// Customer represent Customer in store
//
type Customer struct {
	data.BaseObject

	// AccountID show this record belong to which account
	//
	AccountID string
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
