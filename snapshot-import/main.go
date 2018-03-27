package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	defaultWaitTime = 60
)

type event struct {
	Name   string `json:"name"`
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

type output struct {
	ImportTaskId *string `json:"task_id"`
	WaitTime     int     `json:"wait_time"`
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
		ImportTaskId: res.ImportTaskId,
		WaitTime:     defaultWaitTime,
	}

	return o, nil
}

func main() {
	lambda.Start(importSnapshot)
}
