package mdl

// Object is Casbin object
//
type Object uint8

// global
const (

	// UserAdmin can read or write user
	//
	ObjectUserAdmin Object = 1

	// ProductAdmin can read or write product
	//
	ObjectProductAdmin Object = 2

	// LocationAdmin can read or write Location
	//
	ObjectLocationAdmin Object = 3
)

// per location
const (

	// LocationSetting can read or write setting in location
	//
	ObjectLocationSetting Object = 101

	// LocationMenu can read or write menu in location
	//
	ObjectLocationMenu Object = 102

	// LocationOrder can read or write order in location
	//
	ObjectLocationOrder Object = 6
)
