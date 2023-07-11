package proxmox

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	cloudprovider "k8s.io/cloud-provider"
)

type instance struct {
}

func newInstances() cloudprovider.Instances {
	return &instance{}
}

func (i *instance) NodeAddresses(ctx context.Context, name types.NodeName) ([]v1.NodeAddress, error) {
	return []v1.NodeAddress{}, nil
}

func (i *instance) AddSSHKeyToAllInstances(ctx context.Context, user string, keyData []byte) error {
	return cloudprovider.NotImplemented
}

func (i *instance) CurrentNodeName(ctx context.Context, hostname string) (types.NodeName, error) {
	return types.NodeName(hostname), nil
}

func (i *instance) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
	return true, nil
}

func (i *instance) InstanceID(ctx context.Context, nodeName types.NodeName) (string, error) {
	return "a;asdf", nil
}

func (i *instance) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
	return true, nil
}

func (i *instance) InstanceType(ctx context.Context, name types.NodeName) (string, error) {
	return ";lkj", nil
}

func (i *instance) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
	return "asdf", nil
}

func (i *instance) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
	return []v1.NodeAddress{}, nil
}
