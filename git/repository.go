package git

import (
	"fmt"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/storage"
	"github.com/sinlov-go/go-git-tools/git_info"

	goGit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

const (
	head = "HEAD"

	OriginDefault = "origin"

	defaultCommitCount = 10
)

type repo struct {
	gitRepo    *goGit.Repository
	remoteName string

	transportAuthMethod   *transport.AuthMethod
	transportProxyOptions *transport.ProxyOptions
}

// Repository is an abstraction for git-repository
type Repository interface {
	HeadReference() (*plumbing.Reference, error)

	HeadBranchName() (string, error)

	CheckHasSubmodules() (bool, error)
	CheckSubmodulesIsDirty() (bool, error)
	CheckLocalBranchIsDirty() (bool, error)

	Log(fromRev, toRev string) ([]Commit, error)

	CommitLatestTagByTime() (*Commit, error)
	CommitTagSearchByName(tagName string) (*Commit, error)
	CommitTagSearchByFirstLine(firstLine string) (*Commit, error)
	Commit(commitMessage string, paths ...string) error

	TagLatestByCommitTime() (*object.Tag, error)

	RemoteInfo(remoteName string, configUrlIndex int) (*git_info.GitRemoteInfo, error)

	SetProxyOptions(options transport.ProxyOptions)
	SetAuthMethod(auth transport.AuthMethod)

	PullOrigin() error
	FetchTags() error
}

// NewRepositoryByPath return Repository from path
func NewRepositoryByPath(path string) (Repository, error) {
	return NewRepositoryRemoteByPath(OriginDefault, path)
}

// NewRepositoryRemoteByPath
// remote most is git.OriginDefault
// return Repository from path
func NewRepositoryRemoteByPath(remote string, path string) (Repository, error) {
	gitRepo, err := goGit.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	return &repo{
		gitRepo:    gitRepo,
		remoteName: remote,
	}, nil
}

// NewRepositoryClone
// return Repository from git.Repository
func NewRepositoryClone(s storage.Storer, worktree billy.Filesystem, o *goGit.CloneOptions) (Repository, error) {
	return NewRepositoryRemoteClone(OriginDefault, s, worktree, o)
}

// NewRepositoryRemoteClone
// remote most is git.OriginDefault
// return Repository from git.Repository
func NewRepositoryRemoteClone(remote string, s storage.Storer, worktree billy.Filesystem, o *goGit.CloneOptions) (Repository, error) {
	repository, err := goGit.Clone(s, worktree, o)
	if err != nil {
		return nil, err
	}
	if repository == nil {
		return nil, fmt.Errorf("repository is nil")
	}

	return &repo{
		gitRepo:    repository,
		remoteName: remote,
	}, nil
}
