package global

import (
	"github.com/piyuo/libsrv/db"
)

// Invoice represent Invoice
//
type Invoice struct {
	db.Model

	// pay items
	//

	// total amount
}

func (c *Invoice) Factory() db.Object {
	return &Invoice{}
}

func (c *Invoice) Collection() string {
	return "Invoice"
}
