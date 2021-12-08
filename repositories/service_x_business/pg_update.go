package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Update(input_mqtt_service models.Mqtt_Service) error {

	db := models.Conectar_Pg_DB()

	query := `INSERT INTO bussinessr_service(idbusiness,idservice,isavailable) (select * from unnest($1::int[], $2::int[],$3::boolean[]))`
	if _, err := db.Exec(context.Background(), query, input_mqtt_service.Idbusiness_pg, input_mqtt_service.Idservice_pg, input_mqtt_service.Isavailable_pg); err != nil {
		return err
	}

	return nil
}
