package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Find_All() ([]models.Json_Postgresql, error) {

	db := models.Conectar_Pg_DB()
	q := "select json_build_object('businessdata',(b.idbusiness),'paymentmethods',json_agg(bp.idpayment),'typefoods',json_agg(bt.idtypefood),'services',json_agg(bs.idbusiness)) From business b LEFT JOIN businessr_typefood bt ON b.idbusiness=bt.idbusiness LEFT JOIN business_r_paymenth bp ON b.idbusiness=bp.idbusiness LEFT JOIN bussinessr_service bs ON b.idbusiness=bs.idbusiness GROUP BY b.idbusiness"
	rows, error_show := db.Query(context.Background(), q)
	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListJsonPostgresql []models.Json_Postgresql

	if error_show != nil {
		return oListJsonPostgresql, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var jsonPostgresql models.Json_Postgresql
		rows.Scan(&jsonPostgresql.Business, &jsonPostgresql.Payments, &jsonPostgresql.TypeFoods, &jsonPostgresql.Services)
		oListJsonPostgresql = append(oListJsonPostgresql, jsonPostgresql)
	}

	//Si todo esta bien
	return oListJsonPostgresql, nil
}
