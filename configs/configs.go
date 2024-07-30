package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var configs *config

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
	//Set name of the config
	viper.SetConfigName("config")

	//Set type of the config, example, yaml, env, toml etc.
	viper.SetConfigType("env")

	//Add path from config file passed in function parameter
	viper.AddConfigPath(path)

	//Catch file that will be used to load config
	viper.SetConfigFile(".env")

	//This is a optional option that will give priority to variables defined not exported
	viper.AutomaticEnv()

	//Read config passed previously from viper configuration
	err := viper.ReadInConfig()
	if err != nil {

		panic(err)
	}

	//Unmarshal the config into config struct
	err = viper.Unmarshal(&configs)
	if err != nil {
		panic(err)
	}

	//That will allows our program use generate JWT Tokens
	configs.TokenAuth = jwtauth.New("HS256", []byte(configs.JWTSecret), nil)

	//Return our config structure with our .env configurations
	return configs, nil
}
