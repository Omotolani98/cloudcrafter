package providers

import (
	"cloudcrafter/pkg/logger"
	"cloudcrafter/pkg/models"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AWSProvider struct {
	ec2Client *ec2.EC2
	s3Client  *s3.S3
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
		s3Client:  s3.New(sess),
	}, nil
}

func (p *AWSProvider) CreateResource(resource models.Resource) (*models.ResourceMetadata, error) {
	image, ok := resource.Properties["image"]
	if !ok {
		return nil, fmt.Errorf("missing required property: image")
	}

	machineType, ok := resource.Properties["machineType"]
	if !ok {
		return nil, fmt.Errorf("missing required property: machineType")
	}

	subnet, ok := resource.Properties["subnet"]
	if !ok {
		return nil, fmt.Errorf("missing required property: subnet")
	}

	keyName, ok := resource.Properties["keyName"]
	if !ok {
		return nil, fmt.Errorf("missing required property: keyName")
	}

	securityGroups, ok := resource.Properties["securityGroups"]
	if !ok {
		return nil, fmt.Errorf("missing required property: securityGroups")
	}

	securityGroupIDs := aws.StringSlice(strings.Split(securityGroups, ","))

	tags := []*ec2.Tag{
		{
			Key:   aws.String("Name"),
			Value: aws.String(resource.Properties["name"]),
		},
	}

	tagSpecification := &ec2.TagSpecification{
		ResourceType: aws.String(ec2.ResourceTypeInstance),
		Tags:         tags,
	}

	input := &ec2.RunInstancesInput{
		ImageId:           aws.String(image),
		InstanceType:      aws.String(machineType),
		SubnetId:          aws.String(subnet),
		SecurityGroupIds:  securityGroupIDs,
		KeyName:           aws.String(keyName),
		MinCount:          aws.Int64(1),
		MaxCount:          aws.Int64(1),
		TagSpecifications: []*ec2.TagSpecification{tagSpecification},
	}

	result, err := p.ec2Client.RunInstances(input)
	if err != nil {
		return nil, fmt.Errorf("failed to create EC2 instance: %v", err)
	}

	instance := result.Instances[0]
	metadata := &models.ResourceMetadata{
		ID:        aws.StringValue(instance.InstanceId),
		Name:      resource.Properties["name"],
		Type:      "vm",
		Provider:  "aws",
		Region:    resource.Properties["region"],
		Status:    aws.StringValue(instance.State.Name),
		CreatedAt: time.Now(),
	}

	return metadata, nil
}

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
	input := &ec2.DescribeInstancesInput{}
	result, err := p.ec2Client.DescribeInstances(input)
	if err != nil {
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

			resources = append(resources, resource)
		}
	}

	if len(resources) == 0 {
		logger.Log.Warn("No resources found")
		return nil, fmt.Errorf("no resources found in the region")
	}

	fmt.Printf("\nResources listed successfully: %d\n", len(resources))
	return resources, nil
}

func (p *AWSProvider) CreateBucket(bucketName string) error {
	fmt.Printf("Creating S3 bucket...%s\n", bucketName)
	_, err := p.s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return fmt.Errorf("failed to create bucket: %w", err)
	}
	return nil
}

// ListBuckets lists all S3 buckets
func (p *AWSProvider) ListBuckets() ([]models.S3Bucket, error) {
	fmt.Println("Listing S3 buckets...")
	output, err := p.s3Client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to list buckets: %w", err)
	}

	var buckets []models.S3Bucket
	for _, bucket := range output.Buckets {
		buckets = append(buckets, models.S3Bucket{
			Name:         aws.StringValue(bucket.Name),
			CreationDate: bucket.CreationDate,
		})
	}
	return buckets, nil
}

func (p *AWSProvider) DeleteBucket(bucketName string) error {
	fmt.Printf("Deleting S3 bucket...%s\n", bucketName)
	_, err := p.s3Client.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return fmt.Errorf("failed to delete bucket: %w", err)
	}
	return nil
}

// UploadObject uploads a file to an S3 bucket
func (p *AWSProvider) UploadObject(bucketName, key, filePath string) error {
	fmt.Printf("Uploading object to S3...%s :: %s\n", bucketName, key)
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = p.s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("failed to upload object: %w", err)
	}
	return nil
}

func (p *AWSProvider) DownloadObject(bucketName, key, filePath string) error {
	fmt.Printf("Downloading object from S3...%s :: %s\n", bucketName, key)
	output, err := p.s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to download object: %w", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(output.Body)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = io.Copy(file, output.Body)
	if err != nil {
		return fmt.Errorf("failed to write object to file: %w", err)
	}
	return nil
}
