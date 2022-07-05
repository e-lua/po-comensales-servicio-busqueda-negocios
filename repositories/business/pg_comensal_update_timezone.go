package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Comensal_UpdateTimeZone(input_mqtt_timezone models.Mqtt_TimeZone) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB_Comensal()

	query := `UPDATE Business SET timezone=$1 WHERE idbusiness=$2`
	if _, err_update := db.Exec(ctx, query, input_mqtt_timezone.TimeZone, input_mqtt_timezone.IdBusiness); err_update != nil {
		return err_update
	}

	return nil
}
