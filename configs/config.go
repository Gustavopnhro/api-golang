package configs

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*config, error) {
	var cfg *config

	// Resolve the absolute path of the directory provided
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("error resolving absolute path: %v", err)
	}

	// Set name of the config
	viper.SetConfigName("app_config")

	// Set type of the config, example, yaml, env, toml etc.
	viper.SetConfigType("env")

	// Add path from config file passed in function parameter
	viper.AddConfigPath(absPath)

	// Catch file that will be used to load config
	configFilePath := filepath.Join(absPath, ".env")
	viper.SetConfigFile(configFilePath)

	// Print the path being used for debugging
	log.Printf("Attempting to load configuration from: %s", configFilePath)

	// This is an optional option that will give priority to variables defined not exported
	viper.AutomaticEnv()

	// Read config passed previously from viper configuration
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	// Unmarshal the config into config struct
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %v", err)
	}

	// This will allow our program to use generate JWT Tokens
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	// Return our config structure with our .env configurations
	return cfg, nil
}
