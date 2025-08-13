package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicKey   string
	BotToken	string
	AppID		string
	Port		string
}

func Load() Config {
	_ = godotenv.Load()

	c := Config{
		PublicKey: getEnv("PUBLIC_KEY"),
		BotToken:  getEnv("DISCORD_TOKEN"),
		AppID:     getEnv("APP_ID"),
		Port:      getEnvOr("PORT", "8080"),
	}

	if err := c.Validate(); err != nil {
		panic(err)
	}

	return c
}


func (c Config) Validate() error {
	var missing []string
	if c.PublicKey == "" {
		missing = append(missing, "DISCORD_PUBLIC_KEY")
	}
	if c.BotToken == "" {
		missing = append(missing, "DISCORD_BOT_TOKEN")
	}
	if c.AppID == "" {
		missing = append(missing, "DISCORD_APP_ID")
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required env vars: %s", strings.Join(missing, ", "))
	}

	return nil
}

func (c Config) Addr() string { return ":" + c.Port }

func getEnv(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}
func getEnvOr(key, def string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return def
}
