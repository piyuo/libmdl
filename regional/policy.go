package regional

import (
	"github.com/piyuo/libsrv/data"
)

// Policy is Casbin Policy
//
type Policy struct {
	data.BaseObject

	// CSV is Casbin Policy CSV
	//
	CSV string
}

// PolicyTable return policy table
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
