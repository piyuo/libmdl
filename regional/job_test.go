package regional

import (
	"context"
	"testing"
	"time"

	"github.com/piyuo/libsrv/data"
	"github.com/stretchr/testify/assert"
)

func TestJobCRUD(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	r, err := New(ctx)
	assert.Nil(err)
	defer r.Close()

	jobID, err := r.CreateJob(ctx)
	assert.Nil(err)
	assert.NotEmpty(jobID)

	found, progress, status, err := r.GetJob(ctx, jobID)
	assert.Nil(err)
	assert.True(found)
	assert.Zero(progress)
	assert.Empty(status)

	err = r.UpdateJob(ctx, jobID, 33, "loading")
	assert.Nil(err)

	found, progress, status, err = r.GetJob(ctx, jobID)
	assert.Nil(err)
	assert.True(found)
	assert.Equal(33, progress)
	assert.Equal("loading", status)

	err = r.DeleteJob(ctx, jobID)
	assert.Nil(err)

	found, progress, status, err = r.GetJob(ctx, jobID)
	assert.Nil(err)
	assert.False(found)
	assert.Zero(progress)
	assert.Empty(status)
}

func TestJobCleanup(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	r, err := New(ctx)
	assert.Nil(err)
	defer r.Close()

	//create record need to be remove
	createTime := time.Now().Add(time.Duration(-5) * time.Hour).UTC()
	table := r.JobTable()
	expired := &Job{
		BaseObject: data.BaseObject{
			ID:         "idExpired",
			CreateTime: createTime,
		},
	}
	err = table.Set(ctx, expired)
	assert.Nil(err)

	valid := &Job{
		BaseObject: data.BaseObject{
			ID: "idNotExpire",
		},
	}
	err = table.Set(ctx, valid)
	assert.Nil(err)
	defer table.DeleteObject(ctx, valid)

	// before cleanup
	found, _, _, err := r.GetJob(ctx, expired.ID)
	assert.Nil(err)
	assert.True(found)
	found, _, _, err = r.GetJob(ctx, valid.ID)
	assert.Nil(err)
	assert.True(found)

	// cleanup
	err = r.CleanupJob(ctx)
	assert.Nil(err)

	// after cleanup
	found, _, _, err = r.GetJob(ctx, expired.ID)
	assert.Nil(err)
	assert.False(found)
	found, _, _, err = r.GetJob(ctx, valid.ID)
	assert.Nil(err)
	assert.True(found)

}
