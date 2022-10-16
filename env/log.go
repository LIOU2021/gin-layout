package env

type Log struct {
	FileName string
}

var LogSetting = &Log{}

func init() {
	EnvSlice = append(EnvSlice, LogSetting)
}
