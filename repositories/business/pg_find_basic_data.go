package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Find_BasicData(idbusiness int) (models.Pg_BasicData_ToBusiness, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var basic_data models.Pg_BasicData_ToBusiness

	db := models.Conectar_Pg_DB()
	q := "SELECT b.typesuscription,COALESCE(b.uniquename,'sin nombre'),COALESCE(b.name,'sin nombre'), COALESCE(b.timezone,'0'),COALESCE(CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone THEN true ELSE false END,'false') FROM business as b JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND b.idbusiness=$1"
	error_shown := db.QueryRow(ctx, q, idbusiness).Scan(&basic_data.Typesuscription, &basic_data.Uniquename, &basic_data.Name, &basic_data.TimeZone, &basic_data.IsOpen)

	if error_shown != nil {
		return basic_data, error_shown
	}

	//Si todo esta bien
	return basic_data, nil
}

func Pg_Find_BasicData_WithoutData(idbusiness int) (models.Pg_BasicData_ToBusiness, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var basic_data models.Pg_BasicData_ToBusiness

	db := models.Conectar_Pg_DB()
	q := "SELECT b.typesuscription,COALESCE(b.uniquename,'sin nombre'),COALESCE(b.name,'sin nombre'),COALESCE(b.timezone,'0'),false FROM business as b WHERE b.idbusiness=$1"
	error_shown := db.QueryRow(ctx, q, idbusiness).Scan(&basic_data.Typesuscription, &basic_data.Uniquename, &basic_data.Name, &basic_data.TimeZone, &basic_data.IsOpen)

	if error_shown != nil {
		return basic_data, error_shown
	}

	//Si todo esta bien
	return basic_data, nil
}
