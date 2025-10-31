/*
Copyright 2025. Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"

	"github.com/elmiko/cluster-api-machine-price-operator/pkg/providers"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	capiv1beta1 "sigs.k8s.io/cluster-api/api/core/v1beta1"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var log = logf.Log.WithName("machinedeployment-reconciler")

const (
	clusterApiPriceAnnotation = "cluster.x-k8s.io/machine-current-price"
)

type MachineDeploymentReconciler struct {
	client.Client

	priceProvider providers.InfrastructurePriceProvider
}

func NewMachineDeploymentReconciler(cl client.Client) *MachineDeploymentReconciler {
	return &MachineDeploymentReconciler{
		cl,
		providers.NewInfrastructurePriceProvider(cl),
	}
}

func (r *MachineDeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var machineDeployment capiv1beta1.MachineDeployment
	if err := r.Get(ctx, req.NamespacedName, &machineDeployment); err != nil {
		log.Error(err, "unable to fetch MachineDeployment", "name", machineDeployment.Name, "namespace", machineDeployment.Namespace)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	helper, err := patch.NewHelper(&machineDeployment, r.Client)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to init patch helper: %w", err)
	}

	shouldUpdate := false
	infraRef := machineDeployment.Spec.Template.Spec.InfrastructureRef
	if price, found, err := r.priceProvider.GetPriceFor(infraRef); err != nil {
		if _, ok := err.(providers.UnknownInfrastructureRefError); ok {
			log.Info("no provider found for infra ref", "kind", infraRef.Kind)
			return ctrl.Result{}, nil
		}
		log.Error(err, "unexpected error getting price data")
		return ctrl.Result{}, err
	} else {
		if found {
			log.Info("got price", "value", price)
			// add annotation to the machineDeployment
			if machineDeployment.Annotations == nil {
				machineDeployment.Annotations = map[string]string{}
			}
			machineDeployment.Annotations[clusterApiPriceAnnotation] = fmt.Sprintf("%7.4f", price)
			shouldUpdate = true
		}
	}

	if shouldUpdate {
		if err := helper.Patch(ctx, &machineDeployment); err != nil {
			if !apierrors.IsNotFound(err) {
				log.Error(err, "failed to patch machineDeployment")
				return ctrl.Result{}, err
			}
		}
	}

	return ctrl.Result{}, nil
}

func (r *MachineDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&capiv1beta1.MachineDeployment{}).
		Complete(r)
}
