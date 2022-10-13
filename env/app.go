package env

type App struct {
	HOST     string
	PORT     string
	TimeZone string
}

var AppSetting = &App{}
