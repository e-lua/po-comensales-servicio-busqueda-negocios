package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_UpdateBanner(input_mqtt_banner models.Mqtt_Banner_Cola) error {

	db := models.Conectar_Pg_DB()

	query := `UPDATE Business SET urlbanner=$1 WHERE idbusiness=$2`
	if _, err_update := db.Exec(context.Background(), query, input_mqtt_banner.Url, input_mqtt_banner.IdBusiness); err_update != nil {
		return err_update
	}
	defer db.Close()
	return nil
}
