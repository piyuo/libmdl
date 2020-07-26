package global

import (
	data "github.com/piyuo/libsrv/data"
)

// Counters is collection of global usage counters
//
type Counters struct {
	data.Counters `firestore:"-"`
}

// AccountTotal return account total counter
//
//	counter := d.AccountCounter()
//
func (c *Counters) AccountTotal() data.Counter {
	return c.Counter("AccountTotal", 100)
}
