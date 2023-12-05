[![ci](https://github.com/sinlov-go/go-git-tools/actions/workflows/ci.yml/badge.svg)](https://github.com/sinlov-go/go-git-tools/actions/workflows/ci.yml)
[![license](https://img.shields.io/github/license/sinlov-go/go-git-tools)](https://github.com/sinlov-go/go-git-tools)
[![go mod version](https://img.shields.io/github/go-mod/go-version/sinlov-go/go-git-tools?label=go.mod)](https://github.com/sinlov-go/go-git-tools)
[![GoDoc](https://godoc.org/github.com/sinlov-go/go-git-tools?status.png)](https://godoc.org/github.com/sinlov-go/go-git-tools/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlov-go/go-git-tools)](https://goreportcard.com/report/github.com/sinlov-go/go-git-tools)
[![codecov](https://codecov.io/gh/sinlov-go/go-git-tools/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov-go/go-git-tools)
[![GitHub release)](https://img.shields.io/github/v/release/sinlov-go/go-git-tools)](https://github.com/sinlov-go/go-git-tools/releases)

## for what

- this project used to management git local project with golang

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/sinlov-go/go-git-tools)](https://github.com/sinlov-go/go-git-tools/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

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

- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## env

- minimum go version: go 1.19
- change `go 1.19`, `^1.19`, `1.19.13` to new go version

### libs

| lib                                 | version |
|:------------------------------------|:--------|
| https://github.com/stretchr/testify | v1.8.4  |
| https://github.com/go-git/go-git    | v5.7.0  |
| https://github.com/whilp/git-urls   | v1.0.0  |

## usage

```go
package main

import (
	"github.com/sinlov-go/go-git-tools/git_info"
	"testing"
)

func TestFoo(t *testing.T) {
	var projectRootPath = "/Users/sinlov/go/src/github.com/sinlov-go/go-git-tools"
	url, err := git_info.RepositoryFistRemoteInfo(projectRootPath, "origin")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("url.Host: %s", url.Host)
	t.Logf("url.User: %s", url.User)
	t.Logf("url.Repo: %s", url.Repo)

	branchByPath, err := git_info.RepositoryNowBranchByPath(projectRootPath)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("branchByPath: %s", branchByPath)
}

```

# dev

```bash
# It needs to be executed after the first use or update of dependencies.
$ make init dep
```

- test code

```bash
$ make test testBenchmark
```

add main.go file and run

```bash
# run at env dev use cmd/main.go
$ make dev
```

- ci to fast check

```bash
# check style at local
$ make style

# run ci at local
$ make ci
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```
