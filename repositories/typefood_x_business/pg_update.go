package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Update(input_mqtt_typefood models.Mqtt_TypeFood) error {

	db := models.Conectar_Pg_DB()

	query := `INSERT INTO businessr_typefood(idbusiness,idtypefood,isavailable) (select * from unnest($1::int[], $2::int[],$3::boolean[]))`
	if _, err := db.Exec(context.Background(), query, input_mqtt_typefood.Idbusiness_pg, input_mqtt_typefood.Idtypefood_pg, input_mqtt_typefood.Isavailable_pg); err != nil {
		return err
	}

	defer db.Close()
	return nil
}
