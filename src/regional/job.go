package regional

import (
	"context"
	"time"

	"github.com/piyuo/libsrv/src/db"
	"github.com/piyuo/libsrv/src/log"
)

// Job keep track a long running job
//
type Job struct {
	db.Entity

	// Progress is complete percentage range from 0 to 100
	//
	Progress int `firestore:"Progress,omitempty"`

	// Status is current status of long running job
	//
	Status string `firestore:"Status,omitempty"`
}

func (c *Job) Factory() db.Object {
	return &Job{}
}

func (c *Job) Collection() string {
	return "Job"
}

// CreateJob create job and return job id
//
//	jobID,err := CreateJob(ctx)
//
func CreateJob(ctx context.Context) (string, error) {
	client, err := RegionalClient(ctx)
	if err != nil {
		return "", err
	}

	job := &Job{}
	if err := client.Set(ctx, job); err != nil {
		return "", err
	}
	return job.ID(), nil
}

// GetJob return progress of job
//
//	found,progress,status,err := GetJob(ctx,jobID)
//
func GetJob(ctx context.Context, jobID string) (bool, int, string, error) {
	client, err := RegionalClient(ctx)
	if err != nil {
		return false, 0, "", err
	}
	obj, err := client.Get(ctx, &Job{}, jobID)
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
func UpdateJob(ctx context.Context, jobID string, progress int, status string) error {
	client, err := RegionalClient(ctx)
	if err != nil {
		return err
	}
	job := &Job{}
	job.SetID(jobID)
	return client.Update(ctx, job, map[string]interface{}{
		"Progress": progress,
		"Status":   status,
	})
}

// DeleteJob delete a job
//
//	err := DeleteJob(ctx,jobID)
//
func DeleteJob(ctx context.Context, jobID string) error {
	client, err := RegionalClient(ctx)
	if err != nil {
		return err
	}
	job := &Job{}
	job.SetID(jobID)
	return client.Delete(ctx, job)
}

// DeleteUnusedJob delete jobs created more than one day
//
//	err := DeleteUnusedJob(ctx)
//
func DeleteUnusedJob(ctx context.Context) error {
	client, err := RegionalClient(ctx)
	if err != nil {
		return err
	}
	// a job should not execute longer than 60 min. we do cleanup after 4 hour for safe
	deadline := time.Now().Add(time.Duration(-4) * time.Hour).UTC()
	done, err := client.Query(&Job{}).Where("CreateTime", "<", deadline).Delete(ctx, 100)
	if done {
		log.Info(ctx, "del unused job done")
		return err
	}
	log.Warn(ctx, "del unused job not done")
	return err
}
