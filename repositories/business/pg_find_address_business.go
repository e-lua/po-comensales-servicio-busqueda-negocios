package repositories

import (
	"context"
	"math/rand"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Pg_Find_Address(idbusiness int) (float64, float64, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var latitude float64
	var longitude float64
	var db *pgxpool.Pool

	random := rand.Intn(4)
	if random%2 == 0 {
		db = models.Conectar_Pg_DB()
	} else {
		db = models.Conectar_Pg_DB_Slave()
	}

	q := "SELECT latitude,longitude FROM business where idbusiness=$1"
	error_shown := db.QueryRow(ctx, q, idbusiness).Scan(&latitude, &longitude)

	if error_shown != nil {
		return latitude, longitude, error_shown
	}

	//Si todo esta bien
	return latitude, longitude, nil
}
