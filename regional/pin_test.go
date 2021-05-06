package regional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPin(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	pin := &Pin{}
	assert.NotNil(pin.Factory())
	assert.NotEmpty(pin.Collection())
}
