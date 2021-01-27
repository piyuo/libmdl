package global

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
	g, err := New(ctx)
	assert.Nil(err)
	defer g.Close()

	g.RemoveAllAccount(ctx)
	g.RemoveAllUser(ctx)
	g.RemoveAllDomainName(ctx)
}
