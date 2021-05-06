package global

import (
	"github.com/piyuo/libsrv/db"
)

// Payment represent payment
//
type Payment struct {
	db.Model

	// pay items
	//

	// total amount
}

func (c *Payment) Factory() db.Object {
	return &Payment{}
}

func (c *Payment) Collection() string {
	return "Payment"
}
