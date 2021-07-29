package config

var ENV = &Env{}

var POSTGRES = &PostgresServer{}

func InitConfig() {
	ENV.Initenv()

	POSTGRES.ConnectToDB(*ENV)
}
