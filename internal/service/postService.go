package service

import (
	"agahi-plus-plus/internal/helper"
	"agahi-plus-plus/internal/model"
	"agahi-plus-plus/internal/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
)

type PostService interface {
	Get(user *model.User, serviceName string) (*model.Post, int, error)
	GetPostByUser(ctx *gin.Context, serviceName string) (*model.Post, int, error)
	UpdatePost(post *model.Post) error
}

type postService struct {
	postApiRepo repository.PostApiRepo
	postDbRepo  repository.PostDbRepo
	userService UserService
	config      *helper.ServiceConfig
	logger      *zap.Logger
}

func NewPostService(
	postApiRepo repository.PostApiRepo,
	postDbRepo repository.PostDbRepo,
	userService UserService,
	config *helper.ServiceConfig,
	logger *zap.Logger,
) PostService {
	return &postService{
		postApiRepo: postApiRepo,
		postDbRepo:  postDbRepo,
		userService: userService,
		config:      config,
		logger:      logger,
	}
}

func (s postService) Get(user *model.User, serviceName string) (*model.Post, int, error) {
	if user.PostToken == "" {
		return nil, 0, errors.New("token required")
	}

	post, err := s.postDbRepo.Get(user.PostToken)
	if err != nil {
		return nil, 0, err
	}

	if post == nil {
		post, err = s.postApiRepo.Get(user.PostToken, serviceName)
		if err != nil {
			return nil, 0, err
		}

		post.UserID = user.ID
		post, err = s.postDbRepo.Insert(post)
		if err != nil {
			log.Println("post service line 65 err:", err.Error())
			return nil, 0, err
		}

		return post, user.Balance, nil
	}

	return post, user.Balance, nil
}

func (s postService) GetPostByUser(ctx *gin.Context, serviceName string) (*model.Post, int, error) {
	user, err := s.userService.GetUserWithContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	return s.Get(user, serviceName)
}

func (s postService) UpdatePost(post *model.Post) error {
	return s.postDbRepo.Update(post)
}
