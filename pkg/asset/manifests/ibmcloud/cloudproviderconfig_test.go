package ibmcloud

import (
	"testing"

	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/ibmcloud"
	"github.com/stretchr/testify/assert"
)

func TestCloudProviderConfig(t *testing.T) {
	expectedConfig := `[global]
version = 1.1.0
[kubernetes]
config-file = /mnt/etc/kubernetes/controller-manager-kubeconfig
[load-balancer-deployment]
image = [REGISTRY]/[NAMESPACE]/keepalived:[TAG]
application = keepalived
vlan-ip-config-map = ibm-cloud-provider-vlan-ip-config
[provider]
accountID = 1e1f75646aef447814a6d907cc83fb3c
clusterID = ocp4-8pxks

`

	ic := types.InstallConfig{
		Platform: types.Platform{
			IBMCloud: &ibmcloud.Platform{
				CISInstanceCRN: "crn:v1:bluemix:public:internet-svcs:us-south:a/1e1f75646aef447814a6d907cc83fb3c:instance::",
			},
		},
	}

	actualConfig, err := CloudProviderConfig("ocp4-8pxks", ic)
	assert.NoError(t, err, "failed to create cloud provider config")
	assert.Equal(t, expectedConfig, actualConfig, "unexpected cloud provider config")
}
