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

	// UserAgent when create this refresh token
	//
	UserAgent string

	// Expired time
	//
	Expired time.Time
}
