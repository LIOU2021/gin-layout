package config

import (
	"github.com/LIOU2021/gin-layout/env"
)

var EnvStructSlice []interface{}

func init() {
	EnvStructSlice = append(EnvStructSlice, env.LogSetting)
	EnvStructSlice = append(EnvStructSlice, env.AppSetting)
}
