package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort    string `mapstructure:"SERVER_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBDsn         string
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	config.makeDbConfig()

	return
}

func (c *Config) makeDbConfig() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort)
	c.DBDsn = dsn
}
