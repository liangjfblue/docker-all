package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPConf  *HTTPConfig
	MysqlConf *MysqlConfig
	RedisConf *RedisConfig
	TokenConf *TokenConfig
}

type HTTPConfig struct {
	RunMode      string
	Addr         string
	Name         string
	ReadTimeout  int
	WriteTimeout int
}

type MysqlConfig struct {
	Addr         string
	Db           string
	User         string
	Password     string
	MaxIdleConnS int
	MaxOpenConnS int
}

type RedisConfig struct {
	Host        string
	Port        string
	ClusterHost string
	IsCluster   bool
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

type TokenConfig struct {
	Secret     string
	SecretTime int
}

func NewConfig() *Config {
	return &Config{
		HTTPConf: &HTTPConfig{
			RunMode: viper.GetString("runmode"),
			Addr:    viper.GetString("addr"),
			Name:    viper.GetString("name"),
		},

		MysqlConf: &MysqlConfig{
			Addr:         viper.GetString("mysql.addr"),
			Db:           viper.GetString("mysql.db"),
			User:         viper.GetString("mysql.user"),
			Password:     viper.GetString("mysql.password"),
			MaxIdleConnS: viper.GetInt("mysql.maxIdleConns"),
			MaxOpenConnS: viper.GetInt("mysql.maxOpenConns"),
		},
		RedisConf: &RedisConfig{
			Host:        viper.GetString("redis.host"),
			Port:        viper.GetString("redis.port"),
			ClusterHost: viper.GetString("redis.cluster_host"),
			IsCluster:   viper.GetBool("redis.is_cluster"),
			MaxIdle:     viper.GetInt("redis.maxIdle"),
			MaxActive:   viper.GetInt("redis.maxActive"),
			IdleTimeout: viper.GetInt("redis.idleTimeout"),
		},
		TokenConf: &TokenConfig{
			Secret:     viper.GetString("token.jwtSecretShort"),
			SecretTime: viper.GetInt("token.jwtSecretShortTime"),
		},
	}
}

func init() {
	if err := initConfig(); err != nil {
		panic(err)
	}
	watchConfig()
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("link-gin-models")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s\n", e.Name)
	})
}
