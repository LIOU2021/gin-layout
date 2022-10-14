package env

type App struct {
	HOST     string
	PORT     string
	TimeZone string
	Key      string
}

var AppSetting = &App{}
