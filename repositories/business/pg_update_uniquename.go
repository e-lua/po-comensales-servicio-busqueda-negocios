package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_UpdateUniqueName(input_mqtt_uniquename models.Mqtt_Uniquename) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Cambio de Server y BD, ya que no se puede acceder al rol de superusuario para la busqueda por distancia
	db := models.Conectar_Pg_DB_Comensal()

	query := `UPDATE Business SET uniquename=$1 WHERE idbusiness=$2`
	if _, err_update := db.Exec(ctx, query, input_mqtt_uniquename.Uniquename, input_mqtt_uniquename.IdBusiness); err_update != nil {
		return err_update
	}

	return nil
}
