package git

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Commit stores all git-commit information
type Commit struct {
	Author  Author
	Message string
	Hash    plumbing.Hash
}

// Convert from git-commit
func newCommit(commit *object.Commit) Commit {
	return Commit{
		Message: commit.Message,
		Hash:    commit.Hash,
		Author: Author{
			Name:  commit.Author.Name,
			Email: commit.Author.Email,
			When:  commit.Author.When,
		},
	}
}
