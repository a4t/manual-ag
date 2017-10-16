package manualag

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

func (client Client) now() (int64, error) {
	input := &autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{
			aws.String(client.AgName),
		},
	}

	result, err := client.AutoScaling.DescribeAutoScalingGroups(input)
	if err != nil {
		return 0, err
	}

	if len(result.AutoScalingGroups) == 0 {
		return 0, fmt.Errorf(client.AgName + " is Not Found.")
	}
	return *result.AutoScalingGroups[0].DesiredCapacity, nil
}

func (client Client) Change(change int) (bool, error) {
	before_desired, _ := client.now()
	after_desired := int(before_desired) + change

	params := &autoscaling.UpdateAutoScalingGroupInput{
		AutoScalingGroupName: aws.String(client.AgName),
		DesiredCapacity:      aws.Int64(int64(after_desired)),
	}
	_, err := client.AutoScaling.UpdateAutoScalingGroup(params)
	if err != nil {
		return false, err
	}

	return true, nil
}
