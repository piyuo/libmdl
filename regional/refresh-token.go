package regional

import (
	data "github.com/piyuo/libsrv/data"
)

// RefreshToken let user login using refresh token
//
type RefreshToken struct {
	data.BaseObject

	// StoreID is store id
	//
	StoreID string

	// UserID is user id
	//
	UserID string

	// IP is user ip the token belong to, user can have multiple refresh token in different IP
	//
	IP string
}
