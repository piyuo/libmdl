package global

import (
	"context"
	"testing"

	"github.com/piyuo/libsrv/identifier"
	"github.com/stretchr/testify/assert"
)

func TestAccountByID(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	g, err := New(ctx)
	assert.Nil(err)
	defer g.Close()

	// account not exist
	account, err := g.GetAccountByID(ctx, "not-exist-id")
	assert.Nil(err)
	assert.Nil(account)

	// add account
	id := identifier.UUID()
	account = &Account{}
	account.SetID(id)
	err = g.AccountTable().Set(ctx, account)
	assert.Nil(err)
	defer g.AccountTable().DeleteObject(ctx, account)

	// account found
	account, err = g.GetAccountByID(ctx, id)
	assert.Nil(err)
	assert.NotNil(account)
}
