package regional

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {}

func shutdown() {}

func TestClean(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	r, err := New(ctx)
	assert.Nil(err)
	defer r.Close()

	r.RemoveAllStore(ctx)
	r.RemoveAllLocation(ctx)
	r.RemoveAllJob(ctx)
	r.RemoveAllVerificationCode(ctx)
}
