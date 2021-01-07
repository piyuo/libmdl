package global

// RoleLocation define user role in specifc location
//
type RoleLocation int8

const (
	// RoleLocationManager is manager role, it can access everything
	//
	RoleLocationManager RoleLocation = 0

	// RoleLocationViewer is manage viewer role, it can view everything
	//
	RoleLocationViewer = 1

	// RoleLocationCashier is cashier role
	//
	RoleLocationCashier = 2

	// RoleLocationKitchen is kitchen role
	//
	RoleLocationKitchen = 3

	// RoleLocationWaiter is waiter role
	//
	RoleLocationWaiter = 4

	// RoleLocationDelivery is delivery role
	//
	RoleLocationDelivery = 5

	// RoleLocationCustom1 custom role 1
	//
	RoleLocationCustom1 = 250

	// RoleLocationCustom2 custom role 2
	//
	RoleLocationCustom2 = 251

	// RoleLocationCustom3 custom role 3
	//
	RoleLocationCustom3 = 252

	// RoleLocationCustom4 custom role 4
	//
	RoleLocationCustom4 = 253

	// RoleLocationCustom5 custom role 5
	//
	RoleLocationCustom5 = 254
)
