package global

import (
	"time"

	data "github.com/piyuo/libsrv/data"
)

// Verify keep verify code for email verify
//
type Verify struct {
	data.BaseObject `firestore:"-"`

	// Email is where verify code sent
	//
	Email string

	// Code is verify code
	//
	Code string

	// Time is latest time the code was sent
	//
	Time time.Time
}

// VerifyTable return Verify object
//
//	table := db.VerifyTable()
//
func (db *DB) VerifyTable() *data.Table {
	return &data.Table{
		Connection: db.Connection,
		TableName:  "Verify",
		Factory: func() data.Object {
			return &Verify{}
		},
	}
}
