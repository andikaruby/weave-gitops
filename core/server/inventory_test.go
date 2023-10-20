package server_test

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	helmv2 "github.com/fluxcd/helm-controller/api/v2beta1"
	kustomizev1 "github.com/fluxcd/kustomize-controller/api/v1"
	"github.com/fluxcd/pkg/apis/meta"
	sourcev1 "github.com/fluxcd/source-controller/api/v1"
	. "github.com/onsi/gomega"
	"github.com/weaveworks/weave-gitops/core/clustersmngr/cluster"
	"github.com/weaveworks/weave-gitops/core/server"
	"github.com/weaveworks/weave-gitops/core/server/types"
	pb "github.com/weaveworks/weave-gitops/pkg/api/core"
	"github.com/weaveworks/weave-gitops/pkg/kube"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestGetInventoryKustomization(t *testing.T) {
	g := NewGomegaWithT(t)

	ctx := context.Background()

	automationName := "my-automation"

	ns := corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-namespace",
			Labels: map[string]string{
				"toolkit.fluxcd.io/tenant": "tenant",
			},
		},
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-deployment",
			Namespace: ns.Name,
			UID:       "this-is-not-an-uid",
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					types.AppLabel: automationName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{types.AppLabel: automationName},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  "nginx",
						Image: "nginx",
					}},
				},
			},
		},
	}

	rs := &appsv1.ReplicaSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-123abcd", automationName),
			Namespace: ns.Name,
		},
		Spec: appsv1.ReplicaSetSpec{
			Template: deployment.Spec.Template,
			Selector: deployment.Spec.Selector,
		},
		Status: appsv1.ReplicaSetStatus{
			Replicas: 1,
		},
	}

	rs.SetOwnerReferences([]metav1.OwnerReference{{
		UID:        deployment.UID,
		APIVersion: appsv1.SchemeGroupVersion.String(),
		Kind:       "Deployment",
		Name:       deployment.Name,
	}})

	kust := &kustomizev1.Kustomization{
		ObjectMeta: metav1.ObjectMeta{
			Name:      automationName,
			Namespace: ns.Name,
		},
		Spec: kustomizev1.KustomizationSpec{
			SourceRef: kustomizev1.CrossNamespaceSourceReference{
				Kind: sourcev1.GitRepositoryKind,
			},
		},
		Status: kustomizev1.KustomizationStatus{
			Inventory: &kustomizev1.ResourceInventory{
				Entries: []kustomizev1.ResourceRef{
					{
						ID:      fmt.Sprintf("%s_%s_apps_Deployment", ns.Name, deployment.Name),
						Version: "v1",
					},
				},
			},
		},
	}

	scheme, err := kube.CreateScheme()
	g.Expect(err).To(BeNil())

	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(&ns, kust, deployment, rs).Build()
	cfg := makeServerConfig(client, t, "")
	c := makeServer(cfg, t)

	res, err := c.GetInventory(ctx, &pb.GetInventoryRequest{
		Namespace:    ns.Name,
		ClusterName:  cluster.DefaultCluster,
		Kind:         "Kustomization",
		Name:         kust.Name,
		WithChildren: true,
	})

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(res.Entries).To(HaveLen(1))

	g.Expect(res.Entries[0].Children).To(HaveLen(1))
	g.Expect(res.Entries[0].Tenant).To(Equal("tenant"))
}

