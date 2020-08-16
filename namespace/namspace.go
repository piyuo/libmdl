package namespace

import (
	"context"

	data "github.com/piyuo/libsrv/data"
)

// Namespace represent a account namespace
//
type Namespace struct {
	data.BaseDB
}

// New namespace instance
//
//	db, err := namespace.New(ctx, "")
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//
func New(ctx context.Context, namespace string) (*Namespace, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	conn, err := data.FirestoreRegionalConnection(ctx, namespace)
	if err != nil {
		return nil, err
	}

	c := &Namespace{
		BaseDB: data.BaseDB{Connection: conn},
	}
	return c, nil
}

// Counters return collection of counter
//
func (c *Namespace) Counters() *Counters {
	return &Counters{
		Counters: data.Counters{
			Connection: c.Connection,
			TableName:  "Count",
		},
	}
}

// Serials return collection of serial
//
func (c *Namespace) Serials() *Serials {
	return &Serials{
		Serials: data.Serials{
			Connection: c.Connection,
			TableName:  "Serial",
		}}
}

// Coders return collection of coder
//
func (c *Namespace) Coders() *Coders {
	return &Coders{
		Coders: data.Coders{
			Connection: c.Connection,
			TableName:  "Code",
		}}
}
