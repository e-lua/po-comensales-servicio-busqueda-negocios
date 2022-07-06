package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Find_Address(idbusiness int) (float64, float64, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var latitude float64
	var longitude float64

	//Cambio de Server y BD, ya que no se puede acceder al rol de superusuario para la busqueda por distancia

	db := models.Conectar_Pg_DB_Comensal()

	q := "SELECT latitude,longitude FROM business where idbusiness=$1"
	error_shown := db.QueryRow(ctx, q, idbusiness).Scan(&latitude, &longitude)

	if error_shown != nil {
		return latitude, longitude, error_shown
	}

	//Si todo esta bien
	return latitude, longitude, nil
}
