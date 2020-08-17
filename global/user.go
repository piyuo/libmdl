package global

import (
	"time"

	"github.com/piyuo/libsrv/data"
)

// RefreshToken let user login using refresh token
//
type RefreshToken struct {

	// IP is user ip the token belong to, user can have multiple refresh token in different IP
	//
	IP string

	// Agent is user agent id from request user agent
	//
	Agent string

	// Expired time
	//
	Expired time.Time
}

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

	// Token keep all refresh token id for search
	//
	Tokens []string

	// RefreshTokens keep issued RefreshToken
	//
	RefreshTokens map[string]string
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
