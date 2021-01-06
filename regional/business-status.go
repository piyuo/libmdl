package regional

// BusinessStatus is store or location status
//
type BusinessStatus int8

const (
	// BusinessOpen is store or location open
	//
	BusinessOpen BusinessStatus = 1

	// BusinessTemporarilyClosed is store or location temporarily closed
	//
	BusinessTemporarilyClosed = 0

	// BusinessClosed is store or location permanently closed
	//
	BusinessClosed = -1
)

// ErrorBusinessTemporarilyClosed is store or location temporarily closed
//
const ErrorBusinessTemporarilyClosed = "BUSINESS_TEMP_CLOSED"

// ErrorBusinessClosed is store or location permanently closed
//
const ErrorBusinessClosed = "BUSINESS_CLOSED"

// BusinessStatusToError convert status to error code. return empty if nothing wrong
//
func BusinessStatusToError(status BusinessStatus) string {
	switch status {
	case BusinessTemporarilyClosed:
		return ErrorBusinessTemporarilyClosed
	case BusinessClosed:
		return ErrorBusinessClosed
	}
	return ""
}
