package regional

import (
	"github.com/piyuo/libsrv/data"
)

// Menu represent menu in location
//
type Menu struct {
	data.BaseObject
}

// MenuTable return Menu table
//
//	table := regional.MenuTable()
//
func (c *Regional) MenuTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Menu",
		Factory: func() data.Object {
			return &Menu{}
		},
	}
}
