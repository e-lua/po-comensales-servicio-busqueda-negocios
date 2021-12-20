package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Find_Filter(idcountry int) ([]models.Pg_R_PaymentMethod, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT rp.idpayment,rp.name,rp.urlphoto FROM r_countryr_payment AS rcp JOIN r_paymentmethod AS rp ON rcp.idpayment=rp.idpayment WHERE rcp.idcountry=$1"
	rows, error_show := db.Query(context.Background(), q, idcountry)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListPg_Paymenth []models.Pg_R_PaymentMethod

	if error_show != nil {
		return oListPg_Paymenth, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var paymenth models.Pg_R_PaymentMethod
		rows.Scan(&paymenth.IDPaymenth, &paymenth.Name, &paymenth.Url)
		oListPg_Paymenth = append(oListPg_Paymenth, paymenth)
	}

	//Si todo esta bien
	return oListPg_Paymenth, nil

}
