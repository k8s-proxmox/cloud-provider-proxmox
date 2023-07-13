package proxmox

import (
	"context"
	"fmt"

	"github.com/sp-yduck/proxmox/pkg/service"
	// "github.com/sp-yduck/proxmox/pkg/service/node/vm"
	v1 "k8s.io/api/core/v1"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog/v2"
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
	klog.Info("checking if instance exists")
	return true, nil
}

func (i *instance) InstanceShutdown(ctx context.Context, node *v1.Node) (bool, error) {
	klog.Info("checking if instance is shutdowned")
	return false, nil
}

func (i *instance) InstanceMetadata(ctx context.Context, node *v1.Node) (*cloudprovider.InstanceMetadata, error) {
	providerID := fmt.Sprintf("%s://%s", ProviderName, node.Status.NodeInfo.SystemUUID)
	klog.Infof("initializing node %s with providerID=%s", node.Name, providerID)
	return &cloudprovider.InstanceMetadata{
		ProviderID:    providerID,
		NodeAddresses: []v1.NodeAddress{},
		InstanceType:  "",
		Zone:          "",
		Region:        "",
	}, nil
}
