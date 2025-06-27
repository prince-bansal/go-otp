package config

type DbConfig struct {
	Host     string
	Database string
	User     string
	Password string
}

type Config struct {
	Db DbConfig
}
