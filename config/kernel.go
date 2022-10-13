package config

// register all config setting
func Kernel() {
	envRegister()
	CreateLogFile(LogName())
}
