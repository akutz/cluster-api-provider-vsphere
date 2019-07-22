/*
Copyright 2019 The Kubernetes Authors.

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

package app

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"

	"sigs.k8s.io/cluster-api-provider-vsphere/cmd/manifests/pkg/encoding/slim"
	"sigs.k8s.io/cluster-api-provider-vsphere/cmd/manifests/pkg/kustomize"
)

// FlagStringSlice may be used with flag.Var to register a flag that may be
// specified multiple times on the command line and return a slice of the args.
type FlagStringSlice []string

// String returns the list as a CSV string.
func (s *FlagStringSlice) String() string {
	return strings.Join(*s, ",")
}

// Set is called once, in command line order, for each flag present.
func (s *FlagStringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

var (
	configDirs        FlagStringSlice
	kubernetesVersion = flag.String(
		"kubernetes-version",
		"v1.13.6",
		"The version of Kubernetes to deploy")
	clusterName = flag.String(
		"cluster-name",
		"management-cluster",
		"The name of the cluster")
	podCIDR = flag.String(
		"pod-cidr",
		"100.96.0.0/11",
		"The CIDR for the cluster's pod network.")
	serviceCIDR = flag.String(
		"service-cidr",
		"100.64.0.0/13",
		"The CIDR for the cluster's service network.")
	serviceDomain = flag.String(
		"service-domain",
		"cluster.local",
		"The domain name for the cluster's service network.")
	clusterOutPath = flag.String(
		"cluster-out",
		"cluster.yaml",
		"The path to write the generated cluster manifest")
	machinesOutPath = flag.String(
		"machines-out",
		"machines.yaml",
		"The path to write the generated machines manifest")
	providerComponentsOutPath = flag.String(
		"provider-components-out",
		"provider-components.yaml",
		"The path to write the generated provider components manifest")
)

func init() {
	flag.Var(&configDirs,
		"config-dir",
		"A directory containing Kustomization resources. May be specified more than once.")
}

// Provider describes a type that can return a Cluster & Machine provider spec.
type Provider interface {
	GetTemplateData() map[string]interface{}
	GetClusterProviderSpec() (runtime.Object, error)
	GetMachineProviderSpec() (runtime.Object, error)
}

// Run is the entry point for the application.
func Run(p Provider) error {
	if !flag.Parsed() {
		flag.Parse()
	}
	if err := generateClusterManifest(p); err != nil {
		return err
	}
	if err := generateMachinesManifest(p); err != nil {
		return err
	}
	if err := generateProviderComponentsManifest(p); err != nil {
		return err
	}
	return nil
}

func generateProviderComponentsManifest(p Provider) error {
	fout, err := os.Create(*providerComponentsOutPath)
	if err != nil {
		return err
	}
	defer fout.Close()
	for i, configDirPath := range configDirs {
		buildOptions := &kustomize.BuildOptions{
			Out:               fout,
			KustomizationPath: configDirPath,
			TemplateData:      p.GetTemplateData(),
		}
		if err := kustomize.RunBuild(buildOptions); err != nil {
			return errors.Wrap(err, "failed to run kustomize")
		}
		if i < len(configDirs)-1 {
			if _, err := fmt.Fprintf(fout, "---\n"); err != nil {
				return err
			}
		}
	}
	return nil
}

func generateClusterManifest(p Provider) error {
	providerSpec, err := p.GetClusterProviderSpec()
	if err != nil {
		return err
	}

	encodedProviderSpec, err := slim.EncodeAsRawExtension(providerSpec)
	if err != nil {
		return err
	}

	obj := &clusterv1.Cluster{
		TypeMeta: metav1.TypeMeta{
			APIVersion: clusterv1.SchemeGroupVersion.String(),
			Kind:       "Cluster",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: *clusterName,
		},
		Spec: clusterv1.ClusterSpec{
			ProviderSpec: clusterv1.ProviderSpec{
				Value: encodedProviderSpec,
			},
			ClusterNetwork: clusterv1.ClusterNetworkingConfig{
				Pods: clusterv1.NetworkRanges{
					CIDRBlocks: []string{*podCIDR},
				},
				ServiceDomain: *serviceDomain,
				Services: clusterv1.NetworkRanges{
					CIDRBlocks: []string{*serviceCIDR},
				},
			},
		},
	}

	return writeObjToFile(obj, *clusterOutPath)
}

func generateMachinesManifest(p Provider) error {
	providerSpec, err := p.GetMachineProviderSpec()
	if err != nil {
		return err
	}

	encodedProviderSpec, err := slim.EncodeAsRawExtension(providerSpec)
	if err != nil {
		return err
	}

	obj := &clusterv1.MachineList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: clusterv1.SchemeGroupVersion.String(),
			Kind:       "MachineList",
		},
		Items: []clusterv1.Machine{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: fmt.Sprintf("%s-controlplane-1", *clusterName),
					Labels: map[string]string{
						"cluster.k8s.io/cluster-name": *clusterName,
					},
				},
				TypeMeta: metav1.TypeMeta{
					APIVersion: clusterv1.SchemeGroupVersion.String(),
					Kind:       "Machine",
				},
				Spec: clusterv1.MachineSpec{
					ProviderSpec: clusterv1.ProviderSpec{
						Value: encodedProviderSpec,
					},
					Versions: clusterv1.MachineVersionInfo{
						Kubelet:      *kubernetesVersion,
						ControlPlane: *kubernetesVersion,
					},
				},
			},
		},
	}

	return writeObjToFile(obj, *machinesOutPath)
}

func writeObjToFile(obj runtime.Object, filePath string) error {
	objYAML, err := slim.MarshalYAML(obj)
	if err != nil {
		return err
	}

	fout, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer fout.Close()

	if _, err := io.Copy(fout, bytes.NewReader(objYAML)); err != nil {
		return err
	}

	return nil
}
