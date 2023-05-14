package memorybus

import "time"

type Config struct {
	MsgTypes    []string      `env:"MSG_TYPES" required:"true"`
	MaxMsgs     int           `env:"MAX_MSGS" envDefault:"10"`
	ReadDelayMS time.Duration `env:"READ_DELAY_MS" envDefault:"10ms"`
}
