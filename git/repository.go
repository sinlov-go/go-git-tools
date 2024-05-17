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
	// HeadReference
	// get head reference
	HeadReference() (*plumbing.Reference, error)

	// HeadBranchName
	// get head branch name
	HeadBranchName() (string, error)

	// CheckHasSubmodules
	// check has submodules
	// no has submodules must return false, nil
	CheckHasSubmodules() (bool, error)
	// CheckSubmodulesIsDirty
	// check submodules is dirty.
	// Warning submodule version must same as .gitmodules record
	// like run as: git submodule status --recursive
	// Fix dirty use: git submodule update --init --recursive
	CheckSubmodulesIsDirty() (bool, error)
	// CheckLocalBranchIsDirty
	// find dirty file will return true and error is nil.
	// Like run cmd as: git status --porcelain
	CheckLocalBranchIsDirty() (bool, error)

	// CheckWorkTreeIsDirtyWithGitCmd
	// check work tree is dirty by run exec: git status --porcelain
	CheckWorkTreeIsDirtyWithGitCmd() (bool, error)

	// Log return all commits between <from revision> and <to revision>
	Log(fromRev, toRev string) ([]Commit, error)

	// CommitLatestTagByTime
	// get commit by tag latest by commit time
	CommitLatestTagByTime() (*Commit, error)
	// CommitTagSearchByName
	// get commit by tag search by name
	CommitTagSearchByName(tagName string) (*Commit, error)
	// CommitTagSearchByFirstLine
	// get commit by tag search by first line
	CommitTagSearchByFirstLine(firstLine string) (*Commit, error)
	// TagLatestByCommitTime
	//
	//	latest tag find by commit time, please ensure that the time of the device submitting the tag is synchronized correctly.
	//	check at: git show-ref --tag
	//
	// return latest tag
	TagLatestByCommitTime() (*object.Tag, error)

	Commit(commitMessage string, paths ...string) error // commit with message

	// RemoteInfo
	//
	// remote most is git.OriginDefault
	//
	// configUrlIndex most is 0
	//
	// return Repository from git.Repository
	RemoteInfo(remoteName string, configUrlIndex int) (*git_info.GitRemoteInfo, error)

	// SetAuthMethod
	// auth transport.AuthMethod
	//
	//	auth, err := ssh.NewPublicKeysFromFile("git", valSshKeyPath, valSshKeyPassWord)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	repository.SetAuthMethod(auth)
	SetAuthMethod(auth transport.AuthMethod)
	// SetProxyOptions
	// options transport.ProxyOptions
	SetProxyOptions(options transport.ProxyOptions)

	// PullOrigin
	// warning: direct support public repository, private repository must use SetAuth before.
	// like: git pull origin
	PullOrigin() error
	// FetchTags
	// warning: direct support public repository, private repository must use SetAuth before.
	// fetch tags
	// like run as: git fetch --tags
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
