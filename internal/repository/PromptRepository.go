package repository

import "github.com/gin-gonic/gin"

type PromptRepository interface {
	Generate(ctx *gin.Context, imageUrl string) (string, error)
}
