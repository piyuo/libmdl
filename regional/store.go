package regional

import (
	"github.com/piyuo/libsrv/db"
)

// Store represent store in a location, ID is serial id to keep it short
//
type Store struct {
	db.Model

	// Name is store name
	//
	Name string `firestore:"Name,omitempty"`

	// Status show store is open or closed
	//
	Status StoreStatus `firestore:"Status,omitempty"`

	// DomainName is full domain name, eg.  example.piyuo.com
	//
	DomainName string `firestore:"DomainName,omitempty"`

	// CustomDomain is custom domain name user defined, eg. www.cacake.com
	//
	CustomDomainName string `firestore:"CustomDomainName,omitempty"`
}

// Factory create a empty object, return object must be nil safe, no nil in any field
//
func (c *Store) Factory() db.Object {
	return &Store{}
}

// Collection return the name in database
//
func (c *Store) Collection() string {
	return "Store"
}
