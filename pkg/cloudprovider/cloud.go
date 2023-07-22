package proxmox

import (
	"io"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
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
	instancesV2 cloudprovider.InstancesV2
}

type cloudProviderConfig struct {
	ProxmoxConfig proxmoxConfig `yaml:"proxmox"`
}

type proxmoxConfig struct {
	URL      string `yaml:"url"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	TokenID  string `yaml:"tokenID"`
	Secret   string `yaml:"secret"`
}

func init() {
	klog.Info("registering cloud provider")

	cloudprovider.RegisterCloudProvider(RegisteredProviderName, func(config io.Reader) (cloudprovider.Interface, error) {
		providerConfig, err := readCloudProviderConfig(config)
		if err != nil {
			return nil, err
		}
		return newCloud(providerConfig)
	})
}

func newCloud(config *cloudProviderConfig) (cloudprovider.Interface, error) {
	klog.Info("creating new cloud")
	instance, err := newInstances(config.ProxmoxConfig)
	if err != nil {
		return nil, err
	}
	px := &Proxmox{instancesV2: instance}
	return px, nil
}

func readCloudProviderConfig(configReader io.Reader) (*cloudProviderConfig, error) {
	config := &cloudProviderConfig{}
	if configReader == nil {
		return nil, errors.New("configReader must not be nil")
	}
	if err := yaml.NewDecoder(configReader).Decode(config); err != nil {
		return nil, err
	}
	cfg := config.ProxmoxConfig
	if cfg.URL == "" {
		return nil, errors.New("url must not be empty")
	}
	return config, nil
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
	return nil, false
}

// Instances returns an instances interface. Also returns true if the
// interface is supported, false otherwise.
func (px *Proxmox) InstancesV2() (cloudprovider.InstancesV2, bool) {
	return px.instancesV2, true
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

// HasClusterID returns true if a ClusterID is required and set/
func (px *Proxmox) HasClusterID() bool {
	return true
}
