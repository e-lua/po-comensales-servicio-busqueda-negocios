package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Find_Uniquename(uniquename_string string) (string, error) {

	var var_string string

	db := models.Conectar_Pg_DB()

	q := "SELECT uniquename FROM business WHERE uniquename=$1"
	error_show := db.QueryRow(context.Background(), q, uniquename_string).Scan(&var_string)

	if error_show != nil {
		return var_string, error_show
	}

	//Si todo esta bien
	return var_string, nil
}
