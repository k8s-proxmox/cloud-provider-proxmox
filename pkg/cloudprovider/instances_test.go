package proxmox

import (
	"github.com/sp-yduck/proxmox-go/rest"
)

func (s *TestSuite) TestGetVMFromUUID() {
	uuid := "cefbe890-bc88-4faa-82d1-a33915d71d1d"

	vm, err := s.instance.getVMFromUUID(uuid)
	if err != nil {
		if rest.IsNotFound(err) {
			s.T().Logf("not found vm having uuid=%s", uuid)
			return
		}
		s.T().Fatalf("failed to get vm from uuid: %v", err)

	}
	s.T().Logf("get vm having uuid=%s: %v", uuid, *vm)
}
