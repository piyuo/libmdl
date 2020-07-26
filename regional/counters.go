package rmdl

import (
	data "github.com/piyuo/libsrv/data"
)

// Counters is collection of global usage counters
//
type Counters struct {
	data.Counters `firestore:"-"`
}

// LocationTotal return total location count
//
//	id := d.LocationTotal(ctx)
//
func (c *Counters) LocationTotal() data.Counter {
	return c.Counter("LocationTotal", 10)
}
