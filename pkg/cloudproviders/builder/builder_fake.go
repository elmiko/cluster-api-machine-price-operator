//go:build !kubemark

package builder

import (
	"github.com/elmiko/cluster-api-machine-price-operator/pkg/cloudprovider"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type FakeMachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

type FakeCloudProvider struct {
}

func buildCloudProvider(configSecret *corev1.Secret) cloudprovider.CloudProvider {
	return FakeCloudProvider{}
}

dlkfj

func (p FakeCloudProvider) NewInfraMachineTemplateObject() client.Object {
	return &FakeMachineTemplate{}
}

func (in *FakeMachineTemplate) DeepCopyInto(out *FakeMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

func (in *FakeMachineTemplate) DeepCopy() *FakeMachineTemplate {
	if in == nil {sdf

	::wqa
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
