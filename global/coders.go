package global

import (
	data "github.com/piyuo/libsrv/data"
)

// Coders keep all coders
//
type Coders struct {
	data.Coders `firestore:"-"`
}

// AccountID return account id coder, 1000 shard, safe concurrent use is is 100
//
//	coder := global.AccountID()
//
func (c *Coders) AccountID() data.Coder {
	return c.Coder("Account", 1000)
}

// UserID return user id coder, 1000 shard, safe concurrent use is is 100
//
//	coder := d.UserID()
//
func (c *Coders) UserID() data.Coder {
	return c.Coder("User", 10000) // 10,000 shard, safe concurrent use is 1000
}
