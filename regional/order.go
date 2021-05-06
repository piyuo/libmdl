package regional

import "github.com/piyuo/libsrv/db"

// Order represent Order in location
//
type Order struct {
	db.Model
}

func (c *Order) Factory() db.Object {
	return &Order{}
}

func (c *Order) Collection() string {
	return "Order"
}
