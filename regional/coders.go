package regional

import (
	data "github.com/piyuo/libsrv/data"
)

// Coders keep all coders
//
type Coders struct {
	data.Coders `firestore:"-"`
}
