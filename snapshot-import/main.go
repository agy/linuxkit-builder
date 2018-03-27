package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type event struct {
	Name   string `json:"name"`
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

type bucket struct {
	S3Bucket *string `json:"bucket"`
	S3Key    *string `json:"key"`
}

type detail struct {
	Description   *string  `json:"description"`
	DiskImageSize *float64 `json:"size"`
	Format        *string  `json:"format"`
	Progress      *string  `json:"progress"`
	SnapshotId    *string  `json:"snapshot_id"`
	Status        *string  `json:"status"`
	StatusMessage *string  `json:"status_message"`
	Url           *string  `json:"url"`
	UserBucket    *bucket  `json:"bucket"`
}

type output struct {
	ImportTaskId *string `json:"task_id"`
	Description  *string `json:"description"`
	TaskDetail   *detail `json:"task_detail"`
}

func importSnapshot(ctx context.Context, e event) (*output, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	compute := ec2.New(s)

	input := &ec2.ImportSnapshotInput{
		Description: aws.String(fmt.Sprintf("LinuxKit: %s", e.Name)),
		DiskContainer: &ec2.SnapshotDiskContainer{
			Description: aws.String(fmt.Sprintf("LinuxKit: %s disk", e.Name)),
			Format:      aws.String("raw"),
			UserBucket: &ec2.UserBucket{
				S3Bucket: aws.String(e.Bucket),
				S3Key:    aws.String(e.Key),
			},
		},
	}

	res, err := compute.ImportSnapshot(input)
	if err != nil {
		return nil, err
	}

	o := &output{
		Description:  res.Description,
		ImportTaskId: res.ImportTaskId,
		TaskDetail: &detail{
			Description:   res.SnapshotTaskDetail.Description,
			DiskImageSize: res.SnapshotTaskDetail.DiskImageSize,
			Format:        res.SnapshotTaskDetail.Format,
			Progress:      res.SnapshotTaskDetail.Progress,
			SnapshotId:    res.SnapshotTaskDetail.SnapshotId,
			Status:        res.SnapshotTaskDetail.Status,
			StatusMessage: res.SnapshotTaskDetail.StatusMessage,
			Url:           res.SnapshotTaskDetail.Url,
			UserBucket: &bucket{
				S3Bucket: res.SnapshotTaskDetail.UserBucket.S3Bucket,
				S3Key:    res.SnapshotTaskDetail.UserBucket.S3Key,
			},
		},
	}

	return o, nil
}

func main() {
	lambda.Start(importSnapshot)
}
