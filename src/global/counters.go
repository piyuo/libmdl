package global

import (
	data "github.com/piyuo/libsrv/src/data"
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
	return c.Counter("Account", 1000, data.DateHierarchyFull) // 1,000 shard, safe concurrent use is is 100
}

// UserCounter return user counter
//
//	counter := global.UserCounter()
//
func (c *Counters) UserCounter() data.Counter {
	return c.Counter("User", 1000, data.DateHierarchyFull) // 10,00 shard, safe concurrent use is 100
}
