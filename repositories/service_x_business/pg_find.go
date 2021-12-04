package repositories

import (
	models "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/models"
)

func Pg_Find(idbusiness int) ([]models.Pg_ServiceXBusiness, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT r_service.idservice,r_service.name,servicexbusiness.price,r_service.urlphoto,servicexbusiness.typemoney,servicexbusiness.isavailable FROM servicexbusiness LEFT JOIN r_service ON servicexbusiness.idservice = r_service.idservice WHERE servicexbusiness.idbusiness=$1"
	rows, error_shown := db.Query(q, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	oListPg_ServiceXBusiness := []models.Pg_ServiceXBusiness{}

	if error_shown != nil {

		defer db.Close()
		return oListPg_ServiceXBusiness, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		service_x_Business := models.Pg_ServiceXBusiness{}
		rows.Scan(&service_x_Business.IDService, &service_x_Business.Name, &service_x_Business.Price, &service_x_Business.Url, &service_x_Business.TypeMoney, &service_x_Business.IsAvaiable)
		oListPg_ServiceXBusiness = append(oListPg_ServiceXBusiness, service_x_Business)
	}

	defer db.Close()

	//Si todo esta bien
	return oListPg_ServiceXBusiness, nil

}
