package regional

import (
	data "github.com/piyuo/libsrv/data"
)

// Counters is collection of global usage counters
//
type Counters struct {
	data.Counters `firestore:"-"`
}
