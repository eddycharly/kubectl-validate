/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "k8s.io/api/extensions/v1beta1"
	extensionsv1beta1 "k8s.io/client-go/applyconfigurations/extensions/v1beta1"
	gentype "k8s.io/client-go/gentype"
	typedextensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

// fakeNetworkPolicies implements NetworkPolicyInterface
type fakeNetworkPolicies struct {
	*gentype.FakeClientWithListAndApply[*v1beta1.NetworkPolicy, *v1beta1.NetworkPolicyList, *extensionsv1beta1.NetworkPolicyApplyConfiguration]
	Fake *FakeExtensionsV1beta1
}

func newFakeNetworkPolicies(fake *FakeExtensionsV1beta1, namespace string) typedextensionsv1beta1.NetworkPolicyInterface {
	return &fakeNetworkPolicies{
		gentype.NewFakeClientWithListAndApply[*v1beta1.NetworkPolicy, *v1beta1.NetworkPolicyList, *extensionsv1beta1.NetworkPolicyApplyConfiguration](
			fake.Fake,
			namespace,
			v1beta1.SchemeGroupVersion.WithResource("networkpolicies"),
			v1beta1.SchemeGroupVersion.WithKind("NetworkPolicy"),
			func() *v1beta1.NetworkPolicy { return &v1beta1.NetworkPolicy{} },
			func() *v1beta1.NetworkPolicyList { return &v1beta1.NetworkPolicyList{} },
			func(dst, src *v1beta1.NetworkPolicyList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.NetworkPolicyList) []*v1beta1.NetworkPolicy {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1beta1.NetworkPolicyList, items []*v1beta1.NetworkPolicy) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
