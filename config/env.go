package config

import (
	"os"
	"strconv"
)

type Env struct {
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	ApiPort          int
}

func (env *Env) Initenv() {
	env.PostgresHost = os.Getenv("POSTGRES_HOST")
	env.PostgresUser = os.Getenv("POSTGRES_USER")
	env.PostgresPort, _ = strconv.Atoi(os.Getenv("POSTGRES_POST"))
	env.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	env.PostgresDatabase = os.Getenv("POSTGRES_DATABASE")
	env.ApiPort, _ = strconv.Atoi(os.Getenv("API_PORT"))
}