func TestGetBlankInventoryKustomization(t *testing.T) {
	g := NewGomegaWithT(t)

	ctx := context.Background()

	automationName := "my-automation"
	ns := "test-namespace"

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-deployment",
			Namespace: ns,
			UID:       "this-is-not-an-uid",
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					types.AppLabel: automationName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{types.AppLabel: automationName},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  "nginx",
						Image: "nginx",
					}},
				},
			},
		},
	}

	kust := &kustomizev1.Kustomization{
		ObjectMeta: metav1.ObjectMeta{
			Name:      automationName,
			Namespace: ns,
		},
		Spec: kustomizev1.KustomizationSpec{
			SourceRef: kustomizev1.CrossNamespaceSourceReference{
				Kind: sourcev1.GitRepositoryKind,
			},
		},
		Status: kustomizev1.KustomizationStatus{
			Inventory: nil, // blank inventory
		},
	}

	scheme, err := kube.CreateScheme()
	g.Expect(err).To(BeNil())

	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(kust, deployment).Build()
	cfg := makeServerConfig(client, t, "")
	c := makeServer(cfg, t)

	res, err := c.GetInventory(ctx, &pb.GetInventoryRequest{
		Namespace:    ns,
		ClusterName:  cluster.DefaultCluster,
		Kind:         "Kustomization",
		Name:         kust.Name,
		WithChildren: true,
	})

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(res.Entries).To(HaveLen(0))
}

func TestGetInventoryHelmRelease(t *testing.T) {
	g := NewGomegaWithT(t)

	scheme, err := kube.CreateScheme()
	g.Expect(err).NotTo(HaveOccurred())

	ctx := context.Background()

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-namespace",
		},
	}
	helm1 := &helmv2.HelmRelease{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "first-helm-name",
			Namespace: ns.Name,
		},
		Spec: helmv2.HelmReleaseSpec{},
		Status: helmv2.HelmReleaseStatus{
			LastReleaseRevision: 1,
		},
	}

	cm := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "config-map",
			Namespace: ns.Name,
		},
		Data: map[string]string{
			"key": "value",
		},
	}

	cmData, err := json.Marshal(cm)
	g.Expect(err).NotTo(HaveOccurred())

	// Create helm storage.
	storage := types.HelmReleaseStorage{
		Name:     "",
		Manifest: string(cmData),
	}

	storageData, _ := json.Marshal(storage)

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "sh.helm.release.v1.first-helm-name.v1",
			Namespace: ns.Name,
		},
		Data: map[string][]byte{
			"release": []byte(base64.StdEncoding.EncodeToString(storageData)),
		},
	}

	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(ns, helm1, secret, cm).Build()
	cfg := makeServerConfig(client, t, "")
	c := makeServer(cfg, t)

	res, err := c.GetInventory(ctx, &pb.GetInventoryRequest{
		Namespace:    ns.Name,
		ClusterName:  cluster.DefaultCluster,
		Kind:         "HelmRelease",
		Name:         helm1.Name,
		WithChildren: true,
	})

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(res.Entries).To(HaveLen(1))
}

func TestGetInventoryHelmReleaseWithKubeconfig(t *testing.T) {
	g := NewGomegaWithT(t)

	scheme, err := kube.CreateScheme()
	g.Expect(err).NotTo(HaveOccurred())

	ctx := context.Background()

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-namespace",
		},
	}
	helm1 := &helmv2.HelmRelease{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "first-helm-name",
			Namespace: ns.Name,
		},
		Spec: helmv2.HelmReleaseSpec{
			KubeConfig: &meta.KubeConfigReference{
				SecretRef: meta.SecretKeyReference{
					Name: "kubeconfig",
				},
			},
		},
		Status: helmv2.HelmReleaseStatus{
			LastReleaseRevision: 1,
		},
	}

	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(ns, helm1).Build()
	cfg := makeServerConfig(client, t, "")
	c := makeServer(cfg, t)

	res, err := c.GetInventory(ctx, &pb.GetInventoryRequest{
		Namespace:    ns.Name,
		ClusterName:  cluster.DefaultCluster,
		Kind:         "HelmRelease",
		Name:         helm1.Name,
		WithChildren: true,
	})

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(res.Entries).To(HaveLen(0))
}

