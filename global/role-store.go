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

	// RoleStoreCustom1 custom role 1
	//
	RoleStoreCustom1 = 250

	// RoleStoreCustom2 custom role 2
	//
	RoleStoreCustom2 = 251

	// RoleStoreCustom3 custom role 3
	//
	RoleStoreCustom3 = 252

	// RoleStoreCustom4 custom role 4
	//
	RoleStoreCustom4 = 253

	// RoleStoreCustom5 custom role 5
	//
	RoleStoreCustom5 = 254
)
