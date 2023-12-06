# Kubernetes Proxmox Cloud Provider

## Proxmox Cloud Controller Manager

This repository contains the Kubernetes cloud-controller-manager for Proxmox VE. Proxmox Cloud Controller Manager is compatible with [cluster-api-provider-proxmox](https://github.com/k8s-proxmox/cluster-api-provider-proxmox) so that machine controller of CAPI can match the nodes with their machine object. (See [machine-controller](https://cluster-api.sigs.k8s.io/developer/architecture/controllers/machine.html#machine--controller) for how it works)

## Feature
### Node lifecycle controller

The nodes initialized by proxmox-ccm look like below

```yaml
apiVersion: v1
kind: Node
metadata:
  labels:
    node.kubernetes.io/instance-type: proxmox-qemu.cpu-2.mem-4.0G
  name: worker-1
spec:
  ...
  providerID: proxmox://smbios-system-uuid-ab012345678
```


## Configuration

cloud config looks like below. See [sample manifest](./manifests/cloud-controller-manager.yaml)

```yaml
proxmox:
    url: https://X.X.X.X:8006/api2/json
    tokenID: 'root@pam!api-token-id'
    secret: "aaaaaaaa-bbbb-cccc-dddd-ee12345678"
    # user: user and
    # password: password is also available
```

## Developing
### Integration Testing
```sh
export PROXMOX_URL='http://localhost:8006/api2/json'

# tokenid & secret
export PROXMOX_TOKENID='root@pam!your-token-id'
export PROXMOX_SECRET='aaaaaaaaa-bbb-cccc-dddd-ef0123456789'

# or username & password
# export PROXMOX_USERNAME='root@pam'
# export PROXMOX_PASSWORD='password'

export PROXMOX_TEST_UUID='something-actual-uuid-here'

go test ./... -v -run ^TestSuiteIntegration
```
## License

Copyright 2023 Teppei Sudo.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

