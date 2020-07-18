package rmdl

import (
	data "github.com/piyuo/libsrv/data"
)

// User represent single user
//
type User struct {
	data.Object `firestore:"-"`

	// StoreID indicate user belong to which store, storID is equal to id in account table
	//
	StoreID string

	// Email is unique in account, user need use email or mobile to login to their store
	//
	Email string

	// BackupEmail used when user can't access their email service, they can choose send email to BackupEmail
	//
	BackupEmail string

	// Mobile is unique in account, user need use email or mobile to login to their store
	//
	Mobile string

	// FirstName is user first name
	//
	FirstName string

	// LastName is user last name
	//
	Lastname string
}

// UserTable return user table
//
//	counter := db.UserTable()
//
func (db *DB) UserTable() *data.Table {
	return &data.Table{
		Connection: db.Connection,
		TableName:  "user",
		Factory: func() data.ObjectRef {
			return &User{}
		},
	}
}
