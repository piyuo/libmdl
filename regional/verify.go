package regional

import (
	"time"

	"github.com/piyuo/libsrv/data"
)

// Verify keep verify code for email verify
//
//	ID is email address
//
type Verify struct {
	data.BaseObject

	// Hash is code hash with salt, we do not store code only hash is enough
	//
	Hash uint32

	// Crypted code
	//
	Crypted string

	// Issue is code issue time
	//
	Issue time.Time
}

// VerifyTable return Verify table
//
//	table := VerifyTable(r)
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
