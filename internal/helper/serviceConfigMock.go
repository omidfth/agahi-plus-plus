package helper

func NewServiceConfigMock() *ServiceConfig {
	return &ServiceConfig{
		Prompt: prompt(struct {
			ApiKey     string
			Url        string
			Role       string
			Message    string
			OutputPath string
		}{ApiKey: "aa-IGbV86bG6Axly2U1z5qdvcC7KNiRO4y6lObdYxL3GY2LA9de",
			Url:        "https://api.avalai.ir/v1beta/models/gemini-2.5-flash-image-preview:generateContent",
			Role:       "user",
			Message:    "Remove everything in this house except the carpets and output an image of the house empty of furniture. But leave the structure of the house and the camera angle unchanged.",
			OutputPath: "./unit_test/",
		}),
	}
}
