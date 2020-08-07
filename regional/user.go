package regional

import (
	data "github.com/piyuo/libsrv/data"
)

// User represent single user, ID is global user id
//
type User struct {
	data.BaseObject

	// StoreID indicate user belong to which store, storID is equal to id in account table
	//
	StoreID string

	// Email is unique in account, user need use email or mobile to login to their store
	//
	Email string

	// BackupEmail used when user can't access their email service, they can choose send email to BackupEmail
	//
	BackupEmail string

	// FirstName is user first name
	//
	FirstName string

	// LastName is user last name
	//
	Lastname string
}

// UserTable return user table
//
//	table := db.UserTable()
//
func (c *Regional) UserTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "User",
		Factory: func() data.Object {
			return &User{}
		},
	}
}
