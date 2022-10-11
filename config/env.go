package config

import (
	"log"

	"github.com/go-ini/ini"
)

type Log struct {
	FileName string
}

var LogSetting = &Log{}

var cfg *ini.File

// Setup initialize the configuration instance
func env() {
	var err error
	cfg, err = ini.Load(".env")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse '.env': %v", err)
	}

	mapTo("log", LogSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
