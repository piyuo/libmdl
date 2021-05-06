package regional

import (
	"github.com/piyuo/libsrv/db"
)

// Customer represent Customer in store
//
type Customer struct {
	db.Model
}

func (c *Customer) Factory() db.Object {
	return &Customer{}
}

func (c *Customer) Collection() string {
	return "Customer"
}
