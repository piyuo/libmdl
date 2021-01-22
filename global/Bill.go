package global

import (
	"time"

	"github.com/piyuo/libsrv/data"
)

// Bill represent bill
//
type Bill struct {
	data.BaseObject

	// BeginDate is billing period date begin
	//
	BeginDate time.Time

	// EndDate is billing period date end
	//
	EndDate time.Time

	// Email is owner email where bill send
	//
	Email string

	// FirstName is owner first name where bill send
	//
	FirstName string

	// LastName is owner last name where bill send
	//
	LastName string

	// Plan is account servie plan
	//
	Plan AccountPlan

	// Currency is plan fee currency
	//
	Currency string

	// Plan is account servie plan
	//
	PlanServiceFee int64

	//additional billable item
}

// BillTable return bill table
//
//	table := db.BillTable()
//
func (c *Global) BillTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Bill",
		Factory: func() data.Object {
			return &Bill{}
		},
	}
}
