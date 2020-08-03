package regional

import (
	data "github.com/piyuo/libsrv/data"
)

// Location represent single location
//
type Location struct {
	data.BaseObject `firestore:"-"`
}

// LocationTable return location table
//
//	table := db.LocationTable()
//
func (db *DB) LocationTable() *data.Table {

	return &data.Table{
		Connection: db.Connection,
		TableName:  "Location",
		Factory: func() data.Object {
			return &Location{}
		},
	}
}
