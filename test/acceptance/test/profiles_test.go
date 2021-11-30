package acceptance

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	sourcev1beta1 "github.com/fluxcd/source-controller/api/v1beta1"
	pb "github.com/weaveworks/weave-gitops/pkg/api/profiles"
	"github.com/weaveworks/weave-gitops/pkg/gitproviders"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"path/filepath"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Weave GitOps Profiles API", func() {
	var (
		namespace        = "test-namespace"
		appRepoRemoteURL string
		tip              TestInputs
		profilesSvc      = "profiles-server"
		clientSet        *kubernetes.Clientset
		kClient          client.Client
	)

	BeforeEach(func() {
		Expect(FileExists(WEGO_BIN_PATH)).To(BeTrue())
		Expect(GITHUB_ORG).NotTo(BeEmpty())

		_, _, err := ResetOrCreateCluster(namespace, true)
		Expect(err).NotTo(HaveOccurred())

		private := true
		tip = generateTestInputs()
		_ = initAndCreateEmptyRepo(tip.appRepoName, gitproviders.GitProviderGitHub, private, GITHUB_ORG)

		clientSet, kClient = buildKubernetesClients()
	})

	AfterEach(func() {
		deleteRepo(tip.appRepoName, gitproviders.GitProviderGitHub, GITHUB_ORG)
	})

	It("Works", func() {
		By("Installing the Profiles API and setting up the profile helm repository")
		appRepoRemoteURL = "git@github.com:" + GITHUB_ORG + "/" + tip.appRepoName + ".git"
		installAndVerifyWego(namespace, appRepoRemoteURL)
		deployProfilesHelmRepository(kClient, namespace)

		By("Getting a list of profiles")
		resp, statusCode, err := kubernetesDoRequest(namespace, profilesSvc, "8000", "/v1/profiles", clientSet)
		Expect(err).NotTo(HaveOccurred())
		Expect(statusCode).To(Equal(http.StatusOK))

		profiles := pb.GetProfilesResponse{}
		Expect(json.Unmarshal(resp, &profiles)).To(Succeed())
		Expect(profiles.Profiles).NotTo(BeNil())
		Expect(profiles.Profiles).To(ConsistOf(&pb.Profile{
			Name:        "podinfo",
			Home:        "https://github.com/stefanprodan/podinfo",
			Description: "Podinfo Helm chart for Kubernetes",
			Sources:     []string{"https://github.com/stefanprodan/podinfo"},
			Maintainers: []*pb.Maintainer{
				{
					Name:  "stefanprodan",
					Email: "stefanprodan@users.noreply.github.com",
				},
			},
			//These have to not be nil for the assertion to work. because proto :shrug:
			Keywords:    []string{},
			Annotations: map[string]string{},
		}))

		By("Getting the values for a profile")
		resp, statusCode, err = kubernetesDoRequest(namespace, profilesSvc, "8000", "/v1/profiles/podinfo/6.0.1/values", clientSet)
		Expect(err).NotTo(HaveOccurred())
		Expect(statusCode).To(Equal(http.StatusOK))

		profileValues := pb.GetProfileValuesResponse{}
		Expect(json.Unmarshal(resp, &profileValues)).To(Succeed())
		values, err := base64.StdEncoding.DecodeString(profileValues.Values)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(values)).To(ContainSubstring("# Default values for podinfo"))
	})
})

func buildKubernetesClients() (*kubernetes.Clientset, client.Client) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	Expect(err).NotTo(HaveOccurred())
	clientSet, err := kubernetes.NewForConfig(config)
	Expect(err).NotTo(HaveOccurred())

	scheme := runtime.NewScheme()
	schemeBuilder := runtime.SchemeBuilder{
		sourcev1beta1.AddToScheme,
	}
	Expect(schemeBuilder.AddToScheme(scheme)).To(Succeed())

	kClient, err := client.New(config, client.Options{
		Scheme: scheme,
	})

	Expect(err).NotTo(HaveOccurred())

	return clientSet, kClient
}

func deployProfilesHelmRepository(rawClient client.Client, namespace string) {
	weaveworksChartsHelmRepository := &sourcev1beta1.HelmRepository{
		TypeMeta: metav1.TypeMeta{
			Kind:       sourcev1beta1.HelmRepositoryKind,
			APIVersion: sourcev1beta1.GroupVersion.Identifier(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "weaveworks-charts",
			Namespace: namespace,
		},
		Spec: sourcev1beta1.HelmRepositorySpec{
			URL:      "https://weaveworks.github.io/profiles-examples",
			Interval: metav1.Duration{Duration: time.Minute * 10},
		},
	}
	err = rawClient.Create(context.TODO(), weaveworksChartsHelmRepository)
	Expect(err).NotTo(HaveOccurred())
	Eventually(func() error {
		helmRepo := sourcev1beta1.HelmRepository{}
		err := rawClient.Get(context.TODO(), client.ObjectKeyFromObject(weaveworksChartsHelmRepository), &helmRepo)
		if err != nil {
			return err
		}

		if len(helmRepo.Status.Conditions) == 1 && helmRepo.Status.Conditions[0].Type == "Ready" && helmRepo.Status.Conditions[0].Status == "True" {
			return nil
		}
		return fmt.Errorf("HelmRepository not ready %v", helmRepo.Status)
	}, "10s", "1s").Should(Succeed())
}

func kubernetesDoRequest(namespace, serviceName, servicePort, path string, clientset *kubernetes.Clientset) ([]byte, int, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, 0, err
	}

	responseWrapper := clientset.CoreV1().Services(namespace).ProxyGet("http", serviceName, servicePort, u.String(), nil)

	data, err := responseWrapper.DoRaw(context.TODO())
	if err != nil {
		if se, ok := err.(*errors.StatusError); ok {
			return nil, int(se.Status().Code), nil
		}

		return nil, 0, err
	}

	return data, http.StatusOK, nil
}
