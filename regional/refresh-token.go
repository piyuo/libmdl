package regional

import (
	"time"

	data "github.com/piyuo/libsrv/data"
)

// RefreshToken let user login using refresh token
//
type RefreshToken struct {
	data.BaseObject

	// AccountID is store id
	//
	AccountID string

	// UserID is user id
	//
	UserID string

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

// RefreshTokenTable return RefreshTokenTable table
//
//	table := db.RefreshTokenTable()
//
func (c *Regional) RefreshTokenTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "RefreshToken",
		Factory: func() data.Object {
			return &RefreshToken{}
		},
	}
}
