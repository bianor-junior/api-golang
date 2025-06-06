package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)



var cfg *conf

type conf struct {
	DBDriver string `mapstructure:"BD_DRIVER"`
	DBHost   string `mapstructure:"DB_HOST"`
	DBPort   string `mapstructure:"DB_PORT"`
	DBUser   string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName   string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret string `mapstructure:"JWT_SECRET"`
	JWTExpireIn string `mapstructure:"JWT_EXPIRE_IN"`
	TokenAuth *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	err := viper.Unmarshal(&cfg); if err != nil {
		panic(err.Error())
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, nil
}