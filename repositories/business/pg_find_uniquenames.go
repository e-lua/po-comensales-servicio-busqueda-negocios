package repositories

import (
	"context"
	"math/rand"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Pg_Find_Uniquename(uniquename_string string) (string, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var var_string string

	var db *pgxpool.Pool

	random := rand.Intn(4)
	if random%2 == 0 {
		db = models.Conectar_Pg_DB()
	} else {
		db = models.Conectar_Pg_DB_Slave()
	}

	q := "SELECT uniquename FROM business WHERE uniquename=$1"
	error_show := db.QueryRow(ctx, q, uniquename_string).Scan(&var_string)

	if error_show != nil {
		return var_string, error_show
	}

	//Si todo esta bien
	return var_string, nil
}
