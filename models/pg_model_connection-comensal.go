package models

import (
	"context"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var PostgresCN_Comensal = Conectar_Pg_DB_Comensal()

var (
	once_pg_Comensal sync.Once
	p_pg_Comensal    *pgxpool.Pool
)

func Conectar_Pg_DB_Comensal() *pgxpool.Pool {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	once_pg_Comensal.Do(func() {
		urlString := "postgres://postgresxd00:GFgfk45345GGHdfinhjti5BHerYTu7fsdggn000@postgres-comensal:5432/postgresxd00?pool_max_conns=150"
		config, _ := pgxpool.ParseConfig(urlString)
		p_pg_Comensal, _ = pgxpool.ConnectConfig(ctx, config)
	})
	return p_pg_Comensal
}
