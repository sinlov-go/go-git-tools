# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [1.14.0](https://github.com/sinlov-go/go-git-tools/compare/v1.13.0...v1.14.0) (2025-05-08)

### ✨ Features

* update dependencies and lint settings to go 1.23+ ([70c5aaf1](https://github.com/sinlov-go/go-git-tools/commit/70c5aaf1fe9318ee8cd4c439514cecbb8f8f9836))

* update README v1.14.+ must with (go 1.23.+) ([6fe8913c](https://github.com/sinlov-go/go-git-tools/commit/6fe8913c469c7eb6ab91a1be5e610d252eba639e))

### 👷‍ Build System

* update Go version to 1.23 and add build ID ([edd9012f](https://github.com/sinlov-go/go-git-tools/commit/edd9012f2814b28d013243c898ae2b0097eec115))

* bump codecov/codecov-action from 5.4.0 to 5.4.2 ([98a458ee](https://github.com/sinlov-go/go-git-tools/commit/98a458ee2997970106f102c251f15f3327a17273))

* bump codecov/codecov-action from 5.3.1 to 5.4.0 ([c32a8091](https://github.com/sinlov-go/go-git-tools/commit/c32a80915907fa599b9e4470a4ecb203b84951b3))

* bump codecov/codecov-action from 5.1.2 to 5.3.1 ([4ee1c484](https://github.com/sinlov-go/go-git-tools/commit/4ee1c484d64d50f6c3ba12c9657ed8933c43e142))

* bump codecov/codecov-action from 5.0.7 to 5.1.2 ([54742a11](https://github.com/sinlov-go/go-git-tools/commit/54742a11f4615e52a7fa41e8dc5ac0ef4165a892))

* bump codecov/codecov-action from 4.6.0 to 5.0.7 ([eb1fcd74](https://github.com/sinlov-go/go-git-tools/commit/eb1fcd747456dae258e75920b1684bdd0c305a42))

* bump codecov/codecov-action from 4.5.0 to 4.6.0 ([d932f10b](https://github.com/sinlov-go/go-git-tools/commit/d932f10b95dffef522cfa47677a6514dae502170))

* bump codecov/codecov-action from 4.4.1 to 4.5.0 ([1388b312](https://github.com/sinlov-go/go-git-tools/commit/1388b31202b265b297e1e62d5edaa270488d44a3))

* bump codecov/codecov-action from 4.4.0 to 4.4.1 ([71071092](https://github.com/sinlov-go/go-git-tools/commit/71071092168a0621033fbed675d0a14429de4616))

## [1.13.0](https://github.com/sinlov-go/go-git-tools/compare/v1.12.0...v1.13.0) (2024-05-18)

### 🐛 Bug Fixes

* add `CheckWorkTreeIsDirtyWithGitCmd` to try fix windows error ([c4a3d7e4](https://github.com/sinlov-go/go-git-tools/commit/c4a3d7e45dc13b45b5b73c94770c00c0b65e745d)), fix [#28](https://github.com/sinlov-go/go-git-tools/issues/28)

### ✨ Features

* add `IsCitCmdAvailable` to check git command available ([9d88a383](https://github.com/sinlov-go/go-git-tools/commit/9d88a383b09b4986d8db92fd6858f3230ab5c3d2))

## [1.12.0](https://github.com/sinlov-go/go-git-tools/compare/v1.11.0...v1.12.0) (2024-05-17)

### ✨ Features

* add `SetProxyOptions` and `SetAuthMethod` and `Check*` for local ([b820153c](https://github.com/sinlov-go/go-git-tools/commit/b820153c483e3c115b074298b2a1b9e1703fcfce)), feat [#24](https://github.com/sinlov-go/go-git-tools/issues/24)

* update full build pipline ([ccf28a34](https://github.com/sinlov-go/go-git-tools/commit/ccf28a34b7b95f6b621211d8abc955f413a359c2))

### ♻ Refactor

* let method to repository_impl.go ([7f7f2615](https://github.com/sinlov-go/go-git-tools/commit/7f7f2615de2306733f820ad06537d1c34a3e1876))

* update exmplate_test ([df8e3780](https://github.com/sinlov-go/go-git-tools/commit/df8e378061e9aff27a718690a5e1ab0178ed8b3b))

### 👷‍ Build System

* bump golangci/golangci-lint-action from 5 to 6 ([54465509](https://github.com/sinlov-go/go-git-tools/commit/54465509ec07f1eb7baf1afa172bc63242c34217))

* bump codecov/codecov-action from 4.3.0 to 4.4.0 ([095e0ae2](https://github.com/sinlov-go/go-git-tools/commit/095e0ae285d2a9301843190abd303e08ca24707d))

* bump convention-change/conventional-version-check ([cf016366](https://github.com/sinlov-go/go-git-tools/commit/cf016366d71d89b40419dd2c28535e59cf17c39a))

* bump codecov/codecov-action from 4.1.1 to 4.3.0 ([b5e553fe](https://github.com/sinlov-go/go-git-tools/commit/b5e553fe3f58b84888c56dd0f4e20d9f9ffb1ca6))

* bump golangci/golangci-lint-action from 4 to 5 ([87369c37](https://github.com/sinlov-go/go-git-tools/commit/87369c37993a3015ceddae570be6e4815e52e8e1))

* bump codecov/codecov-action from 4.1.0 to 4.1.1 ([964a8c4e](https://github.com/sinlov-go/go-git-tools/commit/964a8c4e4c4b86649bec92166086a3efe3d5e2bb))

* bump github.com/go-git/go-git/v5 from 5.11.0 to 5.12.0 ([d72a41e6](https://github.com/sinlov-go/go-git-tools/commit/d72a41e6fc011a06eb60a4a56acc22edf256189b))

* bump codecov/codecov-action from 4.0.0 to 4.1.0 ([bce115c6](https://github.com/sinlov-go/go-git-tools/commit/bce115c6dd4a43bcf333a9d44331fa761b3865e4))

* bump golangci/golangci-lint-action from 3 to 4 ([1983ebef](https://github.com/sinlov-go/go-git-tools/commit/1983ebef0ed0d6c48480ea7576a309037be23fb9))

* bump codecov/codecov-action from 3.1.4 to 4.0.0 ([7614df32](https://github.com/sinlov-go/go-git-tools/commit/7614df32e6105c94f7b32e95f7c1f30a25c7dea3))

* changet to tag_changes: ${{ needs.version.outputs.cc_changes }} ([65a1ff5b](https://github.com/sinlov-go/go-git-tools/commit/65a1ff5b5fb10946debc4702a95184c862e7538e))

## [1.11.0](https://github.com/sinlov-go/go-git-tools/compare/v1.10.0...v1.11.0) (2024-01-22)

### ✨ Features

* change to github.com/chainguard-dev/git-urls v1.0.2 and update new action pipline ([c70694d6](https://github.com/sinlov-go/go-git-tools/commit/c70694d6ab1bbcaf3b6d20504d4d0480387eeb62))

### 👷‍ Build System

* bump actions/setup-go from 4 to 5 ([bae845ae](https://github.com/sinlov-go/go-git-tools/commit/bae845aec62209ec797acd18ddafad4c1ac55a8f))

* bump actions/upload-artifact from 3 to 4 ([2a4088a0](https://github.com/sinlov-go/go-git-tools/commit/2a4088a040c999c7d5dcbd0afc1e3df97211f69b))

* bump actions/download-artifact from 3 to 4 ([6ee22d40](https://github.com/sinlov-go/go-git-tools/commit/6ee22d403d79325dccb81881986257822c41073b))

* bump github.com/go-git/go-git/v5 from 5.10.1 to 5.11.0 ([88788203](https://github.com/sinlov-go/go-git-tools/commit/88788203f2e80d19c0cc3b81ea9c0b3f5276ef71))

* change golangci/golangci-lint-action use version latest ([f05cc81f](https://github.com/sinlov-go/go-git-tools/commit/f05cc81fceb1ab79df49dbfb70d270457a96b066))

## [1.10.0](https://github.com/sinlov-go/go-git-tools/compare/v1.9.1...v1.10.0) (2023-12-05)

### ✨ Features

* update to go1.19 for github.com/go-git/go-git/v5 v5.10.1 ([e6703fb3](https://github.com/sinlov-go/go-git-tools/commit/e6703fb3fe5248611f016253f0c83da8c55ef5cb)), fix [#10](https://github.com/sinlov-go/go-git-tools/issues/10)

### 👷‍ Build System

* bump github.com/go-git/go-git/v5 from 5.10.0 to 5.10.1 ([7cae8002](https://github.com/sinlov-go/go-git-tools/commit/7cae800244c9ee1f21bbe4e036bc3929d7023183))

* bump github.com/go-git/go-git/v5 from 5.9.0 to 5.10.0 ([b644fab6](https://github.com/sinlov-go/go-git-tools/commit/b644fab692e860f409540d443a59a1a229792290))

* update go.mod build ([c6f9e964](https://github.com/sinlov-go/go-git-tools/commit/c6f9e96432019e8a3e32a21cd2cd3b3642d18ca5))

* bump github.com/go-git/go-git/v5 from 5.8.1 to 5.9.0 ([fca1a3f9](https://github.com/sinlov-go/go-git-tools/commit/fca1a3f9d62f296ab81a37f5427ca0ae81233a62))

* bump github.com/go-git/go-billy/v5 from 5.4.1 to 5.5.0 ([f3b3c9e1](https://github.com/sinlov-go/go-git-tools/commit/f3b3c9e1e28b0239974b2c9af1c1f9f90a761cc9))

* bump actions/checkout from 3 to 4 ([a3389a85](https://github.com/sinlov-go/go-git-tools/commit/a3389a8535c96291ae526bef6f7118dd7c85e962))

### [1.9.1](https://github.com/sinlov-go/go-git-tools/compare/v1.9.0...v1.9.1) (2023-07-31)

## [1.9.0](https://github.com/sinlov-go/go-git-tools/compare/v1.8.1...v1.9.0) (2023-07-31)

### 👷‍ Build System

* **deps:** bump github.com/go-git/go-git/v5 from 5.7.0 to 5.8.0 ([eb83869](https://github.com/sinlov-go/go-git-tools/commit/eb8386992105a6dba1e7ea81f3f500984c3aed33))

* **deps:** bump github.com/go-git/go-git/v5 from 5.8.0 to 5.8.1 ([a5b769e](https://github.com/sinlov-go/go-git-tools/commit/a5b769ea0e61bf8afd78476780b96dee8ed4605b))

### ✨ Features

* let git log Order by LogOrderCommitterTime ([9c8b68f](https://github.com/sinlov-go/go-git-tools/commit/9c8b68f682f9bf3265198bd631e215efc3de3e82))

## [1.9.0](https://github.com/sinlov-go/go-git-tools/compare/v1.8.1...v1.9.0) (2023-07-31)

### 👷‍ Build System

* **deps:** bump github.com/go-git/go-git/v5 from 5.7.0 to 5.8.0 ([eb83869](https://github.com/sinlov-go/go-git-tools/commit/eb8386992105a6dba1e7ea81f3f500984c3aed33))

* **deps:** bump github.com/go-git/go-git/v5 from 5.8.0 to 5.8.1 ([a5b769e](https://github.com/sinlov-go/go-git-tools/commit/a5b769ea0e61bf8afd78476780b96dee8ed4605b))

### ✨ Features

* let git log Order by LogOrderCommitterTime ([9c8b68f](https://github.com/sinlov-go/go-git-tools/commit/9c8b68f682f9bf3265198bd631e215efc3de3e82))

### [1.8.1](https://github.com/sinlov-go/go-git-tools/compare/v1.8.0...v1.8.1) (2023-07-21)

### 🐛 Bug Fixes

* test case error at CI env, so fix this ([af5f671](https://github.com/sinlov-go/go-git-tools/commit/af5f671ac1919b224b5062f67ba5466373232577))

## [1.8.0](https://github.com/sinlov-go/go-git-tools/compare/v1.7.0...v1.8.0) (2023-07-21)

### ✨ Features

* add git.HeadReference and git.HeadBranchName ([de631c6](https://github.com/sinlov-go/go-git-tools/commit/de631c6162d3696b226fadcb407d61f946e5d2fd))

* add git.RemoteInfo to get RemoteInfo ([c6652da](https://github.com/sinlov-go/go-git-tools/commit/c6652da18148a6ad8fafb0b4a6a56974e12fa383))

## [1.7.0](https://github.com/sinlov-go/go-git-tools/compare/v1.6.0...v1.7.0) (2023-07-21)

### ✨ Features

* add git.NewRepositoryRemoteByPath and git.NewRepositoryRemoteClone ([e910096](https://github.com/sinlov-go/go-git-tools/commit/e910096face51f8da3f7547f7db7ea0691b02c04))

## [1.6.0](https://github.com/sinlov-go/go-git-tools/compare/v1.5.0...v1.6.0) (2023-07-20)

### ✨ Features

* add method FetchTags ([7817371](https://github.com/sinlov-go/go-git-tools/commit/7817371a2d4dfecc5a0bd861bf171a23d332e1cc))

### 🐛 Bug Fixes

* fix RepositoryFistRemoteInfo at http url error ([e09ea2d](https://github.com/sinlov-go/go-git-tools/commit/e09ea2daf26bc639b17bcc5ef5994787a5dbdb13))

## [1.5.0](https://github.com/sinlov-go/go-git-tools/compare/v1.4.0...v1.5.0) (2023-07-20)

### ✨ Features

* add git.CommitLatestTagByTime and TagLatestByCommitTime for find tag latest by commit time ([9b0c6d8](https://github.com/sinlov-go/go-git-tools/commit/9b0c6d83dc5816687005d21967f6efb060ef45aa))

## [1.4.0](https://github.com/sinlov-go/go-git-tools/compare/v1.3.0...v1.4.0) (2023-07-19)

### ✨ Features

* git.repo add method CommitTagSearchByName and CommitTagSearchByFirstLine to get git.Commit ([9393ff1](https://github.com/sinlov-go/go-git-tools/commit/9393ff1e0856af721040c357fb315408486b03e1))

## [1.3.0](https://github.com/sinlov-go/go-git-tools/compare/v1.2.3...v1.3.0) (2023-07-18)

### ✨ Features

* add NewRepositoryClone  NewRepositoryByPath ([70ae0a8](https://github.com/sinlov-go/go-git-tools/commit/70ae0a87ce7b2ef740980c9597f0502d1c0ae236))

### [1.2.3](https://github.com/sinlov-go/go-git-tools/compare/v1.2.2...v1.2.3) (2023-07-14)

### 👷‍ Build System

* open artifact upload ([ae43f1a](https://github.com/sinlov-go/go-git-tools/commit/ae43f1a37fadc52ff9501f7180616c92db4eca4d))

### [1.2.2](https://github.com/sinlov-go/go-git-tools/compare/v1.2.1...v1.2.2) (2023-07-14)

### 👷‍ Build System

* change upload-artifact setting ([498c0b5](https://github.com/sinlov-go/go-git-tools/commit/498c0b539c6692dc3d4d8662668add4560e7762a))

### [1.2.1](https://github.com/sinlov-go/go-git-tools/compare/v1.2.0...v1.2.1) (2023-07-14)

### 👷‍ Build System

* fix go-release upload-artifact ([682957d](https://github.com/sinlov-go/go-git-tools/commit/682957d7efd7861592b3aeff4bc093eaa39afb6c))

## [1.2.0](https://github.com/sinlov-go/go-git-tools/compare/v1.1.0...v1.2.0) (2023-07-14)

### ✨ Features

* let Commit.Hash load by commit object ([a77090f](https://github.com/sinlov-go/go-git-tools/commit/a77090f2f170f71a6ee791c0c7a5e8b6953db464))

## [1.1.0](https://github.com/sinlov-go/go-git-tools/compare/v1.0.0...v1.1.0) (2023-07-07)

### ✨ Features

* add package git for operation git repo ([67315d9](https://github.com/sinlov-go/go-git-tools/commit/67315d99fdb71622ab98a190d9fe9b22b237cb78))

## 1.0.0 (2023-07-04)
