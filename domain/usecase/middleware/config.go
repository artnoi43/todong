package middleware

type Config struct {
	SecretKey string `mapstructure:"secret_key" yaml:"secretKey,omitempty"`
}
