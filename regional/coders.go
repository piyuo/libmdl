package regional

import (
	data "github.com/piyuo/libsrv/data"
)

// Coders keep all coders
//
type Coders struct {
	data.Coders `firestore:"-"`
}

// LocationID return location id coder
//
//	coder := regional.LocationID()
//
func (c *Coders) LocationID() data.Coder {
	return c.Coder("Location", 10) // 100 shard, safe concurrent use is is 1
}
