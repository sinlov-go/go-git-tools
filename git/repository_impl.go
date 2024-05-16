package git

import (
	"errors"
	"fmt"
	"github.com/chainguard-dev/git-urls"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/sinlov-go/go-git-tools/git_info"
	"strings"
	"time"
)

// HeadReference
// get head reference
func (r *repo) HeadReference() (*plumbing.Reference, error) {
	return r.gitRepo.Head()
}

// HeadBranchName
// get head branch name
func (r *repo) HeadBranchName() (string, error) {
	referenceHead, err := r.HeadReference()
	if err != nil {
		return "", fmt.Errorf("HeadBranchName can not get head: %s", err)
	}
	referenceName := referenceHead.Name()
	if !referenceName.IsBranch() {
		return "", fmt.Errorf("HeadBranchName head is not branch: %s", referenceName.String())
	}
	return referenceName.Short(), nil
}

// CheckHasSubmodules
// check has submodules
// no has submodules must return false, nil
func (r *repo) CheckHasSubmodules() (bool, error) {
	if r.gitRepo == nil {
		return true, fmt.Errorf("check has submodules error, git repo not init, or can not read")
	}
	worktree, err := r.gitRepo.Worktree()
	if err != nil {
		return true, fmt.Errorf("check has submodules error, want get work tree err: %v", err)
	}

	submodules, err := worktree.Submodules()
	if err != nil {
		return true, fmt.Errorf("check has submodules error, want get submodules err: %v", err)
	}

	return len(submodules) > 0, nil
}

// CheckSubmodulesIsDirty
// check submodules is dirty.
// Warning submodule version must same as .gitmodules record
// like run as: git submodule status --recursive
// Fix dirty use: git submodule update --init --recursive
func (r *repo) CheckSubmodulesIsDirty() (bool, error) {
	if r.gitRepo == nil {
		return true, fmt.Errorf("check submodules is dirty error, git repo not init, or can not read")
	}
	worktree, err := r.gitRepo.Worktree()
	if err != nil {
		return true, fmt.Errorf("check submodules is dirty error, want get work tree err: %v", err)
	}

	submodules, err := worktree.Submodules()
	if err != nil {
		return true, fmt.Errorf("check submodules is dirty error, want get submodules err: %v", err)
	}
	if len(submodules) == 0 {
		return true, fmt.Errorf("check submodules is dirty error, now repo not has no submodules")
	}
	var isDirty bool
	var errDirty error
	for _, submodule := range submodules {
		status, errSubModule := submodule.Status()
		if errSubModule != nil {
			return true, fmt.Errorf("check submodules is dirty error, want get submodule status err: %v", errSubModule)
		}
		if !status.IsClean() {
			isDirty = true
			errDirty = fmt.Errorf("check submodules is dirty, URL: %s status: %v", submodule.Config().URL, status.String())
			break
		}
	}
	return isDirty, errDirty
}

// CheckLocalBranchIsDirty
// find dirty file will return true and error is nil.
// Like run cmd as: git status --porcelain
func (r *repo) CheckLocalBranchIsDirty() (bool, error) {
	if r.gitRepo == nil {
		return true, fmt.Errorf("check local branch is dirty error, git repo not init, or can not read")
	}
	worktree, err := r.gitRepo.Worktree()
	if err != nil {
		return true, fmt.Errorf("check local branch is dirty error, want get work tree err: %v", err)
	}

	status, errStatus := worktree.Status()
	if errStatus != nil {
		return true, fmt.Errorf("want get work tree status err: %v", errStatus)
	}
	if status.IsClean() {
		return false, nil
	}
	return true, nil
}

// Log return all commits between <from revision> and <to revision>
func (r *repo) Log(fromRev, toRev string) ([]Commit, error) {
	if fromRev == "" {
		fromRev = head
	}

	fromHash, err := r.gitRepo.ResolveRevision(plumbing.Revision(fromRev))
	if err != nil {
		return nil, fmt.Errorf("failed to resolve %s: %w", fromRev, err)
	}

	if toRev == "" {
		return r.logWithStopFn(fromHash, nil, nil)
	}

	toHash, err := r.gitRepo.ResolveRevision(plumbing.Revision(toRev))
	if err != nil {
		return nil, fmt.Errorf("failed to resolve %s: %w", toRev, err)
	}

	return r.logWithStopFn(fromHash, nil, stopAtHash(toHash))
}

func (r *repo) CommitLatestTagByTime() (*Commit, error) {
	tagLatest, err := r.TagLatestByCommitTime()
	if err != nil {
		return nil, err
	}
	commitTag, err := tagLatest.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to get latest commit by tag: %w", err)
	}
	commit := newCommit(commitTag)

	return &commit, nil
}

