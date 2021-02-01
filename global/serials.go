package global

import (
	data "github.com/piyuo/libsrv/src/data"
)

// Serials keep all serial numbers
//
type Serials struct {
	data.Serials `firestore:"-"`
}
