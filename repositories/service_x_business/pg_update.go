package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Update(input_mqtt_service models.Mqtt_Service) error {

	db := models.Conectar_Pg_DB()

	query := `INSERT INTO bussinessr_service(idbusiness,idservice,pricing,typemoney,isavailable) (select * from unnest($1::int[], $2::int[],$4::decimal(8,2)[],$4::int[],$5::boolean[]))`
	if _, err := db.Exec(context.Background(), query, input_mqtt_service.Idbusiness_pg, input_mqtt_service.Idservice_pg, input_mqtt_service.Pricing_pg, input_mqtt_service.TypeMoney_pg, input_mqtt_service.Isavailable_pg); err != nil {
		return err
	}

	return nil
}
