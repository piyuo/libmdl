package regional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	verify := &Verify{}
	assert.NotNil(verify.Factory())
	assert.NotEmpty(verify.Collection())
}
