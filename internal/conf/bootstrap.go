package conf

import "github.com/dizzrt/ellie/config"

type Bootstrap struct {
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

func NewBootstrap() *Bootstrap {
	c := config.NewStdViperConfig()
	if err := c.Load(); err != nil {
		panic(err)
	}

	bootstrap, err := buildBootstrap(c)
	if err != nil {
		panic(err)
	}

	return bootstrap
}

func buildBootstrap(c config.Config) (*Bootstrap, error) {
	var bootstrap Bootstrap
	if err := c.UnmarshalKey("server", &bootstrap.Server); err != nil {
		return nil, err
	}

	if err := c.UnmarshalKey("log", &bootstrap.Log); err != nil {
		return nil, err
	}

	if err := c.UnmarshalKey("db", &bootstrap.DB); err != nil {
		return nil, err
	}

	if err := c.UnmarshalKey("registry", &bootstrap.Registry); err != nil {
		return nil, err
	}

	return &bootstrap, nil
}
