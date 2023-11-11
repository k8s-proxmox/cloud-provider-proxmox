package proxmox

import (
	"context"
)

func (s *TestSuite) TestInstanceExists() {
	exists, err := s.instance.InstanceExists(context.TODO(), &s.node)
	if err != nil {
		s.T().Fatalf("failed get instance: %v", err)
	}
	s.T().Logf("get instance: %v", exists)
}

func (s *TestSuite) TestInstanceMetadata() {
	meta, err := s.instance.InstanceMetadata(context.TODO(), &s.node)
	if err != nil {
		s.T().Fatalf("failed get instance metadata: %v", err)
	}
	s.T().Logf("get instance metadata: %v", *meta)
}

func (s *TestSuite) TestInstanceShutdown() {
	shutdown, err := s.instance.InstanceShutdown(context.TODO(), &s.node)
	if err != nil {
		s.T().Fatalf("failed get shutdown status: %v", err)
	}
	s.T().Logf("get shutdown status: %v", shutdown)
}
