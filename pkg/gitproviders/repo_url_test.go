package gitproviders

import (
	"net/url"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

var _ = DescribeTable("detectGitProviderFromURL", func(input string, expected GitProviderName) {
	result, err := detectGitProviderFromURL(input, map[string]string{
		"bitbucket.weave.works": "bitbucket-server",
	})
	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(Equal(expected))
},
	Entry("ssh+github", "ssh://git@github.com/weaveworks/weave-gitops.git", GitProviderGitHub),
	Entry("ssh+gitlab", "ssh://git@gitlab.com/weaveworks/weave-gitops.git", GitProviderGitLab),
	Entry("https+bitbucket", "https://bitbucket.weave.works/scm/wg/config.git", GitProviderBitBucketServer),
)

var _ = Describe("get owner from url", func() {
	DescribeTable("getOwnerFromURL", func(normalizedURL string, providerName GitProviderName, expected string) {
		u, err := url.Parse(normalizedURL)
		Expect(err).NotTo(HaveOccurred())
		result, err := getOwnerFromURL(*u, providerName)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(expected))
	},
		Entry("github", "ssh://git@github.com/weaveworks/weave-gitops.git", GitProviderGitHub, "weaveworks"),
		Entry("gitlab", "ssh://git@gitlab.com/weaveworks/weave-gitops.git", GitProviderGitLab, "weaveworks"),
		Entry("gitlab", "ssh://git@gitlab.com/weaveworks/infra/weave-gitops.git", GitProviderGitLab, "weaveworks/infra"),
		Entry("gitlab", "ssh://git@gitlab.com/weaveworks/infra/dev/weave-gitops.git", GitProviderGitLab, "weaveworks/infra/dev"),
	)

	It("missing owner", func() {
		normalizedURL := "ssh://git@gitlab.com/weave-gitops.git"
		u, err := url.Parse(normalizedURL)
		Expect(err).NotTo(HaveOccurred())
		_, err = getOwnerFromURL(*u, GitProviderGitLab)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("could not get owner from url ssh://git@gitlab.com/weave-gitops.git"))
	})

	It("empty url", func() {
		normalizedURL := ""
		u, err := url.Parse(normalizedURL)
		Expect(err).NotTo(HaveOccurred())
		_, err = getOwnerFromURL(*u, GitProviderGitLab)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("could not get owner from url "))
	})
})

type expectedRepoURL struct {
	s        string
	owner    string
	name     string
	provider GitProviderName
	protocol RepositoryURLProtocol
}

var _ = DescribeTable("NewRepoURL", func(input, gitProviderEnv string, expected expectedRepoURL) {
	if gitProviderEnv != "" {
		viper.Set("git-host-types", gitProviderEnv)
	}
	result, err := NewRepoURL(input)
	Expect(err).NotTo(HaveOccurred())

	Expect(result.String()).To(Equal(expected.s))
	u, err := url.Parse(expected.s)
	Expect(err).NotTo(HaveOccurred())
	Expect(result.URL()).To(Equal(u))
	Expect(result.Owner()).To(Equal(expected.owner))
	Expect(result.Provider()).To(Equal(expected.provider))
	Expect(result.Protocol()).To(Equal(expected.protocol))
},
	Entry("github git clone style", "git@github.com:someuser/podinfo.git", "", expectedRepoURL{
		s:        "ssh://git@github.com/someuser/podinfo.git",
		owner:    "someuser",
		name:     "podinfo",
		provider: GitProviderGitHub,
		protocol: RepositoryURLProtocolSSH,
	}),
	Entry("github url style", "ssh://git@github.com/someuser/podinfo.git", "", expectedRepoURL{
		s:        "ssh://git@github.com/someuser/podinfo.git",
		owner:    "someuser",
		name:     "podinfo",
		provider: GitProviderGitHub,
		protocol: RepositoryURLProtocolSSH,
	}),
	Entry("github https", "https://github.com/someuser/podinfo.git", "", expectedRepoURL{
		s:        "ssh://git@github.com/someuser/podinfo.git",
		owner:    "someuser",
		name:     "podinfo",
		provider: GitProviderGitHub,
		protocol: RepositoryURLProtocolSSH,
	}),
	Entry("gitlab git clone style", "git@gitlab.com:someuser/podinfo.git", "", expectedRepoURL{
		s:        "ssh://git@gitlab.com/someuser/podinfo.git",
		owner:    "someuser",
		name:     "podinfo",
		provider: GitProviderGitLab,
		protocol: RepositoryURLProtocolSSH,
	}),
	Entry("gitlab https", "https://gitlab.com/someuser/podinfo.git", "", expectedRepoURL{
		s:        "ssh://git@gitlab.com/someuser/podinfo.git",
		owner:    "someuser",
		name:     "podinfo",
		provider: GitProviderGitLab,
		protocol: RepositoryURLProtocolSSH,
	}),
	Entry("trailing slash in url", "https://github.com/sympatheticmoose/podinfo-deploy/", "", expectedRepoURL{
		s:        "ssh://git@github.com/sympatheticmoose/podinfo-deploy.git",
		owner:    "sympatheticmoose",
		name:     "podinfo-deploy",
		provider: GitProviderGitHub,
		protocol: RepositoryURLProtocolSSH,
	}),
	Entry("subsubgroup", "https://github.com/sympatheticmoose/infra/dev/podinfo-deploy/", "", expectedRepoURL{
		s:        "ssh://git@github.com/sympatheticmoose/infra/dev/podinfo-deploy.git",
		owner:    "sympatheticmoose/infra/dev",
		name:     "podinfo-deploy",
		provider: GitProviderGitHub,
		protocol: RepositoryURLProtocolSSH,
	}),
	Entry(
		"custom domain",
		"git@gitlab.acme.org/sympatheticmoose/podinfo-deploy/",
		"gitlab.acme.org=gitlab",
		expectedRepoURL{
			s:        "ssh://git@gitlab.acme.org/sympatheticmoose/podinfo-deploy.git",
			owner:    "sympatheticmoose",
			name:     "podinfo-deploy",
			provider: "gitlab",
			protocol: RepositoryURLProtocolSSH,
		}),
	Entry(
		"bitbucket custom domain with port",
		"https://stash.stashtestserver.link:7990/scm/~someuser/podinfo-deploy.git",
		"stash.stashtestserver.link:7990=bitbucket-server",
		expectedRepoURL{
			s:        "ssh://git@stash.stashtestserver.link:7990/scm/~someuser/podinfo-deploy.git",
			owner:    "scm/~someuser",
			name:     "podinfo-deploy",
			provider: "bitbucket-server",
			protocol: RepositoryURLProtocolSSH,
		}),
	Entry(
		"bitbucket custom domain with port on ssh",
		"ssh://git@stash.stashtestserver.link:7990/scm/~someuser/podinfo-deploy.git",
		"stash.stashtestserver.link:7990=bitbucket-server",
		expectedRepoURL{
			s:        "ssh://git@stash.stashtestserver.link:7990/scm/~someuser/podinfo-deploy.git",
			owner:    "scm/~someuser",
			name:     "podinfo-deploy",
			provider: "bitbucket-server",
			protocol: RepositoryURLProtocolSSH,
		}),
)
