package global

import (
	"github.com/piyuo/libsrv/data"
)

// Invoice represent Invoice
//
type Invoice struct {
	data.BaseObject

	// pay items
	//

	// total amount
}

// InvoiceTable return invoice table
//
//	table := db.InvoiceTable()
//
func (c *Global) InvoiceTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Invoice",
		Factory: func() data.Object {
			return &Invoice{}
		},
	}
}
