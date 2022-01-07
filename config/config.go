package config

import "github.com/spf13/viper"

type Config struct {
	DBHost      string `mapstructure:"POSTGRES_HOST"`
	DBPort      string `mapstructure:"POSTGRES_PORT"`
	DBUser      string `mapstructure:"POSTGRES_USER"`
	DBPass      string `mapstructure:"POSTGRES_PASSWORD"`
	DBName      string `mapstructure:"POSTGRES_NAME"`
	RedisHost   string `mapstructure:"REDIS_HOST"`
	RedisPort   string `mapstructure:"REDIS_PORT"`
	ElasticHost string `mapstructure:"ELASTIC_HOST"`
	ElasticPort string `mapstructure:"ELASTIC_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)

	return config, err
}
