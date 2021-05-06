package global

import (
	"github.com/piyuo/libsrv/db"
)

// Domain keep all registered domain name
//
type Domain struct {
	db.Model
}

// Factory create a empty object, return object must be nil safe, no nil in any field
//
func (c *Domain) Factory() db.Object {
	return &Domain{}
}

// Collection return the name in database
//
func (c *Domain) Collection() string {
	return "Domain"
}
