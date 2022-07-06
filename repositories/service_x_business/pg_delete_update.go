package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Delete_Update(input_mqtt_service models.Mqtt_Service) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Conexion con la BD
	//Cambio de Server y BD, ya que no se puede acceder al rol de superusuario para la busqueda por distancia
	db := models.Conectar_Pg_DB_Comensal()

	//BEGIN
	tx, error_tx := db.Begin(ctx)
	if error_tx != nil {
		tx.Rollback(ctx)
		return error_tx
	}

	//ELIMINAR DATOS
	q_delete := `DELETE FROM bussinessr_service WHERE idbusiness=$1`
	_, err := tx.Exec(ctx, q_delete, input_mqtt_service.IdBusiness)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	//HORARIO
	q_schedulerange := `INSERT INTO bussinessr_service(idbusiness,idservice,pricing,typemoney,isavailable) (select * from unnest($1::int[], $2::int[],$3::decimal(8,2)[],$4::int[],$5::boolean[]))`
	if _, err_schedule := tx.Exec(ctx, q_schedulerange, input_mqtt_service.Idbusiness_pg, input_mqtt_service.Idservice_pg, input_mqtt_service.Pricing_pg, input_mqtt_service.TypeMoney_pg, input_mqtt_service.Isavailable_pg); err_schedule != nil {
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
