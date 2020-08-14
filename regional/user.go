package regional

import (
	"github.com/piyuo/libmdl/mdl"
	"github.com/piyuo/libsrv/data"
)

// UserTable return user table
//
//	table := db.UserTable()
//
func (c *Regional) UserTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "User",
		Factory: func() data.Object {
			return &mdl.User{}
		},
	}
}
