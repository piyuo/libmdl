package regional

import (
	data "github.com/piyuo/libsrv/src/data"
)

// Coders keep all coders
//
type Coders struct {
	data.Coders `firestore:"-"`
}

// StoreID return store id coder, 1000 shard, safe concurrent use is is 100
//
//	coder := regional.StoreID()
//
func (c *Coders) StoreID() data.Coder {
	return c.Coder("Store", 1000)
}

// LocationID return location id coder, 1000 shard, safe concurrent use is 100
//
//	coder := d.LocationID()
//
func (c *Coders) LocationID() data.Coder {
	return c.Coder("Location", 1000)
}
