package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RepositorySuite struct {
	suite.Suite

	r *git.Repository

	repo *repo
}

func (s *RepositorySuite) SetupTest() {
	var err error
	s.r, err = git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/sinlov-go/go-git-tools",
	})
	s.NoError(err)

	s.repo = &repo{
		gitRepo: s.r,
	}
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (s *RepositorySuite) TestLogSuccess() {
	_, gotErr := s.repo.Log("", "")
	s.NoError(gotErr)
}
