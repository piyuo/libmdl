package mdl

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

	// Policy is Casbin Policy
	//
	Policy string

	//CustomRoles keep custom roles
	//
	CustomRoles map[string]string

	//CustomGroups keep custom roles
	//
	CustomGroups map[string]string
}
