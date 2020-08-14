package mdl

// Subject is Casbin subject, set of predefine role
//
type Subject uint8

const (

	// SubjectOwner is owner role
	//
	SubjectOwner Subject = 1

	// SubjectAdministrator is administrator role
	//
	SubjectAdministrator Subject = 2

	// SubjectProductManager is product manager role
	//
	SubjectProductManager Subject = 3

	// SubjectOperatoionManager is operation manager role
	//
	SubjectOperatoionManager Subject = 4
)

const (

	// SubjectLocationManager  is location manager role
	//
	SubjectLocationManager Subject = 101

	// SubjectLocationOperator is location operator role
	//
	SubjectLocationOperator Subject = 102
)
