package api

import (
	"agahi-plus-plus/internal/helper"
	"agahi-plus-plus/internal/repository"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type promptApi struct {
	logger *zap.Logger
	config *helper.ServiceConfig
}

func NewPromptApi(logger *zap.Logger, config *helper.ServiceConfig) repository.PromptRepository {
	return &promptApi{
		logger: logger,
		config: config,
	}
}

type Part struct {
	Text       string      `json:"text,omitempty"`
	InlineData *InlineData `json:"inlineData,omitempty"`
}

type InlineData struct {
	MimeType string `json:"mime_type"`
	Data     string `json:"data"`
}

type Content struct {
	Role  string `json:"role"`
	Parts []Part `json:"parts"`
}

type GenerationConfig struct {
	ResponseModalities []string `json:"response_modalities"`
}

type Payload struct {
	Contents         []Content        `json:"contents"`
	GenerationConfig GenerationConfig `json:"generation_config"`
}

type Response struct {
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Content Content `json:"content"`
}

func (r promptApi) Generate(ctx *gin.Context, imageUrl string) (string, error) {
	imgData, err := r.downloadImage(imageUrl)
	if err != nil {
		panic(fmt.Sprintf("Error reading image file: %v", err))
	}
	imgB64 := base64.StdEncoding.EncodeToString(imgData)

	log.Printf("\nMessage: %s\nRole: %s\nUrl: %s\n", r.config.Prompt.Message, r.config.Prompt.Role, r.config.Prompt.Url)

	payload := Payload{
		Contents: []Content{
			{
				Role: r.config.Prompt.Role,
				Parts: []Part{
					{
						Text: r.config.Prompt.Message,
					},
					{
						InlineData: &InlineData{
							MimeType: "image/jpeg",
							Data:     imgB64,
						},
					},
				},
			},
		},
		GenerationConfig: GenerationConfig{
			ResponseModalities: []string{"TEXT", "IMAGE"},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic(fmt.Sprintf("Error marshaling JSON: %v", err))
	}

	req, err := http.NewRequest("POST", r.config.Prompt.Url, bytes.NewReader(payloadBytes))
	if err != nil {
		panic(fmt.Sprintf("Error creating request: %v", err))
	}

	req.Header.Set("Authorization", "Bearer "+r.config.Prompt.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("Error sending request: %v", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Println(fmt.Sprintf("Request failed with status %d: %s", resp.StatusCode, string(body)))
		return "", fmt.Errorf("Request failed with status %d: %s", resp.StatusCode, string(body))
	}

	//respBytes, err := io.ReadAll(resp.Body)
	//log.Println("BODY: ", string(respBytes))
	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println(fmt.Sprintf("Error decoding response: %v", err))
		return "", err
	}

	log.Printf("RESULT: %+v", result)
	outFile, err := r.saveGeminiImage(result, r.config.Prompt.OutputPath)
	if err != nil {
		log.Println(fmt.Sprintf("Error saving image: %v", err))
		return "", err
	}

	fmt.Printf("Saved %s\n", outFile)

	siteOut := strings.Replace(outFile, "imgs/", "https://cdnapp.agahi-plus.ir/", 1)
	return siteOut, err
}

func (r promptApi) downloadImage(url string) ([]byte, error) {
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

func (r promptApi) saveGeminiImage(result Response, outputDir string) (string, error) {
	if len(result.Candidates) == 0 {
		return "", fmt.Errorf("no candidates in response")
	}

	parts := result.Candidates[0].Content.Parts
	for _, p := range parts {
		if p.InlineData != nil && p.InlineData.Data != "" {
			data, err := base64.StdEncoding.DecodeString(p.InlineData.Data)
			if err != nil {
				return "", fmt.Errorf("error decoding base64 image data: %v", err)
			}

			timestamp := time.Now().Format("20060102-150405")
			uniqueFilenameTmp := fmt.Sprintf("generated_image_%s.png", timestamp)
			uniqueFilename := fmt.Sprintf("generated_image_%s.jpg", timestamp)

			outPathTmp := filepath.Join(outputDir, uniqueFilenameTmp)

			if err := os.WriteFile(outPathTmp, data, 0644); err != nil {
				return "", fmt.Errorf("error writing png image file: %v", err)
			}

			outPath := filepath.Join(outputDir, uniqueFilename)

			cmd := exec.Command("gm", "convert", outPathTmp, outPath)
			if err := cmd.Run(); err != nil {
				return "", fmt.Errorf("error converting image: %v", err)
			}

			if err = os.WriteFile(outPath, data, 0644); err != nil {
				return "", fmt.Errorf("error writing image file: %v", err)
			}

			err = os.Remove(outPathTmp)
			if err != nil {
				return "", err
			}

			return outPath, nil
		}
	}

	return "", fmt.Errorf("no image part found in response")
}
