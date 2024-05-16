package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RepositorySuite struct {
	suite.Suite

	r *git.Repository

	repo *repo
}

func (s *RepositorySuite) SetupTest() {
	var err error
	s.r, err = git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/sinlov-go/go-git-tools",
	})
	s.NoError(err)

	s.repo = &repo{
		gitRepo: s.r,
	}
}

func (s *RepositorySuite) TestLogSuccess() {
	res, gotErr := s.repo.Log("", "")
	s.NoError(gotErr)
	fmt.Printf("res len %d\n", len(res))
	if len(res) > 0 {
		hash := res[0].Hash
		if !hash.IsZero() {
			fmt.Printf("res[0] hash %s\n", hash.String())
		}
	}
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func TestHeadBranchName(t *testing.T) {
	const envFlag = "ENV_TEST_LOCAL_BRANCH_NAME"
	valEnvPath := env_kit.FetchOsEnvStr(envFlag, "")
	if valEnvPath == "" {
		t.Skipf("env is empty: %s", envFlag)
	}

	repository, err := NewRepositoryByPath(valEnvPath)
	if err != nil {
		t.Fatal(err)
	}
	branchName, err := repository.HeadBranchName()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("branchName %s\n", branchName)
}

func TestCheckLocalBranchIsDirty(t *testing.T) {
	const envFlag = "ENV_TEST_CHECK_LOCAL_BRANCH_IS_DIRTY"
	valEnvPath := env_kit.FetchOsEnvStr(envFlag, "")
	if valEnvPath == "" {
		t.Skipf("env is empty: %s", envFlag)
	}
	repository, err := NewRepositoryByPath(valEnvPath)
	if err != nil {
		t.Fatal(err)
	}

	isDirty, err := repository.CheckLocalBranchIsDirty()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("isDirty %v\n", isDirty)
}

func TestCheckHasSubmodules(t *testing.T) {
	const envFlag = "ENV_TEST_CHECK_HAS_SUBMODULES"
	valEnvPath := env_kit.FetchOsEnvStr(envFlag, "")
	if valEnvPath == "" {
		t.Skipf("env is empty: %s", envFlag)
	}
	repository, err := NewRepositoryByPath(valEnvPath)
	if err != nil {
		t.Fatal(err)
	}

	submodules, err := repository.CheckHasSubmodules()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("CheckHasSubmodules: %v\n", submodules)
}

func TestCheckSubmodulesIsDirty(t *testing.T) {
	const envFlag = "ENV_TEST_CHECK_HAS_SUBMODULES"
	valEnvPath := env_kit.FetchOsEnvStr(envFlag, "")
	if valEnvPath == "" {
		t.Skipf("env is empty: %s", envFlag)
	}
	repository, err := NewRepositoryByPath(valEnvPath)
	if err != nil {
		t.Fatal(err)
	}

	dirty, err := repository.CheckSubmodulesIsDirty()
	if err != nil {
		t.Logf("CheckSubmodulesIsDirty isDirty %v err: %v\n", dirty, err)
	} else {
		t.Logf("CheckSubmodulesIsDirty: %v\n", dirty)
	}
}

func TestFetchTags(t *testing.T) {
	const envFlag = "ENV_TEST_FETCH_TAGS"
	valEnvPath := env_kit.FetchOsEnvStr(envFlag, "")
	if valEnvPath == "" {
		t.Skipf("env is empty: %s", envFlag)
	}
	repository, err := NewRepositoryByPath(valEnvPath)
	if err != nil {
		t.Fatal(err)
	}

	errFetchTags := repository.FetchTags()
	if errFetchTags != nil {
		t.Fatal(errFetchTags)
	}
}

func TestPullOriginWithPublicRepo(t *testing.T) {
	const envFlag = "ENV_TEST_LOCAL_PULL_ORIGIN_WITH_PUBLIC_REPO"
	valEnvPath := env_kit.FetchOsEnvStr(envFlag, "")
	if valEnvPath == "" {
		t.Skipf("env is empty: %s", envFlag)
	}
	repository, err := NewRepositoryByPath(valEnvPath)
	if err != nil {
		t.Fatal(err)
	}

	errPullOrigin := repository.PullOrigin()
	if errPullOrigin != nil {
		t.Fatal(errPullOrigin)
	}

}

func TestSetAuth(t *testing.T) {
	const envFlag = "ENV_TEST_REPO_SET_AUTH"
	valEnvPath := env_kit.FetchOsEnvStr(envFlag, "")
	if valEnvPath == "" {
		t.Skipf("env is empty: %s", envFlag)
	}

	const envFlagKeyPath = "ENV_TEST_REPO_SET_AUTH_KEY_PATH"
	valSshKeyPath := env_kit.FetchOsEnvStr(envFlagKeyPath, "")
	if valSshKeyPath == "" {
		t.Skipf("env is empty: %s", envFlagKeyPath)
	}

	const envFlagKeyPassword = "ENV_TEST_REPO_SET_AUTH_KEY_PASS_WORD"
	valSshKeyPassWord := env_kit.FetchOsEnvStr(envFlagKeyPassword, "")

	repository, err := NewRepositoryByPath(valEnvPath)
	if err != nil {
		t.Fatal(err)
	}

	auth, err := ssh.NewPublicKeysFromFile("git", valSshKeyPath, valSshKeyPassWord)
	if err != nil {
		t.Fatal(err)
	}

	repository.SetAuthMethod(auth)

	err = repository.PullOrigin()
	if err != nil {
		t.Fatal(err)
	}
}
