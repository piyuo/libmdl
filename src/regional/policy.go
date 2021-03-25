package regional

import "github.com/piyuo/libsrv/src/db"

// Policy represent a policy usd by an account
//
type Policy struct {
	db.Model

	// Policy is Casbin Policy
	//
	Policy string `firestore:"Policy,omitempty"`

	// Roles keep custom roles
	//
	Roles map[string]string `firestore:"Roles,omitempty"`
}

// Factory create a empty object, return object must be nil safe, no nil in any field
//
func (c *Policy) Factory() db.Object {
	return &Policy{}
}

// Collection return the name in database
//
func (c *Policy) Collection() string {
	return "Policy"
}
