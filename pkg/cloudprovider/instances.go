package proxmox

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/sp-yduck/proxmox-go/rest"
	v1 "k8s.io/api/core/v1"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog/v2"
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
	klog.Info("checking if instance exists")
	return true, nil
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
