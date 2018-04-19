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

func importSnapshot(ctx context.Context, t task.Task) (*task.Task, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	compute := ec2.New(s)

	input := &ec2.ImportSnapshotInput{
		Description: aws.String(fmt.Sprintf("LinuxKit: %s", t.Name)),
		DiskContainer: &ec2.SnapshotDiskContainer{
			Description: aws.String(fmt.Sprintf("LinuxKit: %s disk", t.Name)),
			Format:      aws.String(format),
			UserBucket: &ec2.UserBucket{
				S3Bucket: aws.String(t.Bucket),
				S3Key:    aws.String(t.Key),
			},
		},
	}

	res, err := compute.ImportSnapshot(input)
	if err != nil {
		return nil, err
	}

	output := &task.Task{
		ImportTaskId: res.ImportTaskId,
		WaitTime:     defaultWaitTime,
	}

	return output, nil
}

func main() {
	lambda.Start(importSnapshot)
}
