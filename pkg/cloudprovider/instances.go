package proxmox

import (
	"context"
	"fmt"

	"github.com/sp-yduck/proxmox-go/api"
	"github.com/sp-yduck/proxmox-go/proxmox"
	"github.com/sp-yduck/proxmox-go/rest"
	v1 "k8s.io/api/core/v1"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog/v2"
)

const (
	Giga = 1024 * 1024 * 1024
)

type instance struct {
	compute *proxmox.Service
}

func newInstances(config proxmoxConfig) (cloudprovider.InstancesV2, error) {
	authConfig := proxmox.AuthConfig{
		Username: config.User,
		Password: config.Password,
		TokenID:  config.TokenID,
		Secret:   config.Secret,
	}
	client, err := proxmox.NewService(config.URL, authConfig, true)
	if err != nil {
		return nil, err
	}
	return &instance{compute: client}, nil
}

func (i *instance) InstanceExists(ctx context.Context, node *v1.Node) (bool, error) {
	klog.Infof("checking if instance exists (node=%s)", node.Name)

	_, err := i.compute.VirtualMachineFromUUID(ctx, node.Status.NodeInfo.SystemUUID)
	if err != nil {
		if rest.IsNotFound(err) {
			return false, nil
		}
		return true, err
	}

	return true, nil
}

func (i *instance) InstanceShutdown(ctx context.Context, node *v1.Node) (bool, error) {
	klog.V(2).Info("InstanceShutdown called")

	vm, err := i.compute.VirtualMachineFromUUID(ctx, node.Status.NodeInfo.SystemUUID)
	if err != nil {
		return false, err
	}

	shutdonw := vm.VM.Status == api.ProcessStatusStopped
	return shutdonw, nil
}

func (i *instance) InstanceMetadata(ctx context.Context, node *v1.Node) (*cloudprovider.InstanceMetadata, error) {
	providerID := fmt.Sprintf("%s://%s", ProviderName, node.Status.NodeInfo.SystemUUID)
	klog.Infof("getting metadata for node %s (providerID=%s)", node.Name, providerID)

	vm, err := i.compute.VirtualMachineFromUUID(ctx, node.Status.NodeInfo.SystemUUID)
	if err != nil {
		return nil, err
	}

	cpu := vm.VM.Cpus
	mem := roundBtoGB(vm.VM.MaxMem)
	instanceType := fmt.Sprintf("proxmox-qemu.cpu-%d.mem-%s", cpu, mem)

	return &cloudprovider.InstanceMetadata{
		ProviderID:    providerID,
		NodeAddresses: []v1.NodeAddress{},
		InstanceType:  instanceType,
		Zone:          "",
		Region:        "",
	}, nil
}

func roundBtoGB(size int) string {
	rounded := float32(size) / Giga
	return fmt.Sprintf("%.1fG", rounded)
}
