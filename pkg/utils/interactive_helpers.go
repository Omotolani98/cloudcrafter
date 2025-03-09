package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"

	"github.com/Omotolani98/cloudcrafter/pkg/models"
)

func CollectInteractiveResourceData(networkingChoice string) (models.Configuration, error) {
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

	// Handle Networking
	reader := bufio.NewReader(os.Stdin)
	var networking models.NetworkingConfig

	if networkingChoice == "custom" {
		fmt.Print("Enter VPC CIDR Block (e.g., 10.0.0.0/16): ")
		vpcCidr, _ := reader.ReadString('\n')
		networking.VPC.CIDRBlock = strings.TrimSpace(vpcCidr)

		fmt.Print("Enter Subnet CIDR Block (e.g., 10.0.1.0/24): ")
		subnetCidr, _ := reader.ReadString('\n')
		networking.Subnet.CIDRBlock = strings.TrimSpace(subnetCidr)

		// Add security group configuration if needed
		networking.SecurityGroup.Description = "Allow SSH and HTTP"
		networking.SecurityGroup.IngressRules = []models.IngressRule{
			{Protocol: "tcp", FromPort: 22, ToPort: 22, CIDRIP: "0.0.0.0/0"},
			{Protocol: "tcp", FromPort: 80, ToPort: 80, CIDRIP: "0.0.0.0/0"},
		}
	} else {
		// Apply default networking config
		networking = models.NetworkingConfig{
			VPC: struct {
				CIDRBlock string `json:"cidrBlock"`
			}{
				CIDRBlock: "10.0.0.0/16",
			},
			Subnet: struct {
				CIDRBlock string `json:"cidrBlock"`
			}{
				CIDRBlock: "10.0.1.0/24",
			},
			SecurityGroup: struct {
				Description  string               `json:"description"`
				IngressRules []models.IngressRule `json:"ingressRules"`
			}{
				Description: "Allow SSH and HTTP",
				IngressRules: []models.IngressRule{
					{Protocol: "tcp", FromPort: 22, ToPort: 22, CIDRIP: "0.0.0.0/0"},
					{Protocol: "tcp", FromPort: 80, ToPort: 80, CIDRIP: "0.0.0.0/0"},
				},
			},
		}
	}

	config.Networking = networking
	config.Resources = resources
	return config, nil
}

func promptVMProperties() map[string]string {
	properties := make(map[string]string)

	// subnet, securityGroups,
	var vmName, machineType, region, image, keyName string

	survey.AskOne(&survey.Input{Message: "Enter VM name:"}, &vmName)
	properties["name"] = vmName

	survey.AskOne(&survey.Input{Message: "Enter machine type (e.g., t2.micro):"}, &machineType)
	properties["machineType"] = machineType

	survey.AskOne(&survey.Input{Message: "Enter region (e.g., us-east-1):"}, &region)
	properties["region"] = region

	survey.AskOne(&survey.Input{Message: "Enter image ID (e.g., ami-123456):"}, &image)
	properties["image"] = image

	// survey.AskOne(&survey.Input{Message: "Enter subnet ID:"}, &subnet)
	// properties["subnet"] = subnet

	// survey.AskOne(&survey.Input{Message: "Enter security groups (comma-separated):"}, &securityGroups)
	// properties["securityGroups"] = securityGroups

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

// PromptNetworkingChoice prompts the user to choose between default and custom VPC.
func PromptNetworkingChoice() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose VPC configuration:")
	fmt.Println("1. Default VPC")
	fmt.Println("2. Custom VPC")
	fmt.Print("Enter choice (1 or 2): ")
	choice, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	choice = strings.TrimSpace(choice)

	if choice == "1" {
		return "default", nil
	} else if choice == "2" {
		return "custom", nil
	} else {
		return "", fmt.Errorf("invalid choice")
	}
}
