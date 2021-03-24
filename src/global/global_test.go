package global

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountIDCoder(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()

	coder, err := AccountIDCoder(ctx)
	assert.Nil(err)
	assert.NotNil(coder)
}

func TestAccountCounter(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()

	counter, err := AccountCounter(ctx)
	assert.Nil(err)
	assert.NotNil(counter)
}

func TestUserIDCoder(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()

	coder, err := UserIDCoder(ctx)
	assert.Nil(err)
	assert.NotNil(coder)
}

func TestUserCounter(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()

	counter, err := UserCounter(ctx)
	assert.Nil(err)
	assert.NotNil(counter)
}
