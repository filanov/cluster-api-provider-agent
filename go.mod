module github.com/eranco74/cluster-api-provider-agent

go 1.16

replace sigs.k8s.io/cluster-api => sigs.k8s.io/cluster-api v1.0.0

// Versions to be held for v1beta1
// sigs.k8s.io/controller-runtime on v0.10.x
// k8s.io/* on v0.22.x
// github.com/go-logr/logr on v0.4.x
// k8s.io/klog/v2 on v2.10.x
require (
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.16.0
	github.com/openshift/hive/apis v0.0.0-20211012200111-a691d6f21d9e
	github.com/sirupsen/logrus v1.8.1
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
	sigs.k8s.io/cluster-api v1.0.0
	sigs.k8s.io/controller-runtime v0.10.2
)

replace (
        github.com/metal3-io/baremetal-operator => github.com/openshift/baremetal-operator v0.0.0-20210409032903-31b989a197eb // Use OpenShift fork
        k8s.io/api => k8s.io/api v0.21.1
        k8s.io/client-go => k8s.io/client-go v0.21.1
        sigs.k8s.io/cluster-api-provider-aws => github.com/openshift/cluster-api-provider-aws v0.2.1-0.20201022175424-d30c7a274820
        sigs.k8s.io/cluster-api-provider-azure => github.com/openshift/cluster-api-provider-azure v0.1.0-alpha.3.0.20201016155852-4090a6970205
)
