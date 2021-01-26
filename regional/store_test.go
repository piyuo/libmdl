package regional

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDomainTaken(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	r, err := New(ctx)
	assert.Nil(err)
	defer r.Close()

	// not taken
	taken, err := r.IsDomainTaken(ctx, "not-exist.com")
	assert.Nil(err)
	assert.False(taken)

	// add store with domain
	store := &Store{
		Domain: "access-taken.domain",
	}
	r.StoreTable().Set(ctx, store)
	defer r.StoreTable().DeleteObject(ctx, store)

	//taken
	taken, err = r.IsDomainTaken(ctx, "access-taken.domain")
	assert.Nil(err)
	assert.True(taken)
}
