package regional

import (
	"context"

	data "github.com/piyuo/libsrv/data"
)

// DB represent regional database
//
type DB struct {
	data.BaseDB
}

// NewDB create db instance
//
func NewDB(ctx context.Context, namespace string) (*DB, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	conn, err := data.FirestoreRegionalConnection(ctx, namespace)
	if err != nil {
		return nil, err
	}

	db := &DB{
		BaseDB: data.BaseDB{Connection: conn},
	}
	return db, nil
}

// Counters return collection of counter
//
func (db *DB) Counters() *Counters {
	return &Counters{
		Counters: data.Counters{
			Connection: db.Connection,
			TableName:  "Count",
		},
	}
}

// Serials return collection of serial
//
func (db *DB) Serials() *Serials {
	return &Serials{
		Serials: data.Serials{
			Connection: db.Connection,
			TableName:  "Serial",
		}}
}

// Coders return collection of coder
//
func (db *DB) Coders() *Coders {
	return &Coders{
		Coders: data.Coders{
			Connection: db.Connection,
			TableName:  "Code",
		}}
}
