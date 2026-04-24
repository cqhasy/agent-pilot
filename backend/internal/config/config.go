package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port string

	FeishuBaseURL           string
	FeishuAppID             string
	FeishuAppSecret         string
	FeishuEventMode         string
	FeishuVerificationToken string
	FeishuEncryptKey        string

	OpenAIAPIKey  string
	OpenAIBaseURL string
	OpenAIModel   string

	AssistantName string
}

func LoadFromEnv() (Config, error) {
	cfg := Config{
		Port:                    getEnv("PORT", "8080"),
		FeishuBaseURL:           getEnv("FEISHU_BASE_URL", "https://open.feishu.cn"),
		FeishuAppID:             os.Getenv("FEISHU_APP_ID"),
		FeishuAppSecret:         os.Getenv("FEISHU_APP_SECRET"),
		FeishuEventMode:         getEnv("FEISHU_EVENT_MODE", "webhook"),
		FeishuVerificationToken: os.Getenv("FEISHU_VERIFICATION_TOKEN"),
		FeishuEncryptKey:        os.Getenv("FEISHU_ENCRYPT_KEY"),
		OpenAIAPIKey:            os.Getenv("OPENAI_API_KEY"),
		OpenAIBaseURL:           getEnv("OPENAI_BASE_URL", "https://api.openai.com/v1"),
		OpenAIModel:             getEnv("OPENAI_MODEL", "gpt-4.1-mini"),
		AssistantName:           getEnv("ASSISTANT_NAME", "Feishu Dev Assistant"),
	}

	switch {
	case cfg.FeishuAppID == "":
		return Config{}, fmt.Errorf("FEISHU_APP_ID is required")
	case cfg.FeishuAppSecret == "":
		return Config{}, fmt.Errorf("FEISHU_APP_SECRET is required")
	case cfg.OpenAIAPIKey == "":
		return Config{}, fmt.Errorf("OPENAI_API_KEY is required")
	case cfg.FeishuEventMode != "webhook" && cfg.FeishuEventMode != "ws":
		return Config{}, fmt.Errorf("FEISHU_EVENT_MODE must be one of: webhook, ws")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
