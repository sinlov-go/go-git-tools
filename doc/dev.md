## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/sinlov-go/go-git-tools)](https://github.com/sinlov-go/go-git-tools/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

# env

## env

- minimum go version: go 1.19
- change `go 1.19`, `^1.19`, `1.19.13` to new go version

### libs

| lib                                 | version |
|:------------------------------------|:--------|
| https://github.com/stretchr/testify | v1.8.4  |
| https://github.com/go-git/go-git    | v5.10.1 |
| https://github.com/chainguard-dev/git-urls   | v1.0.2  |

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
