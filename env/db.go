package env

type DB struct {
	UserName     string
	Password     string
	Addr         string
	Port         int
	Database     string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

var DbSetting = &DB{}

func init() {
	EnvSlice = append(EnvSlice, DbSetting)
}
