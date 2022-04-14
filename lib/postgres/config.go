package postgres

type Config struct {
	Host     string `mapstructure:"host" yaml:"host,omitempty"`
	Port     string `mapstructure:"port" yaml:"port,omitempty"`
	User     string `mapstructure:"user" yaml:"user,omitempty"`
	Password string `mapstructure:"password" yaml:"password,omitempty"`
	DBName   string `mapstructure:"name" yaml:"dbName,omitempty"`
}
