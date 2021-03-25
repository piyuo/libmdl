package regional

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocationNilSafety(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()
	r, err := RegionalClient(ctx)
	assert.Nil(err)

	location := &Location{}
	err = r.Set(ctx, location)
	assert.Nil(err)
	defer r.Delete(ctx, location)

	obj, err := r.Get(ctx, &Location{}, location.ID())
	assert.Nil(err)
	location2 := obj.(*Location)
	assert.NotNil(location2.Hours)
}
