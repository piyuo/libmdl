package global

// RoleLocation define user role in specifc location
//
type RoleLocation int32

const (
	// RoleLocationManager is manager role, it can access everything
	//
	RoleLocationManager RoleLocation = 0

	// RoleLocationViewer is manage viewer role, it can view everything
	//
	RoleLocationViewer = 1

	// RoleLocationFrontdesk is frontdesk role
	//
	RoleLocationFrontdesk = 2

	// RoleLocationWaiter is waiter role
	//
	RoleLocationWaiter = 3

	// RoleLocationKitchen is kitchen role
	//
	RoleLocationKitchen = 4

	// RoleLocationCashier is cashier role
	//
	RoleLocationCashier = 5

	// RoleLocationDelivery is delivery role
	//
	RoleLocationDelivery = 6
)
