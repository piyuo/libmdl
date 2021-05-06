package regional

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreIDCoder(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()

	coder, err := StoreIDCoder(ctx)
	assert.Nil(err)
	assert.NotNil(coder)
}

func TestLocationIDCoder(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()

	coder, err := LocationIDCoder(ctx)
	assert.Nil(err)
	assert.NotNil(coder)
}
