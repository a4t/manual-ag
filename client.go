package manualag

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

func Init(agname string, config *Config) Client {
	return Client{
		AutoScaling: autoscalingClient(config),
		AgName:      agname,
	}
}

func getClientSession(config *Config) *session.Session {
	if len(config.ProfileName) != 0 {
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			Config:  aws.Config{Region: aws.String(config.Region)},
			Profile: config.ProfileName,
		}))
		return sess
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(config.Region),
	}))
	return sess
}

func autoscalingClient(config *Config) *autoscaling.AutoScaling {
	if len(config.AwsAccessKeyId) != 0 {
		creds := credentials.NewStaticCredentials(config.AwsAccessKeyId, config.AwsSecretAccessKey, "")
		svc := autoscaling.New(
			getClientSession(config),
			aws.NewConfig().WithRegion(config.Region).WithCredentials(creds),
		)
		return svc
	}

	svc := autoscaling.New(getClientSession(config))
	return svc
}
