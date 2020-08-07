package global

import (
	"context"

	data "github.com/piyuo/libsrv/data"
)

// Global represent global database
//
type Global struct {
	data.BaseDB
}

// New global db instance
//
//	g, err := global.New(ctx)
//	if err != nil {
//		return err
//	}
//	defer g.Close()
//
func New(ctx context.Context) (*Global, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	conn, err := data.FirestoreGlobalConnection(ctx)
	if err != nil {
		return nil, err
	}

	c := &Global{
		BaseDB: data.BaseDB{Connection: conn},
	}
	return c, nil
}

// Counters return collection of counter
//
func (c *Global) Counters() *Counters {
	return &Counters{
		Counters: data.Counters{
			Connection: c.Connection,
			TableName:  "Count",
		},
	}
}

// Serials return collection of serial
//
func (c *Global) Serials() *Serials {
	return &Serials{
		Serials: data.Serials{
			Connection: c.Connection,
			TableName:  "Serial",
		}}
}

// Coders return collection of coder
//
func (c *Global) Coders() *Coders {
	return &Coders{
		Coders: data.Coders{
			Connection: c.Connection,
			TableName:  "Code",
		}}
}
