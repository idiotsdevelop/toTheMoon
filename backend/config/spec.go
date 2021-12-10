package config

type ENV string

const (
	Local ENV = "local"
	Dev  ENV = "dev"
	Prod ENV = "prod"
)

const (
	DBWHost    = "MYSQL.WHOST"
	DBRHost    = "MYSQL.RHOST"
	DBDatabase = "MYSQL.DATABASE"
	DBUser     = "MYSQL.USER"
)
