module sigs.k8s.io/cluster-api-provider-vsphere

go 1.12

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20181127025237-2b1284ed4c93

require (
	github.com/go-logr/logr v0.1.0
	github.com/golang/groupcache v0.0.0-20190129154638-5b532d6fd5ef // indirect
	github.com/google/uuid v1.1.1
	github.com/gophercloud/gophercloud v0.1.0 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/mbenkmann/goformat v0.0.0-20180512004123-256ef38c4271
	github.com/onsi/gomega v1.5.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v0.9.3 // indirect
	github.com/spf13/cobra v0.0.4 // indirect
	github.com/vmware/govmomi v0.20.2
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/lint v0.0.0-20190301231843-5614ed5bae6f
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	golang.org/x/oauth2 v0.0.0-20190523182746-aaccbc9213b0 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	golang.org/x/tools v0.0.0-20190614205625-5aca471b1d59
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apimachinery v0.0.0-20190703205208-4cfb76a8bf76
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/cluster-bootstrap v0.0.0-20190307184522-e5eaaafa58b3
	k8s.io/code-generator v0.0.0-20190308034351-e797d15e3d1a
	k8s.io/gengo v0.0.0-20190308184658-b90029ef6cd8 // indirect
	k8s.io/klog v0.3.3
	k8s.io/kubernetes v1.13.3
	sigs.k8s.io/cluster-api v0.1.7
	sigs.k8s.io/controller-runtime v0.1.12
	sigs.k8s.io/controller-tools v0.1.11
	sigs.k8s.io/kustomize/v3 v3.0.3
	sigs.k8s.io/yaml v1.1.0
	winterdrache.de/goformat v0.0.0-20180512004123-256ef38c4271 // indirect
)
