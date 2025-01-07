package providers

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

type AzureProvider struct {
	vmClient *armcompute.VirtualMachinesClient
}
