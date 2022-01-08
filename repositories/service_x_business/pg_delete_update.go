package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Delete_Update(input_mqtt_service models.Mqtt_Service) error {

	//Conexion con la BD
	db := models.Conectar_Pg_DB()

	//BEGIN
	tx, error_tx := db.Begin(context.Background())
	if error_tx != nil {
		tx.Rollback(context.Background())
		return error_tx
	}

	//ELIMINAR DATOS
	q_delete := `DELETE FROM bussinessr_service WHERE idbusiness=$1`
	_, err := tx.Exec(context.Background(), q_delete, input_mqtt_service.IdBusiness)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	//HORARIO
	q_schedulerange := `INSERT INTO bussinessr_service(idbusiness,idservice,pricing,typemoney,isavailable) (select * from unnest($1::int[], $2::int[],$3::decimal(8,2)[],$4::int[],$5::boolean[]))`
	if _, err_schedule := tx.Exec(context.Background(), q_schedulerange, input_mqtt_service.Idbusiness_pg, input_mqtt_service.Idservice_pg, input_mqtt_service.Pricing_pg, input_mqtt_service.TypeMoney_pg, input_mqtt_service.Isavailable_pg); err_schedule != nil {
		tx.Rollback(context.Background())
		return err_schedule
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(context.Background())
	if err_commit != nil {
		tx.Rollback(context.Background())
		return err_commit
	}

	return nil
}
