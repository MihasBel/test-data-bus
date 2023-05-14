package memorybus

import "time"

type Config struct {
	MsgTypes     []string      `env:"MSG_TYPES" required:"true"`
	MinOffset    int64         `env:"MIN_OFFSET" envDefault:"10"` //TODO to clean old msgs
	ReadDelaySec time.Duration `env:"READ_DELAY_Sec" envDefault:"2s"`
}
