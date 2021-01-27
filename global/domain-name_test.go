package global

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainName(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	g, err := New(ctx)
	assert.Nil(err)
	defer g.Close()

	taken, err := g.IsDomainNameTaken(ctx, "not.exist.com")
	assert.Nil(err)
	assert.False(taken)

	err = g.CreateDomainName(ctx, "test.domainname.com", "account")
	assert.Nil(err)

	taken, err = g.IsDomainNameTaken(ctx, "test.domainname.com")
	assert.Nil(err)
	assert.True(taken)

	err = g.RemoveDomainName(ctx, "test.domainname.com")
	assert.Nil(err)
}

func TestDomainNameRemoveByAccountID(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	g, err := New(ctx)
	assert.Nil(err)
	defer g.Close()

	err = g.CreateDomainName(ctx, "test1.domainname.com", "account1")
	assert.Nil(err)
	err = g.CreateDomainName(ctx, "test2.domainname.com", "account1")
	assert.Nil(err)

	err = g.RemoveDomainNameByAccountID(ctx, "account1")
	assert.Nil(err)

	taken, err := g.IsDomainNameTaken(ctx, "test1.domainname.com")
	assert.Nil(err)
	assert.False(taken)

	taken, err = g.IsDomainNameTaken(ctx, "test2.domainname.com")
	assert.Nil(err)
	assert.False(taken)

}
