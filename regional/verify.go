package regional

import (
	"time"

	"github.com/piyuo/libsrv/db"
)

// Verify keep log to prevent someone guess password
//
type Verify struct {
	db.Entity

	// Ip who do the verify
	//
	Ip string `firestore:"Ip,omitempty"`

	// Error count
	//
	Error uint32 `firestore:"Error,omitempty"`

	// Verify which email
	//
	Email string `firestore:"Email,omitempty"`

	// Lasttime user try verify and fail
	//
	Lasttime time.Time `firestore:"Lasttime"`
}

// Factory create a empty object, return object must be nil safe, no nil in any field
//
func (c *Verify) Factory() db.Object {
	return &Pin{}
}

// Collection return the name in database
//
func (c *Verify) Collection() string {
	return "Verify"
}
