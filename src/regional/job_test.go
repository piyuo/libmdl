package regional

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJobCRUD(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()

	jobID, err := CreateJob(ctx)
	assert.Nil(err)
	assert.NotEmpty(jobID)

	found, progress, status, err := GetJob(ctx, jobID)
	assert.Nil(err)
	assert.True(found)
	assert.Zero(progress)
	assert.Empty(status)

	err = UpdateJob(ctx, jobID, 33, "loading")
	assert.Nil(err)

	found, progress, status, err = GetJob(ctx, jobID)
	assert.Nil(err)
	assert.True(found)
	assert.Equal(33, progress)
	assert.Equal("loading", status)

	err = DeleteJob(ctx, jobID)
	assert.Nil(err)

	found, progress, status, err = GetJob(ctx, jobID)
	assert.Nil(err)
	assert.False(found)
	assert.Zero(progress)
	assert.Empty(status)
}

func TestDeleteUnusedJob(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()
	client, err := RegionalClient(ctx)
	assert.Nil(err)

	//create record need to be remove
	createTime := time.Now().Add(time.Duration(-5) * time.Hour).UTC()
	expired := &Job{}
	expired.SetID("idExpired")
	expired.SetCreateTime(createTime)
	err = client.Set(ctx, expired)
	assert.Nil(err)

	valid := &Job{}
	valid.SetID("idNotExpire")
	err = client.Set(ctx, valid)
	assert.Nil(err)
	defer client.Delete(ctx, valid)

	// before cleanup
	found, _, _, err := GetJob(ctx, expired.ID())
	assert.Nil(err)
	assert.True(found)
	found, _, _, err = GetJob(ctx, valid.ID())
	assert.Nil(err)
	assert.True(found)

	// cleanup
	err = DeleteUnusedJob(ctx)
	assert.Nil(err)

	// after cleanup
	found, _, _, err = GetJob(ctx, expired.ID())
	assert.Nil(err)
	assert.False(found)
	found, _, _, err = GetJob(ctx, valid.ID())
	assert.Nil(err)
	assert.True(found)

}
