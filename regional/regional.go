package regional

import (
	"context"

	data "github.com/piyuo/libsrv/data"
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

	conn, err := data.FirestoreRegionalConnection(ctx, "")
	if err != nil {
		return nil, err
	}

	c := &Regional{
		BaseDB: data.BaseDB{Connection: conn},
	}
	return c, nil
}
