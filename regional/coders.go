package rmdl

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
//	coder := d.LocationID()
//
func (c *Coders) LocationID() data.Coder {
	return c.Coder("LocationID", 100)
}
