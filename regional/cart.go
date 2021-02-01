package regional

import (
	"github.com/piyuo/libsrv/src/data"
)

// Cart represent Cart in store
//
type Cart struct {
	data.DomainObject
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
