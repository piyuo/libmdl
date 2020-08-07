package regional

import (
	data "github.com/piyuo/libsrv/data"
)

// Verify keep verify code for email verify, ID is email address
//
type Verify struct {
	data.BaseObject

	// Code is verify code
	//
	Code string
}

// VerifyTable return Verify object
//
//	table := regional.VerifyTable()
//
func (c *Regional) VerifyTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Verify",
		Factory: func() data.Object {
			return &Verify{}
		},
	}
}
