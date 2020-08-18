package namespace

import (
	"github.com/piyuo/libmdl/shared"
	"github.com/piyuo/libsrv/data"
)

// StoreTable return store table
//
//	table := db.StoreTable()
//
func (c *Namespace) StoreTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Store",
		Factory: func() data.Object {
			return &shared.Store{}
		},
	}
}
