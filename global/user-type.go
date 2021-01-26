package global

// UserType is define user type in entire system
//
type UserType int32

const (
	// UserTypeOwner is account Owner
	//
	UserTypeOwner UserType = 1

	// UserTypeAdministrator is administrator role, it can access all store and location
	//
	UserTypeAdministrator UserType = 2

	// UserTypeManager is general manager role, it can access all store and location
	//
	UserTypeManager UserType = 3

	// UserTypeViewer is geeneral manager role, but can only view everything not edit
	//
	UserTypeViewer UserType = 4

	// UserTypeStaff is staff, it right specify in StoreRoles and LocationRoles
	//
	UserTypeStaff UserType = 5
)
