package models

type Config struct {
	Database  DatabaseConfig
	TimeZone  string
	SecretKey string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	DbName   string
	Username string
	Password string
}
