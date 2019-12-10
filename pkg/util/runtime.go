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

package util

import (
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
)

// CompareObjects compares two objects.
// An empty string is returned if the objects are equal, otherwise the output
// from the diff is returned.
func CompareObjects(a, b interface{}) (string, error) {
	return cmp.Diff(a, b), nil
}

// CompareRuntimeObjects compares two runtime objects as unstructured data.
// An empty string is returned if the objects are equal, otherwise the output
// from the diff is returned.
func CompareRuntimeObjects(a, b runtime.Object, statusOnly bool) (string, error) {
	au, err := runtime.DefaultUnstructuredConverter.ToUnstructured(a)
	if err != nil {
		return "", errors.Wrapf(err, "failed to convert a=%T to unstructured object", a)
	}
	bu, err := runtime.DefaultUnstructuredConverter.ToUnstructured(b)
	if err != nil {
		return "", errors.Wrapf(err, "failed to convert b=%T to unstructured object", b)
	}
	if statusOnly {
		return cmp.Diff(au["status"], bu["status"]), nil
	}
	return cmp.Diff(au, bu), nil
}
