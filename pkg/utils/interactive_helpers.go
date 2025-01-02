package utils

import (
	"cloudcrafter/pkg/models"
	"github.com/AlecAivazis/survey/v2"
)

// CollectInteractiveResourceData gathers resource data interactively
func CollectInteractiveResourceData() (models.Resource, error) {
	var resource models.Resource

	// Name
	_ = survey.AskOne(&survey.Input{
		Message: "Enter Resource Name (e.g., My Server):",
	}, &resource.KeyName)

	// Ask the user for the instance type (machine type)
	_ = survey.AskOne(&survey.Input{
		Message: "Enter Instance Type (e.g., t2.micro):",
	}, &resource.MachineType)

	// Region
	_ = survey.AskOne(&survey.Input{
		Message: "Enter region (e.g., us-east-1):",
	}, &resource.Region)

	// Image
	_ = survey.AskOne(&survey.Input{
		Message: "Enter AMI ID:",
	}, &resource.Image)

	// Subnet
	_ = survey.AskOne(&survey.Input{
		Message: "Enter Subnet ID:",
	}, &resource.Subnet)

	// Security Groups
	_ = survey.AskOne(&survey.Input{
		Message: "Enter Security Group IDs (comma-separated):",
	}, &resource.SecurityGroups)

	// Key Pair
	_ = survey.AskOne(&survey.Input{
		Message: "Enter Key Pair name:",
	}, &resource.KeyName)

	return resource, nil
}
