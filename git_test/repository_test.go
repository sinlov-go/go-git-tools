package git_test

import (
	"fmt"
	"github.com/go-git/go-billy/v5"
	goGit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/sinlov-go/go-git-tools/git"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestNewRepositoryByPath(t *testing.T) {
	currentFolderPath, err := getCurrentFolderPath()
	if err != nil {
		t.Fatal(err)
	}
	gitRootPath := filepath.Dir(currentFolderPath)

	repository, err := git.NewRepositoryByPath(gitRootPath)
	if err != nil {
		t.Fatal(err)
	}
	commits, err := repository.Log("", "")
	if err != nil {
		t.Logf("get commits err %s", err)
	} else {
		t.Logf("commits len %d", len(commits))
	}
}

func TestNewRepositoryClone(t *testing.T) {
	// mock NewRepositoryClone
	tests := []struct {
		name     string
		s        storage.Storer
		worktree billy.Filesystem
		o        *goGit.CloneOptions
		wantErr  bool
		logErr   error
	}{
		{
			name:     "http url",
			s:        memory.NewStorage(),
			worktree: nil,
			o: &goGit.CloneOptions{
				URL: "https://github.com/sinlov-go/go-git-tools.git",
			},
		},
		{
			name:     "git url",
			s:        memory.NewStorage(),
			worktree: nil,
			o: &goGit.CloneOptions{
				URL: "git@github.com:sinlov-go/go-git-tools.git",
			},
			wantErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do NewRepositoryClone
			gotResult, gotErr := git.NewRepositoryClone(tc.s, tc.worktree, tc.o)

			// verify NewRepositoryClone
			assert.Equal(t, tc.wantErr, gotErr != nil)
			if tc.wantErr {
				t.Logf("gotErr: %v", gotErr)
				return
			}
			commits, errLog := gotResult.Log("", "")
			assert.Equal(t, tc.logErr, errLog)
			t.Logf("git commits len: %v", len(commits))
		})
	}
}

func TestCommitTagSearchByName(t *testing.T) {
	// mock CommitTagSearchByName
	tests := []struct {
		name                   string
		s                      storage.Storer
		worktree               billy.Filesystem
		o                      *goGit.CloneOptions
		tagName                string
		wantTagHash            string
		wantCloneErr           bool
		wantTagSearchByNameErr error
	}{
		{
			name:     "tag empty",
			s:        memory.NewStorage(),
			worktree: nil,
			o: &goGit.CloneOptions{
				URL: "https://github.com/sinlov-go/go-git-tools.git",
			},
			tagName:                "",
			wantTagHash:            "",
			wantCloneErr:           false,
			wantTagSearchByNameErr: fmt.Errorf("commit tag search by name is empty"),
		},
		{
			name:     "tag not exist",
			s:        memory.NewStorage(),
			worktree: nil,
			o: &goGit.CloneOptions{
				URL: "https://github.com/sinlov-go/go-git-tools.git",
			},
			tagName:                "0.1.2",
			wantCloneErr:           false,
			wantTagSearchByNameErr: fmt.Errorf("can not find commit at tag: 0.1.2"),
		},
		{
			name:     "tag v1.3.0",
			s:        memory.NewStorage(),
			worktree: nil,
			o: &goGit.CloneOptions{
				URL: "https://github.com/sinlov-go/go-git-tools.git",
			},
			tagName:     "v1.3.0",
			wantTagHash: "231ccf643cffdc30722aa9f9bbc9b8e4ae817bde",
		},
		{
			name:     "tag v1.2.3",
			s:        memory.NewStorage(),
			worktree: nil,
			o: &goGit.CloneOptions{
				URL: "https://github.com/sinlov-go/go-git-tools.git",
			},
			tagName:     "v1.2.3",
			wantTagHash: "4bb54b0fefb2a60f478397fcc2722999c658d72b",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do CommitTagSearchByName
			gotResult, gotErr := git.NewRepositoryClone(tc.s, tc.worktree, tc.o)

			// verify CommitTagSearchByName
			assert.Equal(t, tc.wantCloneErr, gotErr != nil)
			if tc.wantCloneErr {
				t.Logf("gotErr: %v", gotErr)
				return
			}
			commit, errTagSearch := gotResult.CommitTagSearchByName(tc.tagName)
			assert.Equal(t, tc.wantTagSearchByNameErr, errTagSearch)
			if errTagSearch != nil {
				return
			}
			assert.Equal(t, tc.wantTagHash, commit.Hash.String())
		})
	}
}

