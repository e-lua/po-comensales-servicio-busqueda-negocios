package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_UpdateAddress(urlphoto string, idbusiness int) error {

	db := models.Conectar_Pg_DB()

	query := `UPDATE Business SET urlbanner=$1 WHERE idbusiness=$2`
	if _, err_update := db.Exec(context.Background(), query, urlphoto, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
