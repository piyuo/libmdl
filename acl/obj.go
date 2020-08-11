package acl

type obj uint8

const (
	// global

	// Users obj
	//
	Users obj = 1

	// Products obj
	//
	Products obj = 2

	// Locations obj
	//
	Locations obj = 3

	// per location

	// Location obj
	//
	Location obj = 4

	// Menu obj
	//
	Menu obj = 5

	// Order obj is in location
	//
	Order obj = 6
)
