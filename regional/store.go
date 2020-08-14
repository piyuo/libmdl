package regional

import (
	"github.com/piyuo/libsrv/data"
)

// Store represent single store, ID is global account id
//
type Store struct {
	data.BaseObject

	// Name is store name
	//
	Name string

	//CustomRoles keep custom roles
	//
	CustomRoles map[string]string

	//CustomGroups keep custom roles
	//
	CustomGroups map[string]string
}

// StoreTable return store table
//
func (c *Regional) StoreTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Store",
		Factory: func() data.Object {
			return &Store{}
		},
	}
}
