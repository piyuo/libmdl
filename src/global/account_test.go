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
	client, err := GlobalClient(ctx)
	assert.Nil(err)

	// account not exist
	account, err := GetAccountByID(ctx, "no-id")
	assert.Nil(err)
	assert.Nil(account)

	// add account
	id := identifier.UUID()
	account = &Account{}
	account.SetID(id)
	err = client.Set(ctx, account)
	assert.Nil(err)
	defer client.Delete(ctx, account)

	// account found
	account, err = GetAccountByID(ctx, id)
	assert.Nil(err)
	assert.NotNil(account)
}
