package conf

import (
	"sync"

	"github.com/dizzrt/ellie/config"
)

var (
	ac   *AppConfig
	once sync.Once
)

type AppConfig struct {
	Server   Server   `mapstructure:"server"`
	Log      Log      `mapstructure:"log"`
	Registry Registry `mapstructure:"registry"`
	DB       DB       `mapstructure:"db"`
}

type Server struct {
	GRPC GRPCServer `mapstructure:"grpc"`
	HTTP HTTPServer `mapstructure:"http"`
}

type GRPCServer struct {
	Addr string `mapstructure:"addr"`
}

type HTTPServer struct {
	Addr string `mapstructure:"addr"`
}

type Log struct {
	File       string `mapstructure:"file"`
	Symlink    string `mapstructure:"symlink"`
	Level      string `mapstructure:"level"`
	MaxAge     string `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	OutputType string `mapstructure:"output_type"`
}

type DB struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Addr     string `mapstructure:"addr"`
}

type Registry struct {
	Addr string `mapstructure:"addr"`
}

func GetAppConfig() *AppConfig {
	once.Do(func() {
		c := config.NewStdViperConfig()
		if err := c.Load(); err != nil {
			panic(err)
		}

		var tmp AppConfig
		if err := c.Unmarshal(&tmp); err != nil {
			panic(err)
		}

		ac = &tmp
	})

	return ac
}
