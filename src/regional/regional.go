package regional

import (
	"context"

	data "github.com/piyuo/libsrv/src/data"
)

// Regional represent regional database
//
type Regional struct {
	data.BaseDB
}

// New regional db instance
//
//	db, err := regional.New(ctx, "")
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//
func New(ctx context.Context) (*Regional, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	conn, err := data.FirestoreRegionalConnection(ctx)
	if err != nil {
		return nil, err
	}

	c := &Regional{
		BaseDB: data.BaseDB{Connection: conn},
	}
	return c, nil
}

// Counters return collection of counter
//
func (c *Regional) Counters() *Counters {
	return &Counters{
		Counters: data.Counters{
			Connection: c.Connection,
			TableName:  "Count",
		},
	}
}

// Serials return collection of serial
//
func (c *Regional) Serials() *Serials {
	return &Serials{
		Serials: data.Serials{
			Connection: c.Connection,
			TableName:  "Serial",
		}}
}

// Coders return collection of coder
//
func (c *Regional) Coders() *Coders {
	return &Coders{
		Coders: data.Coders{
			Connection: c.Connection,
			TableName:  "Code",
		}}
}
