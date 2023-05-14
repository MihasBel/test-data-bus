package busserver

type Config struct {
	Host             string `env:"GRPC_URL" envDefault:":9080"`
	SecureConnection bool   `env:"GRPC_SECURE_CONNECTION" envDefault:"false"`
	StartTimeout     int    `env:"GRPC_START_TIMEOUT" envDefault:"10"`
	StopTimeout      int    `env:"GRPC_STOP_TIMEOUT" envDefault:"10"`
}
