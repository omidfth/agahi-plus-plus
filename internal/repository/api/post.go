package api

import (
	"agahi-plus-plus/internal/helper"
	"agahi-plus-plus/internal/model"
	"agahi-plus-plus/internal/postgres"
	"agahi-plus-plus/internal/repository"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strings"
)

type postApi struct {
	config *helper.ServiceConfig
	logger *zap.Logger
}

func NewPostApi(
	config *helper.ServiceConfig,
	logger *zap.Logger,
) repository.PostApiRepo {
	return &postApi{
		config: config,
		logger: logger,
	}
}

func (i postApi) Get(token string, serviceName string) (*model.Post, error) {
	log.Println(i.config.Divar)
	endPoint := getPostUrl(i.config.Divar.Api.GetPost, token)
	log.Println("endpoint: ", endPoint)
	method := "GET"
	req, err := http.NewRequest(method, endPoint, nil)
	if err != nil {
		return nil, err
	}

	config := i.config.GetDivarConfig(serviceName)

	req.Header.Add("x-api-key", config.ApiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("method:getDivarPost err status_code:", resp.StatusCode)
		return nil, errors.New(resp.Status)
	}

	var response getPostResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, errors.New("error while decode response")
	}
	return response.toPostModel()
}

func getPostUrl(endpoint string, postToken string) string {
	return strings.Replace(endpoint, "{{token}}", postToken, 1)
}

type getPostResponse struct {
	Token       string   `json:"token"`
	Category    string   `json:"category"`
	City        string   `json:"city"`
	District    string   `json:"district"`
	ChatEnabled bool     `json:"chat_enabled"`
	Data        postData `json:"data"`
}

type postData struct {
	Title  string   `json:"title"`
	Images []string `json:"images"`
}

func (r getPostResponse) toPostModel() (*model.Post, error) {
	title := r.Data.Title

	jd, _ := postgres.MakeJsonb(r.Data.Images)

	return &model.Post{
		Token:       r.Token,
		IsConnected: false,
		Title:       title,
		Images:      jd,
	}, nil
}
