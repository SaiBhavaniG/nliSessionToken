package models

type Config struct {
	AppID        string `env:"APP_ID" validate:"required"`
	AeroHostName string `env:"AERO_HOST" validate:"required"`
	AeroPort     int    `env:"AERO_PORT" validate:"required"`
	HostAddress  string `env:"HOST_ADDR,localhost:8080"`
	BaseClient   string `env:"BASE_CLIENT,gcp"`
	DaoClient    string `env:"DAO_CLIENT,aerospike"`
}
