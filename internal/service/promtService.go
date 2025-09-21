package service

import (
	"agahi-plus-plus/internal/model"
	"agahi-plus-plus/internal/postgres"
	"agahi-plus-plus/internal/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PromptService interface {
	Generate(ctx *gin.Context, selectedImages []string, serviceName string) (*model.Post, int, error)
}

type promptService struct {
	promptRepo  repository.PromptRepository
	postService PostService
	userService UserService
	logger      *zap.Logger
}

func NewPromptService(
	promptRepo repository.PromptRepository,
	postService PostService,
	userService UserService,
	logger *zap.Logger,
) PromptService {
	return &promptService{
		promptRepo:  promptRepo,
		postService: postService,
		userService: userService,
		logger:      logger,
	}
}

func (s promptService) Generate(ctx *gin.Context, selectedImages []string, serviceName string) (*model.Post, int, error) {
	user, err := s.userService.GetUserWithContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	post, balance, err := s.postService.GetPostByUser(ctx, serviceName)
	if err != nil {
		return post, balance, err
	}

	imagesLen := len(selectedImages)

	if imagesLen == 0 {
		return post, balance, nil
	}

	if imagesLen > balance {
		return post, balance, errors.New("not enough token")
	}
	var outputs []string
	for _, imageUrl := range selectedImages {
		o, generateErr := s.promptRepo.Generate(ctx, imageUrl)
		if generateErr != nil {
			continue
		}
		outputs = append(outputs, o)
	}

	user.Balance -= imagesLen

	s.userService.Update(user)

	jb, _ := postgres.MakeJsonb(outputs)
	post.NewImages = jb
	jb2, _ := postgres.MakeJsonb(selectedImages)
	post.SelectedImages = jb2

	s.postService.UpdatePost(post)

	return post, user.Balance, nil
}
