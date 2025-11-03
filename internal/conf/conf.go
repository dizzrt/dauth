package conf

type Bootstrap struct {
	Server Server `mapstructure:"server"`
	Log    Log    `mapstructure:"log"`
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
