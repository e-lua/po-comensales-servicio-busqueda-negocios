package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Delete_Update(input_mqtt_paymenth models.Mqtt_PaymentMethod) error {

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
	q_delete := `DELETE FROM Business_R_Paymenth WHERE idbusiness=$1`
	_, err := tx.Exec(ctx, q_delete, input_mqtt_paymenth.IdBusiness)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	//PAYMENTH
	q_schedulerange := `INSERT INTO Business_R_Paymenth(idbusiness,idPayment,isavailable,phonenumber) (select * from unnest($1::int[], $2::int[],$3::boolean[],$4::varchar(25)[]))`
	if _, err_schedule := tx.Exec(ctx, q_schedulerange, input_mqtt_paymenth.Idbusiness_pg, input_mqtt_paymenth.Idpaymenth_pg, input_mqtt_paymenth.Isavailable_pg, input_mqtt_paymenth.PhoneNumber); err_schedule != nil {
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
