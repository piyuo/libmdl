package global

import (
	"github.com/piyuo/libsrv/data"
)

// Payment represent payment
//
type Payment struct {
	data.DomainObject

	// pay items
	//

	// total amount
}

// PaymentTable return account table
//
//	table := db.PaymentTable()
//
func (c *Global) PaymentTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Payment",
		Factory: func() data.Object {
			return &Payment{}
		},
	}
}
