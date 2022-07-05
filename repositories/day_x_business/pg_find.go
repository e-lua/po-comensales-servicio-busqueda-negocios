package repositories

import (
	"context"
	"math/rand"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Pg_Find(idbusiness int) ([]models.Pg_R_Schedule_ToBusiness, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var db *pgxpool.Pool

	random := rand.Intn(4)
	if random%2 == 0 {
		db = models.Conectar_Pg_DB()
	} else {
		db = models.Conectar_Pg_DB_Slave()
	}

	q := "SELECT r.idschedule,bsch.starttime,bsch.endtime,bsch.isavailable FROM schedule AS r LEFT JOIN businessschedule AS bsch ON bsch.idschedule=r.idschedule WHERE bsch.idbusiness=$1 UNION SELECT r.idschedule,'0','0',false FROM schedule AS r LEFT JOIN businessschedule AS bsch ON bsch.idschedule=r.idschedule WHERE r.idschedule NOT IN (SELECT bsch.idschedule FROM businessschedule AS bsch WHERE bsch.idbusiness=$1)"
	rows, error_show := db.Query(ctx, q, idbusiness)

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
