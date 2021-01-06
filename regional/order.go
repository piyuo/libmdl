package regional

import (
	"github.com/piyuo/libsrv/data"
)

// Order represent Order in location
//
type Order struct {
	data.BaseObject

	// AccountID show this record belong to which account
	//
	AccountID string
}

// OrderTable return Order table
//
//	table := regional.OrderTable()
//
func (c *Regional) OrderTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Order",
		Factory: func() data.Object {
			return &Order{}
		},
	}
}
