package git_info

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	gitConfig "github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	gitUrls "github.com/whilp/git-urls"
	"net/url"
	"strings"
)

// IsPathUnderGitManagement
//
//	@Description: IsPathUnderGitManagement check path is under git management
func IsPathUnderGitManagement(path string) bool {
	_, err := git.PlainOpen(path)
	return err == nil
}

func IsPathGitManagementRoot(path string) (bool, error) {
	repository, err := git.PlainOpen(path)
	if err != nil {
		return false, fmt.Errorf("IsPathGitManagementRoot can not open repository at path %s , err: %s", path, err)
	}
	worktree, err := repository.Worktree()
	if err != nil {
		return false, fmt.Errorf("IsPathGitManagementRoot can not get worktree: %s", err)
	}
	root := worktree.Filesystem.Root()
	if root != path {
		return false, fmt.Errorf("IsPathGitManagementRoot path: %s is not root, root path is %s", path, root)
	}
	return true, nil
}

// RepositoryConfigPath
//
//	cfg, err := git_tools.RepositoryConfigPath(projectRootPath)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Logf("config: %v", cfg.Remotes["origin"].URLs[0])
//
// more use See https://github.com/go-git/go-git/blob/v5.7.0/config/config_test.go#L18
func RepositoryConfigPath(path string) (*gitConfig.Config, error) {
	repository, err := git.PlainOpen(path)
	if err != nil {
		return nil, fmt.Errorf("RepositoryConfigPath can not open repository at path %s , err: %s", path, err)
	}

	return repository.Config()

}

type GitRemoteInfo struct {
	UrlStr   string
	Scheme   string
	Host     string
	Hostname string
	Port     string
	UserInfo *url.Userinfo
	User     string
	Repo     string
}

// RepositoryFistRemoteInfo
//
//	url, err := git_tools.RepositoryFistRemoteInfo(projectRootPath, "origin")
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Logf("url.Host: %s", url.Host)
//	t.Logf("url.User: %s", url.User)
//	t.Logf("url.Repo: %s", url.Repo)
//
// See: https://mirrors.edge.kernel.org/pub/software/scm/git/docs/git-clone.html
func RepositoryFistRemoteInfo(path string, remote string) (*GitRemoteInfo, error) {
	if remote == "" {
		return nil, fmt.Errorf("RepositoryFistRemoteInfo remote is empty")
	}
	cfg, err := RepositoryConfigPath(path)
	if err != nil {
		return nil, fmt.Errorf("RepositoryFistRemoteInfo can not open repository at path %s , err: %s", path, err)
	}
	remoteConfig := cfg.Remotes[remote]
	if remoteConfig == nil {
		return nil, fmt.Errorf("RepositoryFistRemoteInfo remote: %s not found", remote)
	}
	if len(remoteConfig.URLs) == 0 {
		return nil, fmt.Errorf("RepositoryFistRemoteInfo remote: %s URLs is empty", remote)
	}
	urlStr := remoteConfig.URLs[0]
	parse, err := gitUrls.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("RepositoryFistRemoteInfo remote: %s URLs[0]: %s parse err: %s", remote, urlStr, err)
	}
	if parse.Path == "" {
		return nil, fmt.Errorf("RepositoryFistRemoteInfo remote: %s URLs[0]: %s parse path is empty", remote, urlStr)
	}
	removeGitPath := strings.Replace(parse.Path, ".git", "", -1)
	pathSplit := strings.Split(removeGitPath, "/")

	if strings.Contains(parse.Scheme, "http") {
		if len(pathSplit) < 3 {
			return nil, fmt.Errorf("RepositoryFistRemoteInfo remote: %s URLs[0]: %s parse path not support", remote, parse.Path)
		}
		return &GitRemoteInfo{
			UrlStr:   urlStr,
			Scheme:   parse.Scheme,
			Host:     parse.Host,
			Hostname: parse.Hostname(),
			Port:     parse.Port(),
			UserInfo: parse.User,
			User:     pathSplit[1],
			Repo:     pathSplit[2],
		}, nil
	}
	if len(pathSplit) < 2 {
		return nil, fmt.Errorf("RepositoryFistRemoteInfo remote: %s URLs[0]: %s parse path not support", remote, parse.Path)
	}
	return &GitRemoteInfo{
		UrlStr:   urlStr,
		Scheme:   parse.Scheme,
		Host:     parse.Host,
		Hostname: parse.Hostname(),
		Port:     parse.Port(),
		UserInfo: parse.User,
		User:     pathSplit[0],
		Repo:     pathSplit[1],
	}, nil
}

func RepositoryHeadByPath(path string) (*plumbing.Reference, error) {
	repository, err := git.PlainOpen(path)
	if err != nil {
		return nil, fmt.Errorf("RepositoryHeadByPath can not open repository at path %s , err: %s", path, err)
	}
	head, err := repository.Head()
	if err != nil {
		return nil, fmt.Errorf("RepositoryHeadByPath can not get head: %s", err)
	}
	return head, nil
}

func RepositoryNowBranchByPath(path string) (string, error) {
	repository, err := git.PlainOpen(path)
	if err != nil {
		return "", fmt.Errorf("RepositoryNowBranchByPath can not open repository at path %s , err: %s", path, err)
	}
	head, err := repository.Head()
	if err != nil {
		return "", fmt.Errorf("RepositoryNowBranchByPath can not get head: %s", err)
	}
	referenceName := head.Name()
	if !referenceName.IsBranch() {
		return "", fmt.Errorf("RepositoryNowBranchByPath head is not branch: %s", referenceName.String())
	}
	return referenceName.Short(), nil
}
