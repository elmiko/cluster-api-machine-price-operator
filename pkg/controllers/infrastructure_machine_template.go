package controllers

import (
	"context"

	"github.com/elmiko/cluster-api-machine-price-operator/pkg/cloudprovider"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type InfraMachineTemplateReconciler struct {
	client.Client
	provider cloudprovider.CloudProvider
}

func NewInfraMachineTemplateReconciler(client client.Client, provider cloudprovider.CloudProvider) InfraMachineTemplateReconciler {
	return InfraMachineTemplateReconciler{
		Client:   client,
		provider: provider,
	}
}

func (r *InfraMachineTemplateReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// check cache for price, and last update time

	// if not cached, or last update time greater than 12 hours, refresh
	infraMachine := r.provider.NewInfraMachineTemplateObject()
	if err := r.Get(ctx, req.NamespacedName, infraMachine); err != nil {
		//price := r.provider.GetPriceFor(infraMachine)

	}
	/*
		typedObject := r.provider.ObjectTypeRef()
		err := cl.Get(... , typedObject)
		// price := provider.GetPriceForRequest(ctx, req)
		// update price on machine template
		// this might be an unstructured load of the object in question
		typedObject.GetLabels()
		update labels
		r.Update(ctx, ... , infraMachine)
		write
	*/
	return ctrl.Result{}, nil
}

func (r *InfraMachineTemplateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(r.provider.NewInfraMachineTemplateObject()).
		Complete(r)
}
