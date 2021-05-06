package global

// RoleStore define user role in specifc store
//
type RoleStore int32

const (
	// RoleStoreManager is manager role, it can access everything
	//
	RoleStoreManager RoleStore = 0

	// RoleStoreViewer is manage viewer role, it can view everything
	//
	RoleStoreViewer = 1

	// RoleStoreProduct product manager
	//
	RoleStoreProduct = 2

	// RoleStoreOrder order manager
	//
	RoleStoreOrder = 3

	// RoleStoreOperation operation manager
	//
	RoleStoreOperation = 4

)
