package pkg

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/informeai/gia/dto"
)

// GitWrapper is struct for git wrapper
type GitWrapper struct {
	repo *git.Repository
}

// NewGitWrapper return instance of git wrapper
func NewGitWrapper() *GitWrapper {
	return &GitWrapper{}
}

// Init execute clone and initialize repo in memory
func (g *GitWrapper) Init(url string) error {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return err
	}
	g.repo = r
	return nil
}

func (g *GitWrapper) Authors() ([]dto.Author, error) {
	var authors []dto.Author
	ref, err := g.repo.Head()
	if err != nil {
		return nil, err
	}

	cIter, err := g.repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return nil, err
	}

	arrAuthors := make(map[string]dto.Author)

	err = cIter.ForEach(func(c *object.Commit) error {
		if _, ok := arrAuthors[c.Author.Email]; ok {
			author := arrAuthors[c.Author.Email]
			author.CommitCount++
			arrAuthors[c.Author.Email] = author
		} else {
			arrAuthors[c.Author.Email] = dto.Author{
				Name:        c.Author.Name,
				Email:       c.Author.Email,
				CommitCount: 1,
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	for _, alth := range arrAuthors {
		authors = append(authors, alth)
	}
	return authors, nil
}
