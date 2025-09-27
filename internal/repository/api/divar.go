package api

import (
	"agahi-plus-plus/internal/helper"
	"agahi-plus-plus/internal/model"
	"agahi-plus-plus/internal/repository"
	"agahi-plus-plus/internal/response"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"go.uber.org/zap"
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

	images := post.GetAllImages()
	uploadUrl, err := i.getUploadUrl(apikey)

	if err != nil {
		return err
	}
	var path []string
	for _, image := range images {
		p, uErr := i.uploadImage(uploadUrl, image)
		if uErr != nil {
			log.Printf("failed to upload image: %v", uErr)
		}
		path = append(path, p)
	}

	reqPayload := editPostRequest{
		Description: post.Description,
		ImagePaths:  path,
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

type GetUploadUrl struct {
	UploadUrl string `json:"upload_url"`
}

func (i divarApi) getUploadUrl(apiKey string) (string, error) {
	url := i.config.Divar.Api.UploadImage
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("X-API-Key", apiKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(body))

	var uploadUrl GetUploadUrl
	err = json.Unmarshal(body, &uploadUrl)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return uploadUrl.UploadUrl, nil
}

type UploadImagePath struct {
	Path string `json:"path"`
}

func (i divarApi) uploadImage(uploadUrl string, fileUrl string) (string, error) {
	url := uploadUrl
	method := "POST"
	data, _ := i.downloadImage(fileUrl)
	payload := strings.NewReader(string(data))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "image/jpeg")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(body))

	var uploadImagePath UploadImagePath
	err = json.Unmarshal(body, &uploadImagePath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return uploadImagePath.Path, nil
}

func (i divarApi) downloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error downloading image: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download failed with status: %d", resp.StatusCode)
	}

	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading image data: %v", err)
	}

	return imgData, nil
}
