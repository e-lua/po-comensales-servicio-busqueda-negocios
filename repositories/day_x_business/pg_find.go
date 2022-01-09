package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Find(idbusiness int) ([]models.Pg_R_Schedule_ToBusiness, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT r.idschedule,bsch.starttime,bsch.endtime,bsch.isavailable FROM schedule AS r LEFT JOIN businessschedule AS bsch ON bsch.idschedule=r.idschedule WHERE bsch.idbusiness=$1 UNION SELECT r.idschedule,'0','0',false FROM schedule AS r LEFT JOIN businessschedule AS bsch ON bsch.idschedule=r.idschedule WHERE r.idschedule NOT IN (SELECT bsch.idschedule FROM businessschedule AS bsch WHERE bsch.idbusiness=$1)"
	rows, error_show := db.Query(context.Background(), q, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListPg_Schedule []models.Pg_R_Schedule_ToBusiness

	if error_show != nil {
		return oListPg_Schedule, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var schedule models.Pg_R_Schedule_ToBusiness
		rows.Scan(&schedule.IDSchedule, &schedule.Starttime, &schedule.Endtime, &schedule.Available)
		oListPg_Schedule = append(oListPg_Schedule, schedule)
	}

	//Si todo esta bien
	return oListPg_Schedule, nil

}
