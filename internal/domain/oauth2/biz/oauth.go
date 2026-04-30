package biz

import "github.com/dizzrt/dauth/internal/domain/oauth2/repo"

var _ OAuth2Biz = (*oauth2Biz)(nil)

type OAuth2Biz interface {
}

type oauth2Biz struct {
	oauth2Repo repo.OAuth2Repo
}

func NewOAuth2Biz(oauth2Repo repo.OAuth2Repo) OAuth2Biz {
	return &oauth2Biz{
		oauth2Repo: oauth2Repo,
	}
}
