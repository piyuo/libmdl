package mdl

// State is piyuo service status
//
type State int

const (
	// StateActive mean account is active and work normally
	//
	StateActive State = 1

	// StatePending mean account not renew in time and wait for recycle
	//
	StatePending State = 0

	// StateCanceled mean accont has problem and has been canceled manually. this account will not recycle and close permanently
	//
	StateCanceled State = -1
)
