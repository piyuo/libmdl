package global

// UserRole is define user role in entire system
//
type UserRole int8

const (
	// UserRoleOwner is owner role, it can access everything
	//
	UserRoleOwner UserRole = 0

	// UserRoleAdministrator is administrator role, it can access all store and location but not bill relate
	//
	UserRoleAdministrator = 1

	// UserRoleManager is general manager role, it can access all store and location
	//
	UserRoleManager = 2

	// UserRoleViewer is geeneral manager role, but can only view everything not edit
	//
	UserRoleViewer = 3

	// UserRoleStaff is staff, it right specify in StoreRoles and LocationRoles
	//
	UserRoleStaff = 4
)
