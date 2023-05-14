package memorybus

import "time"

type Config struct {
	MsgTypes     []string      `env:"MSG_TYPES" required:"true"`
	MaxMsgs      int           `env:"MAX_MSGS" envDefault:"10"`
	ReadDelaySec time.Duration `env:"READ_DELAY_Sec" envDefault:"2s"`
}
