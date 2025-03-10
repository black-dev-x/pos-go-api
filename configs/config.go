package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type configuration struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBName       string `mapstructure:"DB_DATABASE"`
	DBUser       string `mapstructure:"DB_USERNAME"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	WebPort      string `mapstructure:"PORT"`
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	JWTAuth      *jwtauth.JWTAuth
}

func LoadConfig() (*configuration, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config configuration
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	config.JWTAuth = jwtauth.New("HS256", []byte(config.JWTSecret), nil)
	println("Config loaded")
	println(config.DBHost)
	return &config, nil
}
