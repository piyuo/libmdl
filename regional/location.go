package rmdl

import (
	data "github.com/piyuo/libsrv/data"
)

// Location represent single location
//
type Location struct {
	data.Object `firestore:"-"`
}

// LocationTable return location table
//
//	counter := db.LocationTable()
//
func (db *DB) LocationTable() *data.Table {

	return &data.Table{
		Connection: db.Connection,
		TableName:  "location",
		Factory: func() data.ObjectRef {
			return &Location{}
		},
	}
}
