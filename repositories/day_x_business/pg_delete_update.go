package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Delete_Update(input_mqtt_schedule models.Mqtt_Schedule) error {

	//Conexion con la BD
	db := models.Conectar_Pg_DB()

	//BEGIN
	tx, error_tx := db.Begin(context.Background())
	if error_tx != nil {
		tx.Rollback(context.Background())
		return error_tx
	}

	//ELIMINAR DATOS
	q_delete := `DELETE FROM BusinessSchedule WHERE idbusiness=$1`
	_, err := tx.Exec(context.Background(), q_delete, input_mqtt_schedule.IdBusiness)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	//HORARIO
	q_schedulerange := `INSERT INTO BusinessSchedule(idschedule,idbusiness,starttime,endtime,isavailable,name) (SELECT * FROM unnest($1::int[],$2::int[],$3::varchar(14)[],$4::varchar(14)[],$5::boolean[],$6::varchar(30)[]));`
	if _, err_schedule := tx.Exec(context.Background(), q_schedulerange, input_mqtt_schedule.Idschedule_pg, input_mqtt_schedule.Idbusiness_pg, input_mqtt_schedule.Starttime_pg, input_mqtt_schedule.Endtime_pg, input_mqtt_schedule.Isavailable_pg, input_mqtt_schedule.Name_pg); err_schedule != nil {
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
