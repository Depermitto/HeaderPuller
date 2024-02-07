package repo

import (
	"github.com/go-git/go-git/v5"
	"strings"
)

func IsRepoLink(s string) bool {
	return strings.HasPrefix(s, "http")
}

func DefaultOptions(repoLink string) *git.CloneOptions {
	return &git.CloneOptions{
		URL:   repoLink,
		Depth: 1,
	}
}
