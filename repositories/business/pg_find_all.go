package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Find_All() ([]interface{}, error) {

	db := models.Conectar_Pg_DB()
	q := "select json_build_object('businessdata',(b.idbusiness),'paymentmethods',json_agg(bp.idpayment),'typefoods',json_agg(bt.idtypefood),'services',json_agg(bs.idbusiness)) From business b LEFT JOIN businessr_typefood bt ON b.idbusiness=bt.idbusiness LEFT JOIN business_r_paymenth bp ON b.idbusiness=bp.idbusiness LEFT JOIN bussinessr_service bs ON b.idbusiness=bs.idbusiness GROUP BY b.idbusiness"
	rows, error_show := db.Query(context.Background(), q)
	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListaInterface []interface{}

	if error_show != nil {
		return oListaInterface, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var interfac interface{}
		rows.Scan(&interfac)
		oListaInterface = append(oListaInterface, interfac)
	}

	//Si todo esta bien
	return oListaInterface, nil
}
