package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Delete_Update(input_mqtt_schedule models.Mqtt_Schedule) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Conexion con la BD
	db := models.Conectar_Pg_DB()

	//BEGIN
	tx, error_tx := db.Begin(ctx)
	if error_tx != nil {
		tx.Rollback(ctx)
		return error_tx
	}

	//ELIMINAR DATOS
	q_delete := `DELETE FROM BusinessSchedule WHERE idbusiness=$1`
	_, err := tx.Exec(ctx, q_delete, input_mqtt_schedule.IdBusiness)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	//HORARIO
	q_schedulerange := `INSERT INTO BusinessSchedule(idschedule,idbusiness,starttime,endtime,isavailable) (SELECT * FROM unnest($1::int[],$2::int[],$3::varchar(14)[],$4::varchar(14)[],$5::boolean[]));`
	if _, err_schedule := tx.Exec(ctx, q_schedulerange, input_mqtt_schedule.Idschedule_pg, input_mqtt_schedule.Idbusiness_pg, input_mqtt_schedule.Starttime_pg, input_mqtt_schedule.Endtime_pg, input_mqtt_schedule.Isavailable_pg); err_schedule != nil {
		tx.Rollback(ctx)
		return err_schedule
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(ctx)
	if err_commit != nil {
		tx.Rollback(ctx)
		return err_commit
	}

	return nil
}
