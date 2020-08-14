package acl

type obj uint8

// global
const (

	// UserAdmin can read or write user
	//
	UserAdmin obj = 1

	// ProductAdmin can read or write product
	//
	ProductAdmin obj = 2

	// LocationAdmin can read or write Location
	//
	LocationAdmin obj = 3
)

// per location
const (

	// LocationSetting can read or write setting in location
	//
	LocationSetting obj = 101

	// LocationMenu can read or write menu in location
	//
	LocationMenu obj = 102

	// LocationOrder can read or write order in location
	//
	LocationOrder obj = 6
)
