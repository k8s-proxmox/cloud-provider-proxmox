package proxmox

import (
	"os"
	"testing"

	"github.com/k8s-proxmox/proxmox-go/proxmox"
	"github.com/stretchr/testify/suite"
	v1 "k8s.io/api/core/v1"
)

type TestSuite struct {
	suite.Suite
	instance instance
	node     v1.Node
}

func (s *TestSuite) SetupSuite() {
	s.setupInstance()
	s.setupTestNode()
}

func (s *TestSuite) setupInstance() {
	url := os.Getenv("PROXMOX_URL")
	user := os.Getenv("PROXMOX_USERNAME")
	password := os.Getenv("PROXMOX_PASSWORD")
	tokeid := os.Getenv("PROXMOX_TOKENID")
	secret := os.Getenv("PROXMOX_SECRET")
	if url == "" {
		s.T().Fatal("url must not be empty")
	}
	authConfig := proxmox.AuthConfig{
		Username: user,
		Password: password,
		TokenID:  tokeid,
		Secret:   secret,
	}

	params := proxmox.NewParams(url, authConfig, proxmox.ClientConfig{InsecureSkipVerify: true})
	svc, err := proxmox.NewService(params)
	if err != nil {
		s.T().Logf("username=%s, password=%s, tokenid=%s, secret=%s", user, password, tokeid, secret)
		s.T().Fatalf("failed to create rest client: %v", err)
	}
	s.instance.compute = svc
}

func (s *TestSuite) setupTestNode() {
	uuid := os.Getenv("PROXMOX_TEST_UUID")
	node := v1.Node{
		Status: v1.NodeStatus{
			NodeInfo: v1.NodeSystemInfo{
				SystemUUID: uuid,
			},
		},
	}
	node.SetName("test-node")
	s.node = node
}

func TestSuiteIntegration(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
