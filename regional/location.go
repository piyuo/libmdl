package regional

import (
	"github.com/piyuo/libsrv/db"
)

// Location represent a store location
//
type Location struct {
	db.Model

	// StoreID belong to which store
	//
	StoreID string `firestore:"StoreID,omitempty"`

	// Name is location name
	//
	Name string `firestore:"Name,omitempty"`

	// Status show store is open or closed
	//
	Status LocationStatus `firestore:"Status,omitempty"`

	// Country is location country code like 'US', 'CN', 'AU'
	//
	Country string `firestore:"Country,omitempty"`

	// Address is location's address
	//
	Address string `firestore:"Address,omitempty"`

	// Address is location's optional part like floor/room/building
	//
	Address2 string `firestore:"Address2,omitempty"`

	// Lat is location coordinate
	//
	Lat float64 `firestore:"Lat,omitempty"`

	// Lng is location coordinate
	//
	Lng float64 `firestore:"Lng,omitempty"`

	// Tags is location tags
	//
	Tags []string `firestore:"Tags,omitempty"`

	// PhoneNumber is location phone number
	//
	PhoneNumber string `firestore:"PhoneNumber,omitempty"`

	// Email is location contact email
	//
	Email string `firestore:"Email,omitempty"`

	// Wechat is location contact Wechat
	//
	Wechat string `firestore:"Wechat,omitempty"`

	// Facebook is location contact Facebook
	//
	Facebook string `firestore:"Facebook,omitempty"`

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
	Hours map[string]string `firestore:"Hours,omitempty"`

	// Timezone name for this location
	//
	Timezone string `firestore:"Timezone,omitempty"`

	// Timezone offset for this location
	//
	TimezoneOffset int `firestore:"TimezoneOffset,omitempty"`
}

// Factory create a empty object, return object must be nil safe, no nil in any field
//
func (c *Location) Factory() db.Object {
	return &Location{
		Hours: map[string]string{},
	}
}

// Collection return the name in database
//
func (c *Location) Collection() string {
	return "Location"
}
