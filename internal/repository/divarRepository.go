package repository

import (
	"agahi-plus-plus/internal/model"
	"agahi-plus-plus/internal/response"
)

type DivarRepository interface {
	GetPostTokens(url, apikey, accessToken string) (*response.GetPostsResponse, error)
	EditPost(endpoint, apikey, accessToken string, post *model.Post) error
}
