package git_info_test

import (
	"github.com/sinlov-go/go-git-tools/git_info"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestRepositoryFistRemote(t *testing.T) {
	t.Logf("~> mock RepositoryFistRemoteInfo")
	// mock RepositoryFistRemoteInfo
	url, err := git_info.RepositoryFistRemoteInfo(projectRootPath, "origin")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("url.Host: %s", url.Host)
	t.Logf("url.User: %s", url.User)
	t.Logf("url.Repo: %s", url.Repo)

	t.Logf("~> do RepositoryFistRemoteInfo")
	// do RepositoryFistRemoteInfo

	// verify RepositoryFistRemoteInfo
	_, err = git_info.RepositoryFistRemoteInfo(projectRootPath, "")
	if err == nil {
		t.Fatal("RepositoryFistRemoteInfo with empty remote should not be nil")
	}
	assert.Equal(t, "RepositoryFistRemoteInfo remote is empty", err.Error())
}

func TestRepositoryConfigPath(t *testing.T) {
	t.Logf("~> mock RepositoryConfigPath")
	// mock RepositoryConfigPath
	cfg, err := git_info.RepositoryConfigPath(projectRootPath)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("config: %v", cfg.Remotes["origin"].URLs[0])

	t.Logf("~> do RepositoryConfigPath")
	// do RepositoryConfigPath

	// verify RepositoryConfigPath
}

func TestIsPathUnderGitManagement(t *testing.T) {
	t.Logf("~> mock IsPathUnderGitManagement")
	// mock IsPathUnderGitManagement
	t.Logf("projectRoot: %s", projectRootPath)

	t.Logf("~> do IsPathUnderGitManagement")
	// do IsPathUnderGitManagement
	assert.True(t, git_info.IsPathUnderGitManagement(projectRootPath))

	// verify IsPathUnderGitManagement

	topDir := filepath.Dir(projectRootPath)
	assert.False(t, git_info.IsPathUnderGitManagement(topDir))
}

func TestIsPathGitManagementRoot(t *testing.T) {
	t.Logf("~> mock IsPathGitManagementRoot")
	// mock IsPathGitManagementRoot

	t.Logf("~> do IsPathGitManagementRoot")
	// do IsPathGitManagementRoot
	isRoot, err := git_info.IsPathGitManagementRoot(projectRootPath)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, isRoot)

	dirRoot, err := git_info.IsPathGitManagementRoot(filepath.Dir(projectRootPath))
	if err == nil {
		t.Fatal("should not be nil")
	}
	t.Logf("dirRoot IsPathGitManagementRoot err: %v", err)
	assert.False(t, dirRoot)
	docRoot, err := git_info.IsPathGitManagementRoot(filepath.Join(projectRootPath, "doc"))
	if err == nil {
		t.Fatal("should not be nil")
	}
	t.Logf("docRoot IsPathGitManagementRoot err: %v", err)
	assert.False(t, docRoot)

	// verify IsPathGitManagementRoot
}

func TestRepositoryHeadByPath(t *testing.T) {
	t.Logf("~> mock RepositoryHeadByPath")
	// mock RepositoryHeadByPath
	headByPath, err := git_info.RepositoryHeadByPath(projectRootPath)

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("headByPath: %v", headByPath)

	_, err = git_info.RepositoryHeadByPath(filepath.Dir(projectRootPath))
	if err == nil {
		t.Fatal("should not be nil")
	}
}

func TestRepositoryNowBranchByPath(t *testing.T) {
	t.Logf("~> mock RepositoryNowBranchByPath")
	// mock RepositoryNowBranchByPath
	branchByPath, err := git_info.RepositoryNowBranchByPath(projectRootPath)
	if err != nil {
		t.Logf("RepositoryNowBranchByPath err: %v", err)
	}
	t.Logf("branchByPath: %s", branchByPath)
	_, err = git_info.RepositoryNowBranchByPath(filepath.Dir(projectRootPath))
	if err == nil {
		t.Logf("RepositoryNowBranchByPath err: %v", err)
	}
}
