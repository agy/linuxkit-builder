package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
)

const (
	name = "aws-linuxkit"
)

var (
	stateMachineARN = ""
)

type output struct {
	Name   string `json:"name"`
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

func parseEvent(e events.S3Event) (string, string, error) {
	if len(e.Records) == 0 {
		return "", "", errors.New("no records found")
	}

	s3 := e.Records[0].S3

	return s3.Bucket.Name, s3.Object.Key, nil
}

func invokeSfn(ctx context.Context, e events.S3Event) (*output, error) {
	bucket, key, err := parseEvent(e)
	if err != nil {
		return nil, err
	}

	sfnInput := &output{
		Name:   name,
		Bucket: bucket,
		Key:    key,
	}

	execInput, err := json.Marshal(sfnInput)
	if err != nil {
		return nil, errors.New("could not encode sfn input")
	}

	fmt.Println("execInput:", execInput)

	// TODO(agy): Add timestamp
	execName := name

	input := &sfn.StartExecutionInput{
		Input:           aws.String(string(execInput)),
		Name:            aws.String(execName),
		StateMachineArn: aws.String(stateMachineARN),
	}

	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	svc := sfn.New(sess)

	if _, err := svc.StartExecution(input); err != nil {
		return nil, err
	}

	return sfnInput, nil
}

func main() {
	stateMachineARN = os.Getenv("STATEMACHINEARN")
	if stateMachineARN == "" {
		panic("invalid state machine ARN")
	}

	fmt.Println("STATEMACHINEARN:", stateMachineARN)

	lambda.Start(invokeSfn)
}
