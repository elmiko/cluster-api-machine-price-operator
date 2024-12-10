package cloudprovider

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CloudProvider interface {
	// NewInfraMachineTemplateObject returns an Object of the infrastructure machine template type
	// that will be reconciled by the controller.
	NewInfraMachineTemplateObject() client.Object

	// GetPriceForRequest returns the floating point value for the price as determined
	// by the cloud provider.
	GetPriceForRequest(context.Context, ctrl.Request)

	GetPriceFor(client.Object) float64 
}
