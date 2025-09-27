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
	lastSelectedImages := post.GetSelectedImages()
	imagesLen := 0

	for _, imageUrl := range selectedImages {
		if s.contains(lastSelectedImages, imageUrl) {
			continue
		}
		imagesLen++
	}

	if imagesLen == 0 {
		return post, balance, nil
	}
	if imagesLen > balance {
		return post, balance, errors.New("not enough token")
	}

	lastNewImages := post.GetNewImages()

	var outputs []string
	for _, imageUrl := range selectedImages {
		if s.contains(lastSelectedImages, imageUrl) {
			continue
		}
		o, generateErr := s.promptRepo.Generate(ctx, imageUrl)
		if generateErr != nil {
			continue
		}
		outputs = append(outputs, o)
	}

	outputs = append(outputs, lastNewImages...)

	user.Balance -= imagesLen

	s.userService.Update(user)

	jb, _ := postgres.MakeJsonb(outputs)
	post.NewImages = jb
	jb2, _ := postgres.MakeJsonb(selectedImages)
	post.SelectedImages = jb2

	s.postService.UpdatePost(post)

	return post, user.Balance, nil
}

func (s promptService) contains(sl []string, str string) bool {
	for _, v := range sl {
		if v == str {
			return true
		}
	}

	return false
}
