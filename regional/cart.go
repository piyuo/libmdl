package regional

import (
	"github.com/piyuo/libsrv/data"
)

// Cart represent Cart in store
//
type Cart struct {
	data.BaseObject
}

// CartTable return Cart table
//
//	table := regional.CartTable()
//
func (c *Regional) CartTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Cart",
		Factory: func() data.Object {
			return &Cart{}
		},
	}
}
