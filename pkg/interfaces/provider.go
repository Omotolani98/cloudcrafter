package providers

import "cloudcrafter/pkg/models"

type Provider interface {
type Provider interface {
	CreateResource(resource models.Resource) (*models.ResourceMetadata, error)
	DeleteResource(resourceID string) error
	GetResource(resourceID string) (*models.ResourceMetadata, error)
       ListResources() ([]models.ResourceMetadata, error)

	CreateBucket(bucketName string) error
	ListBuckets() ([]models.S3Bucket, error)
	DeleteBucket(bucketName string) error
	UploadObject(bucketName, key, filePath string) error
}
