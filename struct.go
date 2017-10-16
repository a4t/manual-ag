package manualag

import (
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

type Config struct {
	Region             string
	ProfileName        string
	AwsAccessKeyId     string
	AwsSecretAccessKey string
}

type Client struct {
	AutoScaling *autoscaling.AutoScaling
	AgName      string
}
