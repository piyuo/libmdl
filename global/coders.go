package global

import (
	data "github.com/piyuo/libsrv/data"
)

// Coders keep all coders
//
type Coders struct {
	data.Coders `firestore:"-"`
}

// AccountID return account id coder
//
//	coder := d.AccountID()
//
func (c *Coders) AccountID() data.Coder {
	return c.Coder("AccountID", 100)
}
