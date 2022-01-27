package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Find_Uniquename(uniquename_string string) (string, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var var_string string

	db := models.Conectar_Pg_DB()

	q := "SELECT uniquename FROM business WHERE uniquename=$1"
	error_show := db.QueryRow(ctx, q, uniquename_string).Scan(&var_string)

	if error_show != nil {
		return var_string, error_show
	}

	//Si todo esta bien
	return var_string, nil
}
