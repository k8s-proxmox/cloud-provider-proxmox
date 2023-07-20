package proxmox

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/sp-yduck/proxmox-go/api"
	"github.com/sp-yduck/proxmox-go/rest"
	v1 "k8s.io/api/core/v1"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog/v2"
)

const (
	UUIDFormat = `[a-f\d]{8}-[a-f\d]{4}-[a-f\d]{4}-[a-f\d]{4}-[a-f\d]{12}`
)

type instance struct {
	compute *rest.RESTClient
}

func newInstances(config proxmoxConfig) (cloudprovider.InstancesV2, error) {
	base := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	client, err := rest.NewRESTClient(config.URL, rest.WithUserPassword(config.User, config.Password), rest.WithClient(base))
	if err != nil {
		return nil, err
	}
	return &instance{compute: client}, nil
}

func (i *instance) InstanceExists(ctc context.Context, node *v1.Node) (bool, error) {
	klog.Info("checking if instance exists (node=%s)", node.Name)

	_, err := i.getVMFromUUID(node.Status.NodeInfo.SystemUUID)
	if err != nil {
		if rest.IsNotFound(err) {
			return false, nil
		}
		return true, err
	}

	return true, nil
}

func (i *instance) getVMFromUUID(uuid string) (*api.VirtualMachine, error) {
	nodes, err := i.compute.GetNodes()
	if err != nil {
		return nil, err
	}
	for _, n := range nodes {
		vms, err := i.compute.GetVirtualMachines(n.Node)
		if err != nil {
			return nil, err
		}
		for _, vm := range vms {
			config, err := i.compute.GetVirtualMachineConfig(n.Node, vm.VMID)
			if err != nil {
				return nil, err
			}
			smbios := config.SMBios1
			vmuuid, err := convertSMBiosToUUID(smbios)
			if err != nil {
				return nil, err
			}
			if vmuuid == uuid {
				return vm, nil
			}
		}
	}
	return nil, rest.NotFoundErr
}

func convertSMBiosToUUID(smbios string) (string, error) {
	re := regexp.MustCompile(fmt.Sprintf("uuid=%s", UUIDFormat))
	match := re.FindString(smbios)
	if match == "" {
		return "", errors.New("failed to fetch uuid form smbios")
	}
	// match: uuid=<uuid>
	return strings.Split(match, "=")[1], nil
}

func (i *instance) InstanceShutdown(ctx context.Context, node *v1.Node) (bool, error) {
	klog.Info("checking if instance is shutdowned")
	return false, nil
}

func (i *instance) InstanceMetadata(ctx context.Context, node *v1.Node) (*cloudprovider.InstanceMetadata, error) {
	providerID := fmt.Sprintf("%s://%s", ProviderName, node.Status.NodeInfo.SystemUUID)
	klog.Infof("getting metadata for node %s (providerID=%s)", node.Name, providerID)
	return &cloudprovider.InstanceMetadata{
		ProviderID:    providerID,
		NodeAddresses: []v1.NodeAddress{},
		InstanceType:  "",
		Zone:          "",
		Region:        "",
	}, nil
}
