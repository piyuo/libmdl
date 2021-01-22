package regional

// LocationStatus is location status
//
type LocationStatus int8

const (
	// LocationStatusOpen is location open
	//
	LocationStatusOpen LocationStatus = 1

	// LocationStatusClosed is location permanently closed
	//
	LocationStatusClosed = -1

	// LocationStatusMaintenance is location in maintenance
	//
	LocationStatusMaintenance = -2

	// LocationStatusTemporarilyClosed is location temporarily closed
	//
	LocationStatusTemporarilyClosed = -3
)

// ErrorLocationClosed is location permanently closed
//
const ErrorLocationClosed = "LOCATION_CLOSED"

// ErrorLocationTemporarilyClosed is location temporarily closed
//
const ErrorLocationTemporarilyClosed = "LOCATION_TEMP_CLOSED"

// ErrorLocationMaintenance is location in maintenance
//
const ErrorLocationMaintenance = "LOCATION_MAINTENANCE"

// LocationStatusToError convert status to error code. return empty if nothing wrong
//
func LocationStatusToError(status LocationStatus) string {
	switch status {
	case LocationStatusClosed:
		return ErrorLocationClosed
	case LocationStatusMaintenance:
		return ErrorLocationMaintenance
	case LocationStatusTemporarilyClosed:
		return ErrorLocationTemporarilyClosed
	}
	return ""
}
