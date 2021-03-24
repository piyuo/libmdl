package regional

import (
	"github.com/piyuo/libsrv/src/db"
)

// Cart represent Cart in store
//
type Cart struct {
	db.Model
}

func (c *Cart) Factory() db.Object {
	return &Cart{}
}

func (c *Cart) Collection() string {
	return "Cart"
}
