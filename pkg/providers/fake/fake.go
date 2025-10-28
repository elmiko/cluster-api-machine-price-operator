package fake

import (
	"context"

	"github.com/elmiko/cluster-api-machine-price-operator/pkg/cloudprovider"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type FakeMachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

type FakeCloudProvider struct {
}

func NewProvider() cloudprovider.CloudProvider {
	return FakeCloudProvider{}
}

func (p FakeCloudProvider) NewInfraMachineTemplateObject() client.Object {
	return &FakeMachineTemplate{}
}

func (p FakeCloudProvider) GetPriceForRequest(context.Context, ctrl.Request) float64 {
	return 0.0
}

func (p FakeCloudProvider) GetPriceFor(client.Object) float64 {
	return 0.0
}

func (in *FakeMachineTemplate) DeepCopyInto(out *FakeMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

func (in *FakeMachineTemplate) DeepCopy() *FakeMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(FakeMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

func (in *FakeMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
