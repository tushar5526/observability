package prometheus

import (
	"github.com/gitpod-io/observability/installer/pkg/common"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// extraNamespaceRoleBindings and extraNamespaceRoles are used to give permission to prometheus to scrape metrics
// from endpoints in other namespaces.
// TODO: Add more namespaces from configuration
func extraNamespaceRoleBindings(ctx *common.RenderContext) ([]runtime.Object, error) {
	return []runtime.Object{
		rolebindingFactory(Namespace),
		rolebindingFactory("default"),
		rolebindingFactory("kube-system"),
	}, nil
}

func rolebindingFactory(ns string) *rbacv1.RoleBinding {
	return &rbacv1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "RoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resourceName(),
			Namespace: ns,
			Labels:    common.Labels(Name, Component, App, Version),
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     resourceName(),
		},
		Subjects: []rbacv1.Subject{
			{
				Kind: "ServiceAccount",
				Name: resourceName(),
				// Here we associate the service account used by prometheus
				// which lives in the same namespace as prometheus, and not the role.
				Namespace: Namespace,
			},
		},
	}
}