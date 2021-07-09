package config

import (
	"context"
	"strconv"

	pg "github.com/go-pg/pg/v10"
)

type PostgresServer struct {
	DB *pg.DB
}

func (ps *PostgresServer) ConnectToDB(env Env) {
	ps.DB = pg.Connect(&pg.Options{
		Addr:     env.PostgresHost + ":" + strconv.Itoa(env.PostgresPort),
		User:     env.PostgresUser,
		Password: env.PostgresPassword,
		Database: env.PostgresDatabase,
	})

	ctx := context.Background()

	if err := ps.DB.Ping(ctx); err != nil {
		panic(err)
	}
}
