package global

import (
	"time"

	data "github.com/piyuo/libsrv/data"
)

// Account represent single account
//
type Account struct {
	data.BaseObject `firestore:"-"`

	// OwnerEmail is user id, indicate who own this account
	//
	OwnerEmail string

	// StoreName name is user store name
	//
	StoreName string

	// SubDomain is sub domain in piyuo.com, eg. example.piyuo.com, example is sub domain
	//
	SubDomain string

	// CustomDomain is custom domain name user defined, eg. cacake.com
	//
	CustomDomain string

	// CustomDomain1 is custom domain name user defined, eg. cacake.com
	//
	CustomDomain1 string

	// CustomDomain2 is custom domain name user defined, eg. cacake.com
	//
	CustomDomain2 string

	// RenewalDate is piyuo service renew date
	//
	RenewalDate time.Time

	// Plan is account servie plan
	//
	Plan int

	// Status account status
	//
	Status int

	// PaymentType account payment type
	//
	PaymentType int

	// Category is user choose category when they create account, category is used for generate web template
	//
	Category string
}

// AccountTable return account table
//
//	table := db.AccountTable()
//
func (db *DB) AccountTable() *data.Table {
	return &data.Table{
		Connection: db.Connection,
		TableName:  "Account",
		Factory: func() data.Object {
			return &Account{}
		},
	}
}

// Plan is piyuo service plan, used in account
//
type Plan int

const (
	// Free Plan
	//
	Free Plan = 0

	// Basic Plan
	//
	Basic Plan = 1

	// Premium Plan
	//
	Premium Plan = 2

	// Business Plan
	//
	Business Plan = 3

	// APP Plan
	//
	APP Plan = 4

	// Chain Plan
	//
	Chain Plan = 5

	// Source Plan
	//
	Source Plan = 6
)

// Status is piyuo service status
//
type Status int

const (
	// Active mean account is active and work normally
	//
	Active Status = 1

	// Pending mean account not renew in time and wait for recycle
	//
	Pending Status = 0

	// Canceled mean accont has problem and has been canceled manually. this account will not recycle and close permanently
	//
	Canceled Status = -1
)

// PaymentType is how user pay for service
//
type PaymentType int

const (
	// Subscription is pay by in-App subscription
	//
	Subscription PaymentType = 1

	// Bill is send bill to user
	//
	Bill PaymentType = 0
)
