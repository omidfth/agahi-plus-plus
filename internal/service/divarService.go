package service

import (
	"agahi-plus-plus/internal/helper"
	"agahi-plus-plus/internal/model"
	"agahi-plus-plus/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DivarService interface {
	EditPost(ctx *gin.Context, serviceName string) (*model.Post, error)
}

type divarService struct {
	divarApi    repository.DivarRepository
	postService PostService
	userService UserService
	config      *helper.ServiceConfig
	logger      *zap.Logger
}

func NewDivarService(
	divarApi repository.DivarRepository,
	postService PostService,
	userService UserService,
	config *helper.ServiceConfig,
	logger *zap.Logger,
) DivarService {
	return &divarService{
		divarApi:    divarApi,
		userService: userService,
		postService: postService,
		config:      config,
		logger:      logger,
	}
}

func (s divarService) EditPost(ctx *gin.Context, serviceName string) (*model.Post, error) {
	post, _, err := s.postService.GetPostByUser(ctx, serviceName)
	if err != nil {
		return nil, err
	}

	user, err := s.userService.GetUserWithContext(ctx)
	if err != nil {
		return post, err
	}

	config := s.config.GetDivarConfig(serviceName)
	err = s.divarApi.EditPost(s.config.Divar.Api.EditPost, config.ApiKey, user.AccessToken, post)
	if err != nil {
		return post, err
	}
	post.IsConnected = true
	s.postService.UpdatePost(post)

	return post, err
}
