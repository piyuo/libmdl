package global

import (
	"time"

	data "github.com/piyuo/libsrv/data"
)

// Account represent single store, ID is serial id to keep it short
//
type Account struct {
	data.BaseObject

	// OwnerEmail is owner email, indicate who own this account
	//
	OwnerEmail string

	// FirstName is user first name
	//
	FirstName string

	// LastName is user last name
	//
	Lastname string

	// Region datacenter this account
	//
	Region string

	// Locale is owner locale
	//
	Locale string

	// StoreName name is user store name
	//
	StoreName string

	// SubDomain is sub domain in piyuo.com, eg. example.piyuo.com, example is sub domain
	//
	SubDomain string

	// CustomDomain is custom domain name user defined, eg. cacake.com
	//
	CustomDomain string

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
}

// AccountTable return account table
//
//	table := db.AccountTable()
//
func (c *Global) AccountTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
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
