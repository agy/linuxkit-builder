package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/agy/linuxkit-builder/pkg/task"
)

const (
	defaultWaitTime = 60
	format          = "raw"
)

type event struct {
	Name   string `json:"name"`
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

func importSnapshot(ctx context.Context, e event) (*task.Task, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	compute := ec2.New(s)

	input := &ec2.ImportSnapshotInput{
		Description: aws.String(fmt.Sprintf("LinuxKit: %s", e.Name)),
		DiskContainer: &ec2.SnapshotDiskContainer{
			Description: aws.String(fmt.Sprintf("LinuxKit: %s disk", e.Name)),
			Format:      aws.String(format),
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

	t := &task.Task{
		ImportTaskId: res.ImportTaskId,
		WaitTime:     defaultWaitTime,
	}

	return t, nil
}

func main() {
	lambda.Start(importSnapshot)
}
