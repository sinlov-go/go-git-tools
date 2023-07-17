package git_test

import (
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
		t.Fatal(err)
	}
	t.Logf("commits len %d", len(commits))
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
				URL: "https://github.com/sinlov-go/go-git-tools",
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
