package regional

import (
	"github.com/piyuo/libmdl/mdl"
	"github.com/piyuo/libsrv/data"
)

// Store represent store in a location, ID is serial id to keep it short
//
type Store struct {
	data.BaseObject

	// AccountID show this store belong to which account
	//
	AccountID string

	// Name is location name
	//
	Name string

	// Status show store is open or closed
	//
	Status mdl.StoreStatus

	// Country is location country
	//
	Country string

	// State is location state
	//
	State string

	// City is location city
	//
	City string

	// PostalCode is location postalCode
	//
	PostalCode string

	// AddressLine1 is location AddressLine1
	//
	AddressLine1 string

	// AddressLine2 is location AddressLine2
	//
	AddressLine2 string

	// Coordinate is location coordinate
	//
	Coordinate string

	// PhoneNumber is location phone number
	//
	PhoneNumber string

	// Hours is location hours
	//
	//	"mon":"24hr" // 24 hours
	//	"tue":"" // close
	//	"wed":"13:00-14:00"
	//	"thu":"07:00-21:00"
	//	"fri":"07:00-21:00"
	//	"sat":"07:00-21:00"
	//	"sun":"07:00-21:00"
	//
	Hours map[string]string
}

// StoreTable return store table
//
//	table := regional.StoreTable()
//
func (c *Regional) StoreTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Store",
		Factory: func() data.Object {
			return &Store{}
		},
	}
}
