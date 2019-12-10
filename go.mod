module sigs.k8s.io/cluster-api-provider-vsphere

go 1.12

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/go-logr/logr v0.1.0
	github.com/google/go-cmp v0.3.1
	github.com/google/uuid v1.0.0
	github.com/onsi/ginkgo v1.10.1
	github.com/onsi/gomega v1.7.0
	github.com/pkg/errors v0.8.1
	github.com/vmware/govmomi v0.21.0
	gopkg.in/gcfg.v1 v1.2.3
	gopkg.in/warnings.v0 v0.1.2 // indirect
	k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/client-go v0.0.0-20190918200256-06eb1244587a
	k8s.io/klog v0.4.0
	sigs.k8s.io/cluster-api v0.2.7
	sigs.k8s.io/cluster-api-bootstrap-provider-kubeadm v0.1.5
	sigs.k8s.io/controller-runtime v0.3.0
)
