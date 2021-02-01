package regional

import (
	data "github.com/piyuo/libsrv/src/data"
)

// Counters is collection of regional usage counters
//
type Counters struct {
	data.Counters `firestore:"-"`
}
