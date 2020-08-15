package regional

import (
	data "github.com/piyuo/libsrv/data"
)

// Serials keep all serial numbers
//
type Serials struct {
	data.Serials `firestore:"-"`
}
