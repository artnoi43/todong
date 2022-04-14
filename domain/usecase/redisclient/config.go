package redisclient

type Config struct {
	Address  string `mapstructure:"address" yaml:"address,omitempty"`
	Username string `mapstructure:"username" yaml:"username,omitempty"`
	Password string `mapstructure:"password" yaml:"password,omitempty"`
	DB       int    `mapstructure:"db" yaml:"db,omitempty"`
}
