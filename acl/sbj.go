package acl

// sbj is predefine role
//
type sbj uint8

const (

	// Owner role
	//
	Owner sbj = 1

	// Administrator role
	//
	Administrator sbj = 2

	// ProductManager role
	//
	ProductManager sbj = 3

	// OperatoionManager role
	//
	OperatoionManager sbj = 4
)

const (

	// LocationManager role
	//
	LocationManager sbj = 101

	// LocationOperator role
	//
	LocationOperator sbj = 102
)
