package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Delete(idbusiness int) error {

	db := models.Conectar_Pg_DB()

	//Eliminamos los datos
	q := `DELETE FROM bussinessr_service WHERE idbusiness=$1`
	_, err := db.Exec(context.Background(), q, idbusiness)
	if err != nil {
		return err
	}

	defer db.Close()
	return nil
}
