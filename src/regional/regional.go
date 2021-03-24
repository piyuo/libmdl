package regional

import (
	"context"

	"github.com/piyuo/libsrv/src/db"
	"github.com/piyuo/libsrv/src/google/gaccount"
	"github.com/piyuo/libsrv/src/google/gdb"
	"github.com/pkg/errors"
)

var regionalClient db.Client

// RegionalClient regional global client, client don't need close and it can be resuse in go routines
//
//	client,err := RegionalClient(ctx)
//
func RegionalClient(ctx context.Context) (db.Client, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	if regionalClient != nil && !regionalClient.IsClose() {
		return regionalClient, nil
	}

	cred, err := gaccount.RegionalCredential(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get regional cred")
	}
	client, err := gdb.NewClient(ctx, cred)
	if err != nil {
		return nil, errors.Wrap(err, "new client")
	}
	regionalClient = client
	return regionalClient, nil
}

// StoreIDCoder generate store id, it can generate id at 100/concurrent
//
func StoreIDCoder(client db.Client) db.Coder {
	return client.Coder("StoreID", 1000)
}

// LocationIDCoder generate location id, it can generate id at 100/concurrent
//
func LocationIDCoder(client db.Client) db.Coder {
	return client.Coder("LocationID", 1000)
}
