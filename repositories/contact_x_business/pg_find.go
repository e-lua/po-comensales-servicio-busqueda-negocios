package repositories

import (
	models "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/models"
)

func Pg_Find(idbusiness int) ([]models.Pg_ContactxBusiness, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT contactxbusiness.idcontact,r_contact.name,contactxbusiness.description,contactxbusiness.isavailable FROM contactxbusiness LEFT JOIN r_contact ON contactxbusiness.idcontact = r_contact.idcontact WHERE contactxbusiness.idbusiness=$1"
	rows, error_shown := db.Query(q, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	oListPg_ContactxBusiness := []models.Pg_ContactxBusiness{}

	if error_shown != nil {
		defer db.Close()
		return oListPg_ContactxBusiness, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		contact_x_Business := models.Pg_ContactxBusiness{}
		rows.Scan(&contact_x_Business.IDContact, &contact_x_Business.Name, &contact_x_Business.Description, &contact_x_Business.IsAvaiable)
		oListPg_ContactxBusiness = append(oListPg_ContactxBusiness, contact_x_Business)
	}

	defer db.Close()

	//Si todo esta bien
	return oListPg_ContactxBusiness, nil

}
