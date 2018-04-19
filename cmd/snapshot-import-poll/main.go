package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/agy/linuxkit-builder/pkg/task"
)

func importSnapshotPoll(ctx context.Context, t task.Task) (*task.Task, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	compute := ec2.New(s)

	input := &ec2.DescribeImportSnapshotTasksInput{
		ImportTaskIds: []*string{
			t.ImportTaskId,
		},
	}

	res, err := compute.DescribeImportSnapshotTasks(input)
	if err != nil {
		return nil, err
	}

	if len(res.ImportSnapshotTasks) == 0 {
		return nil, fmt.Errorf("unable to read import snapshot task status")
	}

	detail := res.ImportSnapshotTasks[0].SnapshotTaskDetail

	output := &task.Task{
		Status:       *detail.Status,
		ImportTaskId: t.ImportTaskId,
		WaitTime:     t.WaitTime,
	}

	if detail.SnapshotId != nil {
		output.SnapshotId = *detail.SnapshotId
	}

	return output, nil
}

func main() {
	lambda.Start(importSnapshotPoll)
}
