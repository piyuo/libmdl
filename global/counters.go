package global

import (
	data "github.com/piyuo/libsrv/data"
)

// Counters is collection of global usage counters
//
type Counters struct {
	data.Counters `firestore:"-"`
}

// AccountCounter return account counter
//
//	counter := global.AccountCounter()
//
func (c *Counters) AccountCounter() data.Counter {
	return c.Counter("Account", 1000) // 1,000 shard, safe concurrent use is is 100
}

// UserCounter return user counter
//
//	counter := global.UserCounter()
//
func (c *Counters) UserCounter() data.Counter {
	return c.Counter("User", 10000) // 10,000 shard, safe concurrent use is 1000
}
