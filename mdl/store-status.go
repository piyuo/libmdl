package mdl

// StoreStatus is store status
//
type StoreStatus int8

const (
	// StoreOpen store is Open
	//
	StoreOpen StoreStatus = 1

	// StoreTemporarilyClosed store is temporarily closed
	//
	StoreTemporarilyClosed = 0

	// StorePermanentlyClosed store is permanently closed
	//
	StorePermanentlyClosed = -1
)

// ErrorStoreTemporarilyClosed store is temporarily closed
//
const ErrorStoreTemporarilyClosed = "STORE_TEMP_CLOSED"

// ErrorStoreClosed store is permanently closed
//
const ErrorStoreClosed = "STORE_CLOSED"

// StoreStatusToError convert status to error code. return empty if nothing wrong
//
func StoreStatusToError(status StoreStatus) string {
	switch status {
	case StoreTemporarilyClosed:
		return ErrorStoreTemporarilyClosed
	case StorePermanentlyClosed:
		return ErrorStoreClosed
	}
	return ""
}
