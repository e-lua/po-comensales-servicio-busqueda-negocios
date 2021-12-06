package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_UpdateIsOpen(isOpen bool, idbusiness int) error {

	db := models.Conectar_Pg_DB()

	query := `UPDATE Business SET isOpen=$1 WHERE idbusiness=$2`
	if _, err_update := db.Exec(context.Background(), query, isOpen, idbusiness); err_update != nil {
		return err_update
	}
	defer db.Close()
	return nil
}
