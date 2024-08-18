package configs

import (
	"fmt"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type env_variables struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	WebServerPort string
	JWTSecret     string
	JWTExpiresIn  int
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig() (*env_variables, error) {

	var env_config *env_variables

	viper.SetConfigName("env_variables")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&env_config)

	env_config.TokenAuth = jwtauth.New("HS256", []byte(env_config.JWTSecret), nil)

	return env_config, err
}
