package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_UpdateAddress(input_mqtt_address models.Mqtt_Addres) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Cambio de Server y BD, ya que no se puede acceder al rol de superusuario para la busqueda por distancia
	db := models.Conectar_Pg_DB_Comensal()

	query := `UPDATE business SET latitude=$1,longitude=$2 WHERE idbusiness=$3`
	if _, err := db.Exec(ctx, query, input_mqtt_address.Latitude, input_mqtt_address.Longitude, input_mqtt_address.IdBusiness); err != nil {
		return err
	}

	return nil
}
