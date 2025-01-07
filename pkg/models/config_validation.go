package models

//import (
//	"errors"
//	"strings"
//)
//
//// ValidateResourceConfig validates the resource configuration
//func ValidateResourceConfig(resource Resource) error {
//	// Validate Name
//	if strings.TrimSpace(resource.Name) == "" {
//		return errors.New("resource name cannot be empty")
//	}
//
//	// Validate Machine Type
//	if strings.TrimSpace(resource.MachineType) == "" {
//		return errors.New("machine type cannot be empty")
//	}
//
//	// Validate Region
//	if strings.TrimSpace(resource.Region) == "" {
//		return errors.New("region cannot be empty")
//	}
//
//	// Validate Image
//	if strings.TrimSpace(resource.Image) == "" || !strings.HasPrefix(resource.Image, "ami-") {
//		return errors.New("image ID must be specified and start with 'ami-'")
//	}
//
//	// Validate Subnet
//	if strings.TrimSpace(resource.Subnet) == "" || !strings.HasPrefix(resource.Subnet, "subnet-") {
//		return errors.New("subnet ID must be specified and start with 'subnet-'")
//	}
//
//	// Validate Security Groups
//	if len(resource.SecurityGroups) == 0 {
//		return errors.New("at least one security group must be specified")
//	}
//	for _, sg := range resource.SecurityGroups {
//		if !strings.HasPrefix(sg, "sg-") {
//			return errors.New("each security group ID must start with 'sg-'")
//		}
//	}
//
//	// Validate Key Name
//	if strings.TrimSpace(resource.KeyName) == "" {
//		return errors.New("key pair name cannot be empty")
//	}
//
//	return nil
//}
