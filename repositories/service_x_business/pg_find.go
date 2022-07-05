package repositories

import (
	"context"
	"math/rand"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Pg_Find(idbusiness int, idcountry int) ([]models.Pg_R_Service_ToBusiness, error) {

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

	q := "SELECT r.idservice,r.name,bs.pricing,bs.typemoney,bs.isavailable FROM r_service AS r LEFT JOIN bussinessr_service AS bs ON bs.idservice=r.idservice WHERE bs.idbusiness=$1 UNION SELECT r.idservice,r.name,0,0,false FROM r_service AS r LEFT JOIN bussinessr_service AS bs ON bs.idservice=r.idservice LEFT JOIN r_countryr_service AS rr ON rr.idservice=r.idservice WHERE r.idservice NOT IN (SELECT bs.idservice FROM bussinessr_service AS bs WHERE bs.idbusiness=$1)AND rr.idcountry=$2"
	rows, error_show := db.Query(ctx, q, idbusiness, idcountry)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListPg_Service []models.Pg_R_Service_ToBusiness

	if error_show != nil {
		return oListPg_Service, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var service models.Pg_R_Service_ToBusiness
		rows.Scan(&service.IDservice, &service.Name, &service.Pricing, &service.TypeMoney, &service.IsAvailable)
		oListPg_Service = append(oListPg_Service, service)
	}

	//Si todo esta bien
	return oListPg_Service, nil

}