func (r *repo) CommitTagSearchByName(tagName string) (*Commit, error) {
	if tagName == "" {
		return nil, fmt.Errorf("commit tag search by name is empty")
	}
	tags, err := r.gitRepo.Tags()
	if err != nil {
		return nil, fmt.Errorf("failed to get tags: %w", err)
	}
	var wantCommit *Commit
	errForEach := tags.ForEach(func(ref *plumbing.Reference) error {
		referNameShort := ref.Name().Short()
		if referNameShort == tagName {
			revision, errRR := r.gitRepo.ResolveRevision(plumbing.Revision(ref.Name()))
			if errRR != nil {
				return errRR
			}
			cIter, errLog := r.gitRepo.Log(&git.LogOptions{
				From:  *revision,
				Order: git.LogOrderCommitterTime,
			})
			if errLog != nil {
				return errLog
			}
			errRR = cIter.ForEach(func(c *object.Commit) error {
				if c.Hash.String() == revision.String() {
					cmt := newCommit(c)
					wantCommit = &cmt
				}
				return nil
			})
			if errRR != nil {
				return errRR
			}
			return nil
		}
		return nil
	})
	if errForEach != nil {
		return nil, fmt.Errorf("failed to ForEach tags: %w", errForEach)
	}
	if wantCommit == nil {
		return nil, fmt.Errorf("can not find commit at tag: %s", tagName)
	}
	return wantCommit, nil
}

func (r *repo) CommitTagSearchByFirstLine(firstLine string) (*Commit, error) {
	if firstLine == "" {
		return nil, fmt.Errorf("commit tag search by firstLine is empty")
	}
	tags, err := r.gitRepo.Tags()
	if err != nil {
		return nil, fmt.Errorf("failed to get tags: %w", err)
	}
	var wantCommit *Commit
	errForEach := tags.ForEach(func(ref *plumbing.Reference) error {
		revision, errRR := r.gitRepo.ResolveRevision(plumbing.Revision(ref.Name()))
		if errRR != nil {
			return errRR
		}
		cIter, errLog := r.gitRepo.Log(&git.LogOptions{
			From:  *revision,
			Order: git.LogOrderCommitterTime,
		})
		if errLog != nil {
			return errLog
		}
		errRR = cIter.ForEach(func(c *object.Commit) error {
			if c.Message != "" {
				split := strings.Split(c.Message, "\n")
				if len(split) > 0 {
					if split[0] == firstLine {
						cmt := newCommit(c)
						wantCommit = &cmt
					}
				}
			}
			return nil
		})
		if errRR != nil {
			return errRR
		}
		return nil
	})
	if errForEach != nil {
		return nil, fmt.Errorf("failed to ForEach tags: %w", errForEach)
	}
	if wantCommit == nil {
		return nil, fmt.Errorf("can not find tagBy FristLine: %s", firstLine)
	}
	return wantCommit, nil
}

// TagLatestByCommitTime
//
//	latest tag find by commit time, please ensure that the time of the device submitting the tag is synchronized correctly.
//	check at: git show-ref --tag
//
// return latest tag
func (r *repo) TagLatestByCommitTime() (*object.Tag, error) {
	tagObjs, err := r.gitRepo.TagObjects()
	if err != nil {
		return nil, fmt.Errorf("failed to get tagObjs: %w", err)
	}
	var wantTag *object.Tag
	commitTime := new(time.Time)
	defer tagObjs.Close()
	for {
		obj, errNext := tagObjs.Next()
		if errNext != nil {
			break
		}
		commit, errNext := obj.Commit()
		if errNext != nil {
			continue
		}
		//name := obj.Name
		//strings.TrimSpace(name)
		//objString := obj.String()
		//strings.TrimSpace(objString)
		//commitHash := commit.Hash.String()
		//strings.TrimSpace(commitHash)
		commitWhen := commit.Author.When
		if commitWhen.After(*commitTime) {
			commitTime = &commitWhen
			wantTag = obj
		}
	}
	if wantTag == nil {
		return nil, fmt.Errorf("can not find latest tag")
	}
	return wantTag, nil
}

func (r *repo) Commit(commitMessage string, paths ...string) error {
	gitWorktree, err := r.gitRepo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get git worktree: %w", err)
	}

	for _, path := range paths {
		if errAdd := gitWorktree.AddWithOptions(&git.AddOptions{
			Path: path,
		}); errAdd != nil {
			return fmt.Errorf("failed to git add %s: %w", path, errAdd)
		}
	}

	if _, errCommit := gitWorktree.Commit(commitMessage, &git.CommitOptions{}); errCommit != nil {
		return fmt.Errorf("failed to git commit: %w", errCommit)
	}

	return nil
}

