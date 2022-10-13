package env

import "time"

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	ApiTimeout   time.Duration
}

var ServerSetting = &Server{}
