package regional

import (
	"context"
	"fmt"
	"time"

	"github.com/piyuo/libsrv/data"
)

// Job keep track a long running job
//
type Job struct {
	data.BaseObject

	// Progress is complete percentage range from 0 to 100
	//
	Progress int

	// Status is current status of long running job
	//
	Status string
}

// JobTable return job table
//
//	table := JobTable(r)
//
func (c *Regional) JobTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "Job",
		Factory: func() data.Object {
			return &Job{}
		},
	}
}

// CreateJob create job and return job id
//
//	table := CreateJob(ctx)
//
func (c *Regional) CreateJob(ctx context.Context) (string, error) {
	job := &Job{}
	if err := c.JobTable().Set(ctx, job); err != nil {
		return "", err
	}
	return job.ID, nil
}

// GetJob return progress of job
//
//	found,progress,status,err := GetJob(ctx,jobID)
//
func (c *Regional) GetJob(ctx context.Context, jobID string) (bool, int, string, error) {
	obj, err := c.JobTable().Get(ctx, jobID)
	if err != nil {
		return false, 0, "", err
	}
	if obj == nil {
		return false, 0, "", nil
	}
	job := obj.(*Job)
	return true, job.Progress, job.Status, nil
}

// UpdateJob set job progress
//
//	err := UpdateJob(ctx,jobID,30,"copy files...")
//
func (c *Regional) UpdateJob(ctx context.Context, jobID string, progress int, status string) error {
	return c.JobTable().Update(ctx, jobID, map[string]interface{}{
		"Progress": progress,
		"Status":   status,
	})
}

// DeleteJob delete a job
//
//	err := DeleteJob(ctx,jobID)
//
func (c *Regional) DeleteJob(ctx context.Context, jobID string) error {
	return c.JobTable().Delete(ctx, jobID)
}

// ClearJob clear all job
//
//	err := ClearJob(ctx)
//
func (c *Regional) ClearJob(ctx context.Context) error {
	return c.JobTable().Clear(ctx)
}

// CleanupJob cleanup jobs created more than one day
//
//	err := CleanupJob(ctx)
//
func (c *Regional) CleanupJob(ctx context.Context) error {
	// a job should not execute longer than 60 min. we do cleanup after 4 hour for safe
	deadline := time.Now().Add(time.Duration(-4) * time.Hour).UTC()
	count, err := c.JobTable().Query().Where("CreateTime", "<", deadline).Clear(ctx)
	if count > 0 {
		fmt.Printf("cleanup %v Job\n", count)
	}
	return err
}
