package config

import (
	"log"

	"github.com/LIOU2021/gin-layout/env"
	"github.com/LIOU2021/gin-layout/helpers"
	"github.com/go-ini/ini"
)

var cfg *ini.File

// Setup initialize the configuration instance
func envRegister() {
	var err error
	cfg, err = ini.Load(".env")

	if err != nil {
		log.Fatalf("setting.Setup, fail to parse '.env': %v", err)
	}

	foreachMapTo(env.EnvSlice)
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
