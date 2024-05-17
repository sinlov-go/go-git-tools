[![ci](https://github.com/sinlov-go/go-git-tools/actions/workflows/ci.yml/badge.svg)](https://github.com/sinlov-go/go-git-tools/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/sinlov-go/go-git-tools?label=go.mod)](https://github.com/sinlov-go/go-git-tools)
[![GoDoc](https://godoc.org/github.com/sinlov-go/go-git-tools?status.png)](https://godoc.org/github.com/sinlov-go/go-git-tools)
[![goreportcard](https://goreportcard.com/badge/github.com/sinlov-go/go-git-tools)](https://goreportcard.com/report/github.com/sinlov-go/go-git-tools)

[![GitHub license](https://img.shields.io/github/license/sinlov-go/go-git-tools)](https://github.com/sinlov-go/go-git-tools)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/sinlov-go/go-git-tools)](https://github.com/sinlov-go/go-git-tools/tags)
[![GitHub release)](https://img.shields.io/github/v/release/sinlov-go/go-git-tools)](https://github.com/sinlov-go/go-git-tools/releases)
[![codecov](https://codecov.io/gh/sinlov-go/go-git-tools/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov-go/go-git-tools)

## for what

- this project used to management git local project with golang

## depends

in go mod project

```bash
# warning use private git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "https://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q https://github.com/sinlov-go/go-git-tools.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/sinlov-go/go-git-tools
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/sinlov-go/go-git-tools | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## Features

### pkg `github.com/sinlov-go/go-git-tools/git`

```go
package foo

import (
	"github.com/sinlov-go/go-git-tools/git"
)
```

- [x] `git.NewRepositoryByPath(gitRootPath)` to load local git project
    - [x] `HeadReference() (*plumbing.Reference, error)`  get head reference
    - [x] `HeadBranchName() (string, error)` get head branch name
    - [x] `CheckHasSubmodules() (bool, error)` check has submodules
    - [x] `CheckSubmodulesIsDirty() (bool, error)` check submodules is dirty
    - [x] `CheckLocalBranchIsDirty() (bool, error)` check local branch is dirty
    - [x] `IsCitCmdAvailable() bool` check git command is available (v1.13+)
    - [x] `CheckWorkTreeIsDirtyWithGitCmd() (bool, error)` check work tree is dirty by `git status --porcelain` (v1.13+)
- [x] get log or commit
    - [x] `Log(fromRev, toRev string) ([]Commit, error)` return all commits between <from revision> and <to revision>
    - [x] `CommitLatestTagByTime() (*Commit, error)` return latest tag commit
    - [x] `CommitTagSearchByName(tagName string) (*Commit, error)` return tag commit
    - [x] `CommitTagSearchByFirstLine(firstLine string) (*Commit, error)` return tag commit
    - [x] `Commit(commitMessage string, paths ...string) error` commit

- [x] support setting ` SetAuthMethod(auth transport.AuthMethod)` and `SetAuthMethod(auth transport.AuthMethod)` for
  private git remote

```go
package foo_test

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

	repository.SetAuth(auth)

	err = repository.PullOrigin()
	if err != nil {
		t.Fatal(err)
	}
}
```

### pkg `github.com/sinlov-go/go-git-tools/git_info`

```go
package foo

import (
	"github.com/sinlov-go/go-git-tools/git_info"
)
```

- [x] `git_info` package to get info of repository
    - [x] `IsPathUnderGitManagement` check path is under git management
    - [x] `RepositoryFistRemoteInfo` get repository first remote info
    - [x] `RepositoryHeadByPath` get repository head by local path
    - [x] `RepositoryNowBranchByPath` get repository now branch by local path
    - [x] `RepositoryConfigPath` get repository config by local path

- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

# dev

- see [doc/dev.md](doc/dev.md)