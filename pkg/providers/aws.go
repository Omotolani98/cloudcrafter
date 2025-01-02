package providers

import (
	"cloudcrafter/pkg/logger"
	"cloudcrafter/pkg/models"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"go.uber.org/zap"
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
//func (p *AWSProvider) CreateResource(resource models.Resource) (*models.ResourceMetadata, error) {
//	if resource.Type != "vm" {
//		return nil, errors.New("unsupported resource type for AWS")
//	}
//
//	input := &ec2.RunInstancesInput{
//		ImageId:      aws.String(resource.Image),
//		InstanceType: aws.String(resource.MachineType),
//		MinCount:     aws.Int64(1),
//		MaxCount:     aws.Int64(1),
//	}
//
//	result, err := p.ec2Client.RunInstances(input)
//	if err != nil {
//		return nil, fmt.Errorf("failed to create instance: %v", err)
//	}
//
//	instance := result.Instances[0]
//	return &models.ResourceMetadata{
//		ID:        *instance.InstanceId,
//		Name:      resource.Name,
//		Type:      resource.Type,
//		Provider:  "aws",
//		Region:    resource.Region,
//		Status:    *instance.State.Name,
//		CreatedAt: *instance.LaunchTime,
//	}, nil
//}

func (p *AWSProvider) CreateResource(resource models.Resource) (*models.ResourceMetadata, error) {
	// Define tags (including the instance name)
	tags := []*ec2.Tag{
		{
			Key:   aws.String("Name"),
			Value: aws.String(resource.Name), // Set the instance name
		},
	}

	// Create a tag specification for the EC2 instance
	tagSpecification := &ec2.TagSpecification{
		ResourceType: aws.String(ec2.ResourceTypeInstance), // Specify that these tags apply to an instance
		Tags:         tags,
	}

	// Define input for EC2 instance
	input := &ec2.RunInstancesInput{
		ImageId:           aws.String(resource.Image),       // AMI ID
		InstanceType:      aws.String(resource.MachineType), // Instance Type
		SubnetId:          aws.String(resource.Subnet),      // Subnet ID
		SecurityGroupIds:  aws.StringSlice(resource.SecurityGroups),
		KeyName:           aws.String(resource.KeyName), // Key Pair Name
		MinCount:          aws.Int64(1),                 // Number of instances to launch
		MaxCount:          aws.Int64(1),
		TagSpecifications: []*ec2.TagSpecification{tagSpecification}, // Attach tags
	}

	// Make the API call to AWS
	result, err := p.ec2Client.RunInstances(input)
	if err != nil {
		return nil, fmt.Errorf("failed to create EC2 instance: %v", err)
	}

	// Extract instance metadata
	instance := result.Instances[0]
	metadata := &models.ResourceMetadata{
		ID:        aws.StringValue(instance.InstanceId),
		Name:      resource.Name,
		Type:      "vm",
		Provider:  "aws",
		Region:    resource.Region,
		Status:    aws.StringValue(instance.State.Name),
		CreatedAt: time.Now(),
	}

	// Return metadata
	return metadata, nil
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

func (p *AWSProvider) ListResources() ([]models.ResourceMetadata, error) {
	logger.Log.Info("Listing resources for AWS provider")

	input := &ec2.DescribeInstancesInput{}
	result, err := p.ec2Client.DescribeInstances(input)
	if err != nil {
		logger.Log.Error("Failed to list instances", zap.Error(err))
		return nil, fmt.Errorf("failed to list instances: %v", err)
	}

	var resources []models.ResourceMetadata
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			var name string
			for _, tag := range instance.Tags {
				if *tag.Key == "Name" {
					name = *tag.Value
					break
				}
			}

			resource := models.ResourceMetadata{
				ID:        *instance.InstanceId,
				Name:      name,
				Type:      "vm",
				Provider:  "aws",
				Region:    *instance.Placement.AvailabilityZone,
				Status:    *instance.State.Name,
				CreatedAt: *instance.LaunchTime,
				UpdatedAt: time.Now(),
			}
			logger.Log.Info("Resource discovered",
				zap.String("id", resource.ID),
				zap.String("name", resource.Name),
				zap.String("status", resource.Status),
			)
			resources = append(resources, resource)
		}
	}

	if len(resources) == 0 {
		logger.Log.Warn("No resources found")
		return nil, fmt.Errorf("no resources found in the region")
	}

	logger.Log.Info("Resources listed successfully", zap.Int("count", len(resources)))
	return resources, nil
}
