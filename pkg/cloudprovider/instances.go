package proxmox

import (
	"context"
	"fmt"

	"github.com/sp-yduck/proxmox/pkg/service"
	// "github.com/sp-yduck/proxmox/pkg/service/node/vm"
	v1 "k8s.io/api/core/v1"
	cloudprovider "k8s.io/cloud-provider"
)

type instance struct {
	compute *service.Service
}

func newInstances(config proxmoxConfig) (cloudprovider.InstancesV2, error) {
	svc, err := service.NewServiceWithLogin(config.URL, config.User, config.Password)
	if err != nil {
		return nil, err
	}
	return &instance{compute: svc}, nil
}

func (i *instance) InstanceExists(ctc context.Context, node *v1.Node) (bool, error) {
	return true, nil
}

func (i *instance) InstanceShutdown(ctx context.Context, node *v1.Node) (bool, error) {
	return false, nil
}

func (i *instance) InstanceMetadata(ctx context.Context, node *v1.Node) (*cloudprovider.InstanceMetadata, error) {
	providerID := fmt.Sprintf("%s://%s", ProviderName, node.Status.NodeInfo.SystemUUID)
	return &cloudprovider.InstanceMetadata{
		ProviderID: providerID,
	}, nil
}
