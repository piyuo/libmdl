package regional

import (
	"github.com/piyuo/libsrv/data"
)

// Location represent a store location
//
type Location struct {
	data.DomainObject

	// StoreID belong to which store
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

	// Email is location contact email
	//
	Email string

	// Wechat is location contact Wechat
	//
	Wechat string

	// Facebook is location contact Facebook
	//
	Facebook string

	// Hours is location hours
	//
	//	"mon":"all" // 24 hours
	//	"tue":"closed" // close
	//	"wed":"13001400"
	//	"thu":"07002100"
	//	"fri":"07002100"
	//	"sat":"07002100"
	//	"sun":"07002100"
	//
	Hours map[string]string

	// Timezone name for this location
	//
	Timezone string

	// Timezone offset for this location
	//
	TimezoneOffset int
}

// LocationTable return location table
//
//	table := regional.locationTable()
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
