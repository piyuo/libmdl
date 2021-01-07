package global

// UserType is define user role in entire system
//
type UserType int8

const (
	// UserTypeOwner is owner role, it can access everything
	//
	UserTypeOwner UserType = 0

	// UserTypeAdministrator is administrator role, it can access all store and location but not bill relate
	//
	UserTypeAdministrator = 1

	// UserTypeManager is general manager role, it can access all store and location
	//
	UserTypeManager = 2

	// UserTypeViewer is geeneral manager role, but can only view everything not edit
	//
	UserTypeViewer = 3

	// UserTypeStaff is staff, it right specify in StoreRoles and LocationRoles
	//
	UserTypeStaff = 4
)
