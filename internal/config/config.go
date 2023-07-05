package config

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"0.0.0.0"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	Storage      StorageConfig `yaml:"storage"`
	JwtKey       string        `yaml:"jwt_key" env-required:"true"`
	SocketJwtKey string        `yaml:"socket_jwt_key" env-required:"true"`
	AppVersion   string        `yaml:"app_version" env-required:"true"`
	MaxFileSize  int64         `yaml:"max_file_size" env-required:"true"`
	Email1       string        `yaml:"email_1" env-required:"true"`
	Email2       string        `yaml:"email_2" env-required:"true"`
	Email3       string        `yaml:"email_3" env-required:"true"`

	Password1 string `yaml:"password_1" env-required:"true"`
	Password2 string `yaml:"password_2" env-required:"true"`
	Password3 string `yaml:"password_3" env-required:"true"`

	EmailPort int    `yaml:"email_port" env-required:"true"`
	EmailHost string `yaml:"email_host" env-required:"true"`
}

type StorageConfig struct {
	PgPoolMaxConn int    `yaml:"pg_pool_max_conn" env-required:"true"`
	Host          string `json:"host"`
	Port          string `json:"port"`
	Database      string `json:"database"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

// type EmailConfig struct {
// 	Email     string `yaml:"email" env-required:"true"`
// 	Password  string `yaml:"password" env-required:"true"`
// 	EmailPort int    `yaml:"email_port" env-required:"true"`
// 	EmailHost string `yaml:"email_host" env-required:"true"`
// }

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {

		// TODO path config
		pathConfig := "./../../config.yml"

		instance = &Config{}

		if err := cleanenv.ReadConfig(pathConfig, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			fmt.Println(help)
			//logger.Info(help)
			//logger.Fatal(err)
		}
	})
	return instance
}
