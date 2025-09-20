package repository

import "agahi-plus-plus/internal/response"

type DivarRepository interface {
	GetPostTokens(url, apikey, accessToken string) (*response.GetPostsResponse, error)
}
