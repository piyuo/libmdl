package regional

// StoreStatus is store status
//
type StoreStatus int8

const (
	// StoreStatusOpen is store open
	//
	StoreStatusOpen StoreStatus = 1

	// StoreStatusClosed is store permanently closed
	//
	StoreStatusClosed = -1

	// StoreStatusMaintenance is store in maintenance
	//
	StoreStatusMaintenance = -2

	// StoreStatusTemporarilyClosed is store temporarily closed
	//
	StoreStatusTemporarilyClosed = -3
)

// ErrorStoreClosed is store permanently closed
//
const ErrorStoreClosed = "STORE_CLOSED"

// ErrorStoreTemporarilyClosed is store temporarily closed
//
const ErrorStoreTemporarilyClosed = "STORE_TEMP_CLOSED"

// ErrorStoreMaintenance is store in maintenance
//
const ErrorStoreMaintenance = "STORE_MAINTENANCE"

// StoreStatusToError convert status to error code. return empty if nothing wrong
//
func StoreStatusToError(status StoreStatus) string {
	switch status {
	case StoreStatusClosed:
		return ErrorStoreClosed
	case StoreStatusMaintenance:
		return ErrorStoreMaintenance
	case StoreStatusTemporarilyClosed:
		return ErrorStoreTemporarilyClosed
	}
	return ""
}
