package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// User AI config
type Config struct {
	Provider string
	APIKey   string
	Model    string
	Style    string
	Language string
}

// Loading cfg
func LoadConfig() (*Config, error) {
	viper.SetConfigName("scarycommit")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetDefault("Provider", "openrouter")
	viper.SetDefault("Model", "minimax/minimax-m2:free")
	viper.SetDefault("Style", "conventional")
	viper.SetDefault("Language", "en")

	_ = viper.ReadInConfig()

	return &Config{
		Provider: viper.GetString("Provider"),
		APIKey:   viper.GetString("APIKey"),
		Model:    viper.GetString("Model"),
		Style:    viper.GetString("Style"),
		Language: viper.GetString("Language"),
	}, nil
}

// Default cfg
func SaveDefaultConfig() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("❌ Failed to get current directory:", err)
		return
	}

	path := filepath.Join(cwd, "scarycommit.yaml")

	viper.Set("Provider", "openrouter")
	viper.Set("APIKey", "API_KEY_HERE")
	viper.Set("Model", "minimax/minimax-m2:free")
	viper.Set("Style", "conventional")
	viper.Set("Language", "en")

	err = viper.WriteConfigAs(path)
	if err != nil {
		fmt.Println("❌ Failed to save default config:", err)
		return
	}

	fmt.Println("✅ Saved default config to", path)
}
