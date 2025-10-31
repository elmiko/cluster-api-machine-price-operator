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

package kubemark

import (
	"context"
	"strconv"

	corev1 "k8s.io/api/core/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	kubemarkv1 "sigs.k8s.io/cluster-api-provider-kubemark/api/v1alpha4"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

const InfrastructureRefKind = "KubemarkMachineTemplate"

var log = logf.Log.WithName("kubemark-price-provider")

func GetPriceFor(ctx context.Context, cl client.Client, ref corev1.ObjectReference) (float64, bool, error) {
	utilruntime.Must(kubemarkv1.AddToScheme(cl.Scheme()))
	key := client.ObjectKey{Namespace: ref.Namespace, Name: ref.Name}
	template := kubemarkv1.KubemarkMachineTemplate{}
	if err := cl.Get(ctx, key, &template); err != nil {
		log.Error(err, "unable to fetch KubemarkMachineTemplate", "name", key.Name, "namespace", key.Namespace)
		return 0.0, false, err
	}

	if len(template.Status.Prices.Current) != 0 {
		if price, err := strconv.ParseFloat(template.Status.Prices.Current, 64); err != nil {
			return 0.0, false, err
		} else {
			return price, true, nil
		}
	}

	return 0.0, false, nil
}
