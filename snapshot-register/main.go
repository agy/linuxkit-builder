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
	arch                = "x86_64"
	deleteOnTermination = true
	name                = "linuxkit"
	rootDevice          = "/dev/sda1"
	virtType            = "hvm"
	volumeType          = "standard"
)

var (
	sriovNetSupport string
	enaSupport      bool
)

type event struct {
	SnapshotId string `json:"snapshot_id"`
}

type output struct {
	ImageId string `json:"image_id"`
}

func registerImage(ctx context.Context, e event) (*output, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	compute := ec2.New(s)

	input := &ec2.RegisterImageInput{
		Name:         aws.String(name),
		Architecture: aws.String(arch),
		BlockDeviceMappings: []*ec2.BlockDeviceMapping{
			{
				DeviceName: aws.String(rootDevice),
				Ebs: &ec2.EbsBlockDevice{
					DeleteOnTermination: aws.Bool(deleteOnTermination),
					SnapshotId:          aws.String(e.SnapshotId),
					VolumeType:          aws.String(volumeType),
				},
			},
		},
		Description:        aws.String(fmt.Sprintf("LinuxKit: %s image", name)),
		RootDeviceName:     aws.String(rootDevice),
		VirtualizationType: aws.String(virtType),
		EnaSupport:         aws.Bool(enaSupport),
	}

	if sriovNetSupport != "" {
		input = input.SetSriovNetSupport(sriovNetSupport)
	}

	res, err := compute.RegisterImage(input)
	if err != nil {
		return nil, err
	}

	o := &output{
		ImageId: *res.ImageId,
	}

	return o, nil
}

func main() {
	lambda.Start(registerImage)
}