// RemoteInfo
//
// remote most is git.OriginDefault
//
// configUrlIndex most is 0
//
// return Repository from git.Repository
func (r *repo) RemoteInfo(remoteName string, configUrlIndex int) (*git_info.GitRemoteInfo, error) {
	remote, err := r.gitRepo.Remote(remoteName)
	if err != nil {
		return nil, err
	}
	remoteConfig := remote.Config()
	if remoteConfig == nil {
		return nil, fmt.Errorf("RepositoryFistRemoteInfo remote: %s not found", remote)
	}
	if len(remoteConfig.URLs) < configUrlIndex {
		return nil, fmt.Errorf("RepositoryFistRemoteInfo remote: %s URLs is empty", remote)
	}
	urlStr := remoteConfig.URLs[configUrlIndex]
	parse, err := giturls.Parse(urlStr)
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
		return &git_info.GitRemoteInfo{
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
	return &git_info.GitRemoteInfo{
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

func (r *repo) logWithStopFn(fromHash *plumbing.Hash, beginFn, endFn stopFn) ([]Commit, error) {
	cIter, err := r.gitRepo.Log(&git.LogOptions{
		From:  *fromHash,
		Order: git.LogOrderCommitterTime,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to git log: %w", err)
	}

	commits := make([]Commit, 0, defaultCommitCount)

	if errFind := cIter.ForEach(newIterFn(&commits, beginFn, endFn)); errFind != nil {
		return nil, fmt.Errorf("failed to iterate each git log: %w", errFind)
	}

	return commits, nil
}

type stopFn func(*object.Commit) error

func stopAtHash(hash *plumbing.Hash) stopFn {
	return func(c *object.Commit) error {
		if c.Hash == *hash {
			return storer.ErrStop
		}

		return nil
	}
}

func newIterFn(commits *[]Commit, beginFn, endFn stopFn) func(c *object.Commit) error {
	if beginFn == nil && endFn == nil {
		return func(c *object.Commit) error {
			commit := newCommit(c)
			*commits = append(*commits, commit)

			return nil
		}
	}

	if beginFn == nil {
		return func(c *object.Commit) error {
			commit := newCommit(c)
			*commits = append(*commits, commit)

			if err := endFn(c); err != nil {
				return err
			}

			return nil
		}
	}

	if endFn == nil {
		return func(c *object.Commit) error {
			if err := beginFn(c); err != nil {
				return err
			}

			commit := newCommit(c)
			*commits = append(*commits, commit)

			return nil
		}
	}

	return func(c *object.Commit) error {
		if err := beginFn(c); err != nil {
			return err
		}

		commit := newCommit(c)
		*commits = append(*commits, commit)

		if err := endFn(c); err != nil {
			return err
		}

		return nil
	}
}

// SetAuthMethod
// auth transport.AuthMethod
//
//	auth, err := ssh.NewPublicKeysFromFile("git", valSshKeyPath, valSshKeyPassWord)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	repository.SetAuthMethod(auth)
func (r *repo) SetAuthMethod(auth transport.AuthMethod) {
	r.transportAuthMethod = &auth
}

// SetProxyOptions
// options transport.ProxyOptions
func (r *repo) SetProxyOptions(options transport.ProxyOptions) {
	r.transportProxyOptions = &options
}

// PullOrigin
// warning: direct support public repository, private repository must use SetAuth before.
// like: git pull origin
func (r *repo) PullOrigin() error {
	if r.gitRepo == nil {
		return fmt.Errorf("git pull origin error, repo not init, or can not read")
	}
	worktree, err := r.gitRepo.Worktree()
	if err != nil {
		return fmt.Errorf("git pull origin error, want get work tree err: %v", err)
	}

	options := git.PullOptions{
		RemoteName: r.remoteName,
	}

	if r.transportAuthMethod != nil {
		options.Auth = *r.transportAuthMethod
	}

	if r.transportProxyOptions != nil {
		options.ProxyOptions = *r.transportProxyOptions
	}

	errPull := worktree.Pull(&options)
	if errPull != nil {
		if errors.Is(errPull, git.NoErrAlreadyUpToDate) {
			return nil
		}
		return errPull
	}

	return nil
}

// FetchTags
// warning: direct support public repository, private repository must use SetAuth before.
// fetch tags
// like run as: git fetch --tags
func (r *repo) FetchTags() error {
	//remotes, err := r.gitRepo.Remotes()
	//if err != nil {
	//	return err
	//}
	//if len(remotes) == 0 {
	//	return fmt.Errorf("no remote found")
	//}
	//remote := remotes[0]
	remote, err := r.gitRepo.Remote(r.remoteName)
	if err != nil {
		return fmt.Errorf("failed fetch --tags to find remote %w", err)
	}
	options := git.FetchOptions{
		RemoteName: r.remoteName,
		Tags:       git.AllTags,
	}

	if r.transportAuthMethod != nil {
		options.Auth = *r.transportAuthMethod
	}

	if r.transportProxyOptions != nil {
		options.ProxyOptions = *r.transportProxyOptions
	}

	errFetch := remote.Fetch(&options)
	if errFetch != nil {
		if errors.Is(errFetch, git.NoErrAlreadyUpToDate) {
			return nil
		}
		return fmt.Errorf("failed to fetch --tags err: %w", errFetch)
	}

	return nil
}
