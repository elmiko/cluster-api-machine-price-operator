package kubemark

import (
	"context"

	"github.com/elmiko/cluster-api-machine-price-operator/pkg/cloudprovider"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewProvider() cloudprovider.CloudProvider {
	return KubemarkCloudProvider{}
}

type KubemarkCloudProvider struct {
}

func (p KubemarkCloudProvider) NewInfraMachineTemplateObject() client.Object {
	return nil
}

func (p KubemarkCloudProvider) GetPriceForRequest(context.Context, ctrl.Request) float64 {
	return 0.0
}

func (p KubemarkCloudProvider) GetPriceFor(client.Object) float64 {
	return 0.0
}
