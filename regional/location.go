package regional

import (
	data "github.com/piyuo/libsrv/data"
)

// Location represent single location , ID is serial id to keep it short
//
type Location struct {
	data.BaseObject

	// Name is location name
	//
	Name string
}

// LocationTable return location table
//
//	table := regional.LocationTable()
//
func (c *Regional) LocationTable() *data.Table {

	return &data.Table{
		Connection: c.Connection,
		TableName:  "Location",
		Factory: func() data.Object {
			return &Location{}
		},
	}
}
