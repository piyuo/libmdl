package global

import (
	"context"
	"testing"
	"time"

	"github.com/piyuo/libsrv/src/identifier"
	"github.com/stretchr/testify/assert"
)

func TestAccountStatus(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	account := Account{}

	account.RenewalDate = time.Now().UTC()
	assert.Equal(AccountStatusOK, account.Status())

	account.RenewalDate = time.Now().AddDate(0, 0, 1).UTC()
	assert.Equal(AccountStatusOK, account.Status())

	account.RenewalDate = time.Now().AddDate(0, 0, -61).UTC()
	assert.Equal(AccountStatusSuspend, account.Status())

	account.RenewalDate = time.Now().AddDate(0, 0, -1).UTC()
	assert.Equal(AccountStatusNonRenewal, account.Status())
}

func TestAccountSuspendDate(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	account := Account{
		RenewalDate: time.Now().UTC(),
	}

	suspendDate := account.SuspendDate()
	assert.True(suspendDate.After(time.Now().UTC()))
}

func TestAccountByID(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()
	g, err := GlobalClient(ctx)
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
	g, err := GlobalClient(ctx)
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
