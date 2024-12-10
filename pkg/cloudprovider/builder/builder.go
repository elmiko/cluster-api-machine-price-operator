package builder

import (
	"github.com/elmiko/cluster-api-machine-price-operator/pkg/cloudprovider"
	corev1 "k8s.io/api/core/v1"
)

// NewCloudProvider will return a CloudProvider associated with the cloud that was
// selected at compile time through build tags. This function hides the internal
// implementation from callers outside the package.
func NewCloudProvider(configSecret *corev1.Secret) cloudprovider.CloudProvider {
	return buildCloudProvider(configSecret)
}
