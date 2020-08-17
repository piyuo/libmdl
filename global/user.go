package global

import (
	"github.com/piyuo/libsrv/data"
)

// User represent single user, ID is serial id to keep it short
//
type User struct {
	data.BaseObject

	// AccountID indicate user belong to which account
	//
	AccountID string

	// Region meas user belong to which data center
	//
	Region string

	// Email is unique in User table, user need use email to login to their store
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
	LastName string

	// LocationID mean user belong to which location
	//
	LocationID string

	// Roles is user belong which policy role
	//
	Roles []string
}

// UserTable return user table
//
//	table := db.UserTable()
//
func (c *Global) UserTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "User",
		Factory: func() data.Object {
			return &User{}
		},
	}
}
