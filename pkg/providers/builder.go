package providers

import (
	"github.com/elmiko/cluster-api-machine-price-operator/pkg/cloudprovider"
	"github.com/elmiko/cluster-api-machine-price-operator/pkg/options"
	"github.com/elmiko/cluster-api-machine-price-operator/pkg/providers/fake"
	"github.com/elmiko/cluster-api-machine-price-operator/pkg/providers/kubemark"
)

const (
	FakeProviderName     = "fake"
	KubemarkProviderName = "kubemark"
)

var AvailableCloudProviders = []string{
	FakeProviderName,
	KubemarkProviderName,
}

// NewCloudProvider returns a new CloudProvider based on the configuration options.
func NewCloudProvider(opts options.CampoOptions) cloudprovider.CloudProvider {
	switch opts.ProviderName {
	case FakeProviderName:
		return fake.NewProvider()
	case KubemarkProviderName:
		return kubemark.NewProvider()
	}
	return nil
}
