package comm

import (
	data "github.com/piyuo/libsrv/data"
)

// User represent single user, ID is serial id to keep it short
//
type User struct {
	data.BaseObject

	// AccountID indicate user belong to which account
	//
	AccountID string

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

	// Groups is user belong which location group
	//
	Groups []string

	// Roles is user belong which policy role
	//
	Roles []string
}
