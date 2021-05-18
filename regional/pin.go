package regional

import (
	"github.com/piyuo/libsrv/db"
)

// Pin keep verification code
//
type Pin struct {
	db.Entity

	// ID is Email

	// Hash is code hash with salt, we do not store code only hash is enough
	//
	Hash uint32 `firestore:"Hash,omitempty"`

	// Crypted code
	//
	Crypted string `firestore:"Crypted,omitempty"`

	// Send is the code send history, it is time=ip mapping
	//
	Send map[string]string `firestore:"Send,omitempty"`

	// Enter is the code enter history, it is time=ip mapping
	//
	Enter map[string]string `firestore:"Enter,omitempty"`
}

// Factory create a empty object, return object must be nil safe, no nil in any field
//
func (c *Pin) Factory() db.Object {
	return &Pin{}
}

// Collection return the name in database
//
func (c *Pin) Collection() string {
	return "Pin"
}
