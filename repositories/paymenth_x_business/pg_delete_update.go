package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Delete_Update(input_mqtt_paymenth models.Mqtt_PaymentMethod) error {

	//Conexion con la BD
	db := models.Conectar_Pg_DB()

	//BEGIN
	tx, error_tx := db.Begin(context.Background())
	if error_tx != nil {
		tx.Rollback(context.Background())
		return error_tx
	}

	//ELIMINAR DATOS
	q_delete := `DELETE FROM Business_R_Paymenth WHERE idbusiness=$1`
	_, err := tx.Exec(context.Background(), q_delete, input_mqtt_paymenth.IdBusiness)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	//PAYMENTH
	q_schedulerange := `INSERT INTO Business_R_Paymenth(idbusiness,idPayment,isavailable,phonenumber) (select * from unnest($1::int[], $2::int[],$3::boolean[],$4::varchar(25)[]))`
	if _, err_schedule := tx.Exec(context.Background(), q_schedulerange, input_mqtt_paymenth.Idbusiness_pg, input_mqtt_paymenth.Idpaymenth_pg, input_mqtt_paymenth.Isavailable_pg, input_mqtt_paymenth.PhoneNumber); err_schedule != nil {
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
