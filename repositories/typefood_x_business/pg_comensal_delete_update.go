package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Comensal_Delete_Update(input_mqtt_typefood models.Mqtt_TypeFood) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Conexion con la BD
	db := models.Conectar_Pg_DB_Comensal()

	//BEGIN
	tx, error_tx := db.Begin(ctx)
	if error_tx != nil {
		tx.Rollback(ctx)
		return error_tx
	}

	//ELIMINAR DATOS
	q_delete := `DELETE FROM businessr_typefood WHERE idbusiness=$1`
	_, err := tx.Exec(ctx, q_delete, input_mqtt_typefood.IdBusiness)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	//HORARIO
	q_schedulerange := `INSERT INTO businessr_typefood(idbusiness,idtypefood,isavailable) (select * from unnest($1::int[], $2::int[],$3::boolean[]))`
	if _, err_schedule := tx.Exec(ctx, q_schedulerange, input_mqtt_typefood.Idbusiness_pg, input_mqtt_typefood.Idtypefood_pg, input_mqtt_typefood.Isavailable_pg); err_schedule != nil {
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
