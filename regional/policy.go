package regional

import (
	"github.com/piyuo/libsrv/data"
)

// Policy represent a policy usd by an account
//
type Policy struct {
	data.BaseObject

	// Policy is Casbin Policy
	//
	Policy string

	// Roles keep custom roles
	//
	Roles map[string]string
}

// PolicyTable return policy table
//
//	table := db.PolicyTable()
//
func (c *Regional) PolicyTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Policy",
		Factory: func() data.Object {
			return &Policy{}
		},
	}
}
