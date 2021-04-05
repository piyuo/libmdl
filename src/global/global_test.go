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

func TestUserIDCoder(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()

	coder, err := UserIDCoder(ctx)
	assert.Nil(err)
	assert.NotNil(coder)
}
