package global

// UserType is define user type in entire system
//
type UserType int32

const (
	// UserTypeOwner is account Owner
	//
	UserTypeOwner = 1

	// UserTypeAdministrator is administrator role, it can access all store and location
	//
	UserTypeAdministrator = 2

	// UserTypeManager is general manager role, it can access all store and location
	//
	UserTypeManager = 3

	// UserTypeViewer is geeneral manager role, but can only view everything not edit
	//
	UserTypeViewer = 4

	// UserTypeStaff is staff, it right specify in StoreRoles and LocationRoles
	//
	UserTypeStaff = 5
)
