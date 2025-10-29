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

package providers

import (
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type UnknownInfrastructureRef struct{}

func (e UnknownInfrastructureRef) Error() string {
	return "unknown infrastructure reference type"
}

type InfrastructurePriceProvider struct {
	client client.Client
}

func NewInfrastructurePriceProvider(cl client.Client) InfrastructurePriceProvider {
	return InfrastructurePriceProvider{
		cl,
	}
}

func (p InfrastructurePriceProvider) GetPriceFor(ref corev1.ObjectReference) (float64, error) {
	return 0.0, UnknownInfrastructureRef{}
}
