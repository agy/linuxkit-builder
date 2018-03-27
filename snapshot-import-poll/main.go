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
	ImportTaskId string `json:"task_id"`
}

type output struct {
	Status     string `json:"status"`
	SnapshotId string `json:"snapshot_id,omitempty"`
}

func importSnapshotPoll(ctx context.Context, e event) (*output, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	compute := ec2.New(s)

	input := &ec2.DescribeImportSnapshotTasksInput{
		ImportTaskIds: []*string{
			aws.String(e.ImportTaskId),
		},
	}

	res, err := compute.DescribeImportSnapshotTasks(input)
	if err != nil {
		return nil, err
	}

	if len(res.ImportSnapshotTasks) == 0 {
		return nil, fmt.Errorf("unable to read import snapshot task status")
	}

	o := &output{
		Status:     *res.ImportSnapshotTasks[0].SnapshotTaskDetail.Status,
		SnapshotId: *res.ImportSnapshotTasks[0].SnapshotTaskDetail.SnapshotId,
	}

	return o, nil
}

func main() {
	lambda.Start(importSnapshotPoll)
}
