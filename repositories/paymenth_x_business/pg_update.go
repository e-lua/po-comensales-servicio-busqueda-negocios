package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Update(input_mqtt_paymenth models.Mqtt_PaymentMethod) error {

	db := models.Conectar_Pg_DB()

	query := `INSERT INTO Business_R_Paymenth(idbusiness,idPayment,isavailable) (select * from unnest($1::int[], $2::int[],$3::boolean[]))`
	if _, err := db.Exec(context.Background(), query, input_mqtt_paymenth.Idbusiness_pg, input_mqtt_paymenth.Idpaymenth_pg, input_mqtt_paymenth.Isavailable_pg); err != nil {
		return err
	}

	defer db.Close()
	return nil
}