func TestCommitTagSearchByFirstLine(t *testing.T) {
	// mock CommitTagSearchByFirstLine
	tests := []struct {
		name                   string
		s                      storage.Storer
		worktree               billy.Filesystem
		o                      *goGit.CloneOptions
		firstLine              string
		wantTagHash            string
		wantCloneErr           bool
		wantTagSearchByNameErr error
	}{
		{
			name:     "first line empty",
			s:        memory.NewStorage(),
			worktree: nil,
			o: &goGit.CloneOptions{
				URL: "https://github.com/sinlov-go/go-git-tools.git",
			},
			firstLine:              "",
			wantTagHash:            "",
			wantCloneErr:           false,
			wantTagSearchByNameErr: fmt.Errorf("commit tag search by firstLine is empty"),
		},
		{
			name:     "first line not exist",
			s:        memory.NewStorage(),
			worktree: nil,
			o: &goGit.CloneOptions{
				URL: "https://github.com/sinlov-go/go-git-tools.git",
			},
			firstLine:              "foo first line",
			wantCloneErr:           false,
			wantTagSearchByNameErr: fmt.Errorf("can not find tagBy FristLine: foo first line"),
		},
		{
			name:     "first line release 1.3.0",
			s:        memory.NewStorage(),
			worktree: nil,
			o: &goGit.CloneOptions{
				URL: "https://github.com/sinlov-go/go-git-tools.git",
			},
			firstLine:   "chore(release): 1.3.0",
			wantTagHash: "231ccf643cffdc30722aa9f9bbc9b8e4ae817bde",
		},
		{
			name:     "first line release 1.2.3",
			s:        memory.NewStorage(),
			worktree: nil,
			o: &goGit.CloneOptions{
				URL: "https://github.com/sinlov-go/go-git-tools.git",
			},
			firstLine:   "chore(release): 1.2.3",
			wantTagHash: "4bb54b0fefb2a60f478397fcc2722999c658d72b",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do CommitTagSearchByFirstLine
			gotResult, gotErr := git.NewRepositoryClone(tc.s, tc.worktree, tc.o)

			// verify CommitTagSearchByFirstLine
			assert.Equal(t, tc.wantCloneErr, gotErr != nil)
			if tc.wantCloneErr {
				t.Logf("gotErr: %v", gotErr)
				return
			}
			commit, errTagSearch := gotResult.CommitTagSearchByFirstLine(tc.firstLine)
			assert.Equal(t, tc.wantTagSearchByNameErr, errTagSearch)
			if errTagSearch != nil {
				return
			}
			assert.Equal(t, tc.wantTagHash, commit.Hash.String())
		})
	}
}

func TestCommitLatestTag(t *testing.T) {
	// mock CommitLatestTagByTime
	currentFolderPath, err := getCurrentFolderPath()
	if err != nil {
		t.Fatal(err)
	}
	gitRootPath := filepath.Dir(currentFolderPath)
	tests := []struct {
		name string

		cloneUrl     string
		wantCloneErr bool

		repoLocalPath    string
		wantLocalPathErr bool

		wantLatestTagErr bool
	}{
		{
			name:     "has tag clone",
			cloneUrl: "https://github.com/sinlov-go/go-git-tools.git",
		},
		{
			name:          "has tag local",
			repoLocalPath: gitRootPath,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do CommitLatestTagByTime
			var gotCloneErr error
			var gotLocalErr error
			var gotResult git.Repository
			if tc.cloneUrl != "" {
				result, err := git.NewRepositoryClone(memory.NewStorage(), nil, &goGit.CloneOptions{
					URL: tc.cloneUrl,
				})
				gotResult = result
				gotCloneErr = err
			}

			if tc.repoLocalPath != "" {
				result, err := git.NewRepositoryByPath(tc.repoLocalPath)
				gotResult = result
				gotLocalErr = err
			}

			// verify CommitLatestTagByTime
			assert.Equal(t, tc.wantCloneErr, gotCloneErr != nil)
			assert.Equal(t, tc.wantLocalPathErr, gotLocalErr != nil)
			if tc.wantCloneErr {
				t.Logf("gotErr: %v", gotCloneErr)
				return
			}
			if tc.wantLocalPathErr {
				t.Logf("gotErr: %v", gotLocalErr)
				return
			}

			commitLatestTag, gotLatestTagErr := gotResult.CommitLatestTagByTime()

			assert.Equal(t, tc.wantLatestTagErr, gotLatestTagErr != nil)

			t.Logf("commitLatestTag Message %s", commitLatestTag.Message)
			hash := commitLatestTag.Hash
			assert.False(t, hash.IsZero())
			t.Logf("commitLatestTag Hash %s", hash.String())
		})
	}
}
