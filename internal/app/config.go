package app

import (
	"github.com/MihasBel/test-data-bus/internal/memorybus"
	"time"

	"github.com/MihasBel/test-data-bus/delivery/grpc/busserver"
)

type Config struct {
	LogLevel     string `env:"LOG_LEVEL" envDefault:"info"`
	BrokerConfig memorybus.Config
	GRPCConfig   busserver.Config

	StartTimeout time.Duration `env:"START_TIMEOUT"`
	StopTimeout  time.Duration `env:"STOP_TIMEOUT"`
}
