package cloudproviders

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CloudProvider interface {
	// NewInfraMachineTemplateObject returns an Object of the infrastructure machine template type
	// that will be reconciled by the controller.
	NewInfraMachineTemplateObject() client.Object

	// GetPriceForRequest returns the floating point value for the price of the request as determined
	// by the cloud provider.
	GetPriceForRequest(context.Context, ctrl.Request)

	// GetPriceFor returns the floating point value for the price for the object as determined
	// by the cloud provider.
	GetPriceFor(client.Object) float64
}
