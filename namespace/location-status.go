package namespace

// LocationStatus is location status
//
type LocationStatus int

const (
	// LocationStatusOpen mark location is Open
	//
	LocationStatusOpen LocationStatus = 1

	// LocationStatusTemporarilyClosed mark location temporarily closed
	//
	LocationStatusTemporarilyClosed LocationStatus = 0

	// LocationStatusPermanentlyClosed is permanently closed
	//
	LocationStatusPermanentlyClosed LocationStatus = -1
)
