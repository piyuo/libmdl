package regional

import (
	"github.com/piyuo/libsrv/db"
)

// Pin keep verification code
//
type Pin struct {
	db.Entity

	// Hash is code hash with salt, we do not store code only hash is enough
	//
	Hash uint32 `firestore:"Hash,omitempty"`

	// Crypted code
	//
	Crypted string `firestore:"Crypted,omitempty"`
}

// Factory create a empty object, return object must be nil safe, no nil in any field
//
func (c *Pin) Factory() db.Object {
	return &Pin{}
}

// Collection return the name in database
//
func (c *Pin) Collection() string {
	return "Verify"
}
