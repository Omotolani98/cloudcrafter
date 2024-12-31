package providers

import (
	"cloudcrafter/pkg/models"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AWSProvider struct {
	ec2Client *ec2.EC2
}

// NewAWSProvider initializes the AWS provider
func NewAWSProvider(region string) (*AWSProvider, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %v", err)
	}

	return &AWSProvider{
		ec2Client: ec2.New(sess),
	}, nil
}

// CreateResource creates a resource on AWS
func (p *AWSProvider) CreateResource(resource models.Resource) (*models.ResourceMetadata, error) {
	if resource.Type != "vm" {
		return nil, errors.New("unsupported resource type for AWS")
	}

	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(resource.Image),
		InstanceType: aws.String(resource.MachineType),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}

	result, err := p.ec2Client.RunInstances(input)
	if err != nil {
		return nil, fmt.Errorf("failed to create instance: %v", err)
	}

	instance := result.Instances[0]
	return &models.ResourceMetadata{
		ID:        *instance.InstanceId,
		Name:      resource.Name,
		Type:      resource.Type,
		Provider:  "aws",
		Region:    resource.Region,
		Status:    *instance.State.Name,
		CreatedAt: *instance.LaunchTime,
	}, nil
}

// DeleteResource deletes a resource on AWS
func (p *AWSProvider) DeleteResource(resourceID string) error {
	input := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{aws.String(resourceID)},
	}

	_, err := p.ec2Client.TerminateInstances(input)
	if err != nil {
		return fmt.Errorf("failed to terminate instance: %v", err)
	}

	return nil
}

// GetResource retrieves metadata about a resource
func (p *AWSProvider) GetResource(resourceID string) (*models.ResourceMetadata, error) {
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String(resourceID)},
	}

	result, err := p.ec2Client.DescribeInstances(input)
	if err != nil {
		return nil, fmt.Errorf("failed to describe instance: %v", err)
	}

	instance := result.Reservations[0].Instances[0]
	return &models.ResourceMetadata{
		ID:        *instance.InstanceId,
		Name:      "", // AWS doesn't allow custom names by default
		Type:      "vm",
		Provider:  "aws",
		Region:    *instance.Placement.AvailabilityZone,
		Status:    *instance.State.Name,
		CreatedAt: *instance.LaunchTime,
	}, nil
}
