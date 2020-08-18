package shared

import "github.com/piyuo/libsrv/data"

// Store represent single store, ID is global account id
//
type Store struct {
	data.BaseObject

	// AccountID show this store belong to which account
	//
	AccountID string

	// Name is store name
	//
	Name string

	// Domain is domain in piyuo.com, eg. example.piyuo.com, example is domain
	//
	Domain string

	// CustomDomain is custom domain name user defined, eg. cacake.com
	//
	CustomDomain string

	// Locations keep all locations
	//
	Locations map[string]string

	// Policy is Casbin Policy
	//
	Policy string

	// Roles keep custom roles
	//
	Roles map[string]string
}
