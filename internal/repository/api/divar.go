package api

import (
	"agahi-plus-plus/internal/helper"
	"agahi-plus-plus/internal/model"
	"agahi-plus-plus/internal/repository"
	"agahi-plus-plus/internal/response"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
)

type divarApi struct {
	config *helper.ServiceConfig
	logger *zap.Logger
}

func NewDivarApi(
	config *helper.ServiceConfig,
	logger *zap.Logger,
) repository.DivarRepository {
	return &divarApi{
		config: config,
		logger: logger,
	}
}

func (i divarApi) GetPostTokens(url, apikey, accessToken string) (*response.GetPostsResponse, error) {
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Add("x-api-key", apikey)
	req.Header.Add("Authorization", "Bearer "+accessToken)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}
	var resp response.GetPostsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &resp, nil
}

type editPostRequest struct {
	Description string   `json:"description"`
	ImagePaths  []string `json:"image_paths"`
	Title       string   `json:"title"`
}

func (i divarApi) EditPost(endpoint, apikey, accessToken string, post *model.Post) error {
	url := getPostUrl(endpoint, post.Token)
	method := "PUT"

	reqPayload := editPostRequest{
		Description: post.Description,
		ImagePaths:  post.GetImages(),
		Title:       post.Title,
	}

	j, jErr := json.Marshal(reqPayload)
	if jErr != nil {
		return jErr
	}
	payload := strings.NewReader(string(j))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", apikey)
	req.Header.Add("Authorization", "Bearer "+accessToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	return nil
}
