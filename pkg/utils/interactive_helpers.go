package utils

import (
	"cloudcrafter/pkg/models"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

func CollectInteractiveResourceData() (models.Configuration, error) {
	var config models.Configuration

	providerPrompt := &survey.Select{
		Message: "Choose a cloud provider:",
		Options: []string{"aws", "azure", "gcp"},
	}
	if err := survey.AskOne(providerPrompt, &config.Provider); err != nil {
		return config, err
	}

	var resources []map[string]models.Resource
	for {
		var resourceType string
		resourcePrompt := &survey.Select{
			Message: "Select a resource type to add:",
			Options: []string{"vm", "storage", "done"},
		}
		if err := survey.AskOne(resourcePrompt, &resourceType); err != nil {
			return config, err
		}

		if resourceType == "done" {
			break
		}

		properties := make(map[string]string)
		switch resourceType {
		case "vm":
			properties = promptVMProperties()
		case "storage":
			properties = promptStorageProperties()
		default:
			fmt.Println("Unknown resource type")
			continue
		}

		resource := map[string]models.Resource{
			resourceType: {
				Type:       resourceType,
				Properties: properties,
			},
		}
		resources = append(resources, resource)
	}

	config.Resources = resources
	return config, nil
}

func promptVMProperties() map[string]string {
	properties := make(map[string]string)

	var vmName, machineType, region, image, subnet, securityGroups, keyName string

	survey.AskOne(&survey.Input{Message: "Enter VM name:"}, &vmName)
	properties["name"] = vmName

	survey.AskOne(&survey.Input{Message: "Enter machine type (e.g., t2.micro):"}, &machineType)
	properties["machineType"] = machineType

	survey.AskOne(&survey.Input{Message: "Enter region (e.g., us-east-1):"}, &region)
	properties["region"] = region

	survey.AskOne(&survey.Input{Message: "Enter image ID (e.g., ami-123456):"}, &image)
	properties["image"] = image

	survey.AskOne(&survey.Input{Message: "Enter subnet ID:"}, &subnet)
	properties["subnet"] = subnet

	survey.AskOne(&survey.Input{Message: "Enter security groups (comma-separated):"}, &securityGroups)
	properties["securityGroups"] = securityGroups

	survey.AskOne(&survey.Input{Message: "Enter key name:"}, &keyName)
	properties["keyName"] = keyName

	return properties
}

func promptStorageProperties() map[string]string {
	properties := make(map[string]string)

	var bucketName string
	err := survey.AskOne(&survey.Input{Message: "Enter bucket name:"}, &bucketName)
	if err != nil {
		_ = fmt.Errorf("%v\n", err)
		return nil
	}
	properties["bucketName"] = bucketName

	return properties
}
