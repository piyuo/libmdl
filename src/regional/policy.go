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

func (c *Policy) Factory() db.Object {
	return &Policy{}
}

func (c *Policy) Collection() string {
	return "Policy"
}
