package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_UpdateIsOpen(input_mqtt_open models.Mqtt_IsOpen) error {

	db := models.Conectar_Pg_DB()

	query := `UPDATE Business SET isOpen=$1 WHERE idbusiness=$2`
	if _, err_update := db.Exec(context.Background(), query, input_mqtt_open.IsaOpen, input_mqtt_open.IdBusiness); err_update != nil {
		return err_update
	}
	defer db.Close()
	return nil
}
