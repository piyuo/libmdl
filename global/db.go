package global

import (
	"context"

	data "github.com/piyuo/libsrv/data"
)

// DB represent global database
//
type DB struct {
	data.BaseDB
}

// NewDB create db instance
//
func NewDB(ctx context.Context) (*DB, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	conn, err := data.FirestoreGlobalConnection(ctx)
	if err != nil {
		return nil, err
	}

	db := &DB{
		BaseDB: data.BaseDB{CurrentConnection: conn},
	}
	return db, nil
}

// Counters return collection of counter
//
func (db *DB) Counters() *Counters {
	return &Counters{
		Counters: data.Counters{
			CurrentConnection: db.CurrentConnection,
			TableName:         "count",
		},
	}
}

// Serials return collection of serial
//
func (db *DB) Serials() *Serials {
	return &Serials{
		Serials: data.Serials{
			CurrentConnection: db.CurrentConnection,
			TableName:         "serial",
		}}
}

// Coders return collection of coder
//
func (db *DB) Coders() *Coders {
	return &Coders{
		Coders: data.Coders{
			CurrentConnection: db.CurrentConnection,
			TableName:         "code",
		}}
}
