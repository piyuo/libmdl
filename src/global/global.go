package global

import (
	"context"

	"github.com/piyuo/libsrv/src/db"
	"github.com/piyuo/libsrv/src/google/gaccount"
	"github.com/piyuo/libsrv/src/google/gdb"
	"github.com/pkg/errors"
)

var globalClient db.Client

// GlobalClient return global client, client don't need close and it can be resuse in go routines
//
//	client,err := GlobalClient(ctx)
//
func GlobalClient(ctx context.Context) (db.Client, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	if globalClient != nil && !globalClient.IsClose() {
		return globalClient, nil
	}

	cred, err := gaccount.GlobalCredential(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get global cred")
	}
	client, err := gdb.NewClient(ctx, cred)
	if err != nil {
		return nil, errors.Wrap(err, "new client")
	}
	globalClient = client
	return globalClient, nil
}

// AccountIDCoder can generate id at 100/concurrent
//
func AccountIDCoder(ctx context.Context) (db.Coder, error) {
	client, err := GlobalClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.Coder("AccountID", 1000), nil
}

// AccountCounter can count at 100/concurrent
//
func AccountCounter(ctx context.Context) (db.Counter, error) {
	client, err := GlobalClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.Counter("AccountCount", 1000, db.DateHierarchyFull), nil
}

// UserIDCoder can generate id at 100/concurrent
//
func UserIDCoder(ctx context.Context) (db.Coder, error) {
	client, err := GlobalClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.Coder("UserID", 1000), nil
}

// UserCounter can count at 100/concurrent
//
func UserCounter(ctx context.Context) (db.Counter, error) {
	client, err := GlobalClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.Counter("UserCount", 1000, db.DateHierarchyFull), nil
}
