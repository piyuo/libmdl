package global

import (
	"context"
	"testing"

	"github.com/piyuo/libsrv/identifier"
	"github.com/stretchr/testify/assert"
)

func TestAccountByID(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()
	g, err := Client(ctx)
	assert.Nil(err)

	// account not exist
	account, err := GetAccountByID(ctx, "no-id")
	assert.Nil(err)
	assert.Nil(account)

	// add account
	id := identifier.UUID()
	account = &Account{}
	account.SetID(id)
	err = g.Set(ctx, account)
	assert.Nil(err)
	defer g.Delete(ctx, account)

	// account found
	account, err = GetAccountByID(ctx, id)
	assert.Nil(err)
	assert.NotNil(account)
}

func TestAccountNilSafety(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()
	g, err := Client(ctx)
	assert.Nil(err)

	account := &Account{}
	err = g.Set(ctx, account)
	assert.Nil(err)
	defer g.Delete(ctx, account)

	obj, err := g.Get(ctx, &Account{}, account.ID())
	assert.Nil(err)
	account2 := obj.(*Account)
	assert.NotNil(account2.Roles)
}
