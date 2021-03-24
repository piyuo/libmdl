package global

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainName(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()

	taken, err := IsDomainTaken(ctx, "not.exist.com")
	assert.Nil(err)
	assert.False(taken)

	err = CreateDomain(ctx, "test.domainname.com", "account")
	assert.Nil(err)

	taken, err = IsDomainTaken(ctx, "test.domainname.com")
	assert.Nil(err)
	assert.True(taken)

	err = DeleteDomain(ctx, "test.domainname.com")
	assert.Nil(err)
}

func TestDomainNameDeleteByAccountID(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()

	err := CreateDomain(ctx, "test1.domainname.com", "account1")
	assert.Nil(err)
	err = CreateDomain(ctx, "test2.domainname.com", "account1")
	assert.Nil(err)

	err = DeleteDomainByAccountID(ctx, "account1")
	assert.Nil(err)

	taken, err := IsDomainTaken(ctx, "test1.domainname.com")
	assert.Nil(err)
	assert.False(taken)

	taken, err = IsDomainTaken(ctx, "test2.domainname.com")
	assert.Nil(err)
	assert.False(taken)

}
