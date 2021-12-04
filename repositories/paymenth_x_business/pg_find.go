package repositories

import (
	models "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/models"
)

func Pg_Find(idbusiness int) ([]models.Pg_PaymenthMethXBusiness, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT r_paymentmethod.idpayment,r_paymentmethod.name,r_paymentmethod.hasnumber,paymentmethxbusiness.phone,r_paymentmethod.urlphoto,paymentmethxbusiness.isavailable FROM paymentmethxbusiness LEFT JOIN r_paymentmethod ON paymentmethxbusiness.idpayment = r_paymentmethod.idpayment WHERE paymentmethxbusiness.idbusiness=$1"
	rows, error_shown := db.Query(q, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	oListPg_PaymenthMethXBusiness := []models.Pg_PaymenthMethXBusiness{}

	if error_shown != nil {
		defer db.Close()
		return oListPg_PaymenthMethXBusiness, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		payment_x_Business := models.Pg_PaymenthMethXBusiness{}
		rows.Scan(&payment_x_Business.IDPaymenth, &payment_x_Business.Name, &payment_x_Business.HasNumber, &payment_x_Business.PhoneNumber, &payment_x_Business.Url, &payment_x_Business.IsAvaiable)
		oListPg_PaymenthMethXBusiness = append(oListPg_PaymenthMethXBusiness, payment_x_Business)
	}

	defer db.Close()

	//Si todo esta bien
	return oListPg_PaymenthMethXBusiness, nil

}
