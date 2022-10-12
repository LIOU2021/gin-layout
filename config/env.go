package config

import (
	"log"

	"github.com/LIOU2021/gin-layout/config/env"
	"github.com/LIOU2021/gin-layout/helpers"
	"github.com/go-ini/ini"
)

// type Log struct {
// 	FileName string
// }

// var LogSetting = &Log{}

var cfg *ini.File

var EnvStructSlice []interface{}

// Setup initialize the configuration instance
func envRegister() {
	var err error
	cfg, err = ini.Load(".env")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse '.env': %v", err)
	}

	EnvStructSlice = append(EnvStructSlice, env.LogSetting)
	foreachMapTo(EnvStructSlice)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

func foreachMapTo(v []interface{}) {
	for _, element := range v {
		mapTo(helpers.GetType(element), element)
	}
}
