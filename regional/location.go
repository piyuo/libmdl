package regional

import (
	data "github.com/piyuo/libsrv/data"
)

// Location represent single location , ID is serial id to keep it short
//
type Location struct {
	data.BaseObject

	// AccountID show this store belong to which account
	//
	AccountID string

	// StoreID mean this location belong to which store
	//
	StoreID string

	// Name is location name
	//
	Name string

	// Status show store is open or closed
	//
	Status LocationStatus

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

// LocationTable return location table
//
//	table := regional.LocationTable()
//
func (c *Regional) LocationTable() *data.Table {

	return &data.Table{
		Connection: c.Connection,
		TableName:  "Location",
		Factory: func() data.Object {
			return &Location{}
		},
	}
}
