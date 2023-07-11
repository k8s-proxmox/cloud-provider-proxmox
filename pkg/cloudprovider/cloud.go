package proxmox

import (
	"io"

	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/klog/v2"
)

const (
	// RegisteredProviderName is the name of the cloud provider registered with
	// Kubernetes.
	RegisteredProviderName string = "proxmox"

	// ProviderName is the name used for constructing Provider ID
	ProviderName string = "proxmox"

	// ClientName is the user agent passed into the controller client builder.
	ClientName string = "proxmox-cloud-controller-manager"

	// dualStackFeatureGateEnv is a required environment variable when enabling dual-stack nodes
	// dualStackFeatureGateEnv string = "ENABLE_ALPHA_DUAL_STACK"
)

type Proxmox struct {
	instances cloudprovider.Instances
}

func init() {
	klog.Info("registering cloud provider")
	cloudprovider.RegisterCloudProvider(RegisteredProviderName, func(config io.Reader) (cloudprovider.Interface, error) {
		return newCloud(config)
	})
}

func newCloud(configReader io.Reader) (cloudprovider.Interface, error) {
	klog.Info("creating new cloud")
	return newProxmox()
}

func newProxmox() (*Proxmox, error) {
	px := &Proxmox{}
	return px, nil
}

func (px *Proxmox) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
}

// LoadBalancer returns a balancer interface. Also returns true if the
// interface is supported, false otherwise.
func (px *Proxmox) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return nil, false
}

// Instances returns an instances interface. Also returns true if the
// interface is supported, false otherwise.
func (px *Proxmox) Instances() (cloudprovider.Instances, bool) {
	return px.instances, true
}

// Instances returns an instances interface. Also returns true if the
// interface is supported, false otherwise.
func (px *Proxmox) InstancesV2() (cloudprovider.InstancesV2, bool) {
	return nil, false
}

// Zones returns a zones interface. Also returns true if the interface
// is supported, false otherwise.
func (px *Proxmox) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}

// Clusters returns a clusters interface.  Also returns true if the interface
// is supported, false otherwise.
func (px *Proxmox) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

// Routes returns a routes interface along with whether the interface
// is supported.
func (px *Proxmox) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

// ProviderName returns the cloud provider ID.
func (px *Proxmox) ProviderName() string {
	return ProviderName
}

// ScrubDNS
func (px *Proxmox) ScrubDNS(nameservers, searches []string) (nsOut, srchOut []string) {
	return nil, nil
}

// HasClusterID returns true if a ClusterID is required and set/
func (px *Proxmox) HasClusterID() bool {
	return true
}
