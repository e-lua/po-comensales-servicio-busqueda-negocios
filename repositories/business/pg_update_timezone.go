package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_UpdateTimeZone(input_mqtt_timezone models.Mqtt_TimeZone) error {

	db := models.Conectar_Pg_DB()

	query := `UPDATE Business SET timezone=$1 WHERE idbusiness=$2`
	if _, err_update := db.Exec(context.Background(), query, input_mqtt_timezone.TimeZone, input_mqtt_timezone.IdBusiness); err_update != nil {
		return err_update
	}

	return nil
}