func TestGetFluxLikeInventory(t *testing.T) {
	g := NewGomegaWithT(t)

	ctx := context.Background()

	ks := &kustomizev1.Kustomization{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-kustomization",
			Namespace: "my-namespace",
		},
		Spec: kustomizev1.KustomizationSpec{
			SourceRef: kustomizev1.CrossNamespaceSourceReference{
				Kind: sourcev1.GitRepositoryKind,
			},
		},
		Status: kustomizev1.KustomizationStatus{
			Inventory: &kustomizev1.ResourceInventory{
				Entries: []kustomizev1.ResourceRef{
					{
						ID:      "my-namespace_my-deployment_apps_Deployment",
						Version: "v1",
					},
				},
			},
		},
	}

	scheme, err := kube.CreateScheme()
	g.Expect(err).To(BeNil())

	k8sClient := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(ks).Build()

	gvk := kustomizev1.GroupVersion.WithKind("Kustomization")
	entries, err := server.GetFluxLikeInventory(ctx, k8sClient, ks.Name, ks.Namespace, gvk)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(entries).To(HaveLen(1))

	expected := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "my-deployment",
				"namespace": "my-namespace",
			},
		},
	}

	g.Expect(entries[0]).To(Equal(expected))
}

func TestParseInventoryFromUnstructured(t *testing.T) {
	// inv lives at status.inventory.entries
	stdErr := errors.New("no status.inventory found on resource, it hasn't been synced yet or is not queryable from this endpoint")

	testCases := []struct {
		name        string
		obj         *unstructured.Unstructured
		expected    []*unstructured.Unstructured
		expectedErr error
	}{
		{
			name:        "no status field",
			obj:         &unstructured.Unstructured{},
			expected:    nil,
			expectedErr: stdErr,
		},
		{
			name: "empty status",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"status": map[string]interface{}{},
				},
			},
			expected:    nil,
			expectedErr: stdErr,
		},
		{
			name: "empty inventory",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"status": map[string]interface{}{
						"inventory": map[string]interface{}{},
					},
				},
			},
			expected: nil,
		},
		{
			name: "empty entry item",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"status": map[string]interface{}{
						"inventory": map[string]interface{}{
							"entries": []interface{}{
								map[string]interface{}{},
							},
						},
					},
				},
			},
			expected:    nil,
			expectedErr: errors.New("unable to parse stored object metadata: "),
		},
		{
			name: "invalid inventory",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"status": map[string]interface{}{
						"inventory": map[string]interface{}{
							"entries": []interface{}{
								map[string]interface{}{
									"v":  "v1",
									"id": "foo",
								},
							},
						},
					},
				},
			},
			expected:    nil,
			expectedErr: errors.New("unable to parse stored object metadata: foo"),
		},
		{
			name: "valid inventory",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"status": map[string]interface{}{
						"inventory": map[string]interface{}{
							"entries": []interface{}{
								map[string]interface{}{
									"v":  "v1",
									"id": "my-namespace_my-deployment_apps_Deployment",
								},
								map[string]interface{}{
									"v":  "v1",
									"id": "my-other-namespace_my-configmap__ConfigMap",
								},
							},
						},
					},
				},
			},
			expected: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "apps/v1",
						"kind":       "Deployment",
						"metadata": map[string]interface{}{
							"name":      "my-deployment",
							"namespace": "my-namespace",
						},
					},
				},
				{
					Object: map[string]interface{}{
						"apiVersion": "v1",
						"kind":       "ConfigMap",
						"metadata": map[string]interface{}{
							"name":      "my-configmap",
							"namespace": "my-other-namespace",
						},
					},
				},
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		// subtests...
		t.Run(tc.name, func(tt *testing.T) {
			gg := NewGomegaWithT(tt)
			// Parse inventory from unstructured
			entries, err := server.ParseInventoryFromUnstructured(tc.obj)

			if err != nil || tc.expectedErr != nil {
				gg.Expect(err).To(MatchError(tc.expectedErr))
			}

			gg.Expect(entries).To(ConsistOf(tc.expected))
		})
	}
}
