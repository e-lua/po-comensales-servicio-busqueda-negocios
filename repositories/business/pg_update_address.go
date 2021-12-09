package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_UpdateAddress(input_mqtt_address models.Mqtt_Addres) error {

	db := models.Conectar_Pg_DB()

	query := `UPDATE business SET latitude=$1,longitude=$2 WHERE idbusiness=$3`
	if _, err := db.Exec(context.Background(), query, input_mqtt_address.Latitude, input_mqtt_address.Longitude, input_mqtt_address.IdBusiness); err != nil {
		return err
	}

	return nil
}
