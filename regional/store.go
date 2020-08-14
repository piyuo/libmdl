package regional

import (
	"github.com/piyuo/libmdl/mdl"
	"github.com/piyuo/libsrv/data"
)

// StoreTable return store table
//
//	table := db.StoreTable()
//
func (c *Regional) StoreTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Store",
		Factory: func() data.Object {
			return &mdl.Store{}
		},
	}
}
