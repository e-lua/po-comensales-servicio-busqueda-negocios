package repositories

import (
	models "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/models"
)

func Pg_Find(idbusiness int) ([]models.Pg_TypeFoodXBusiness, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT r_typefood.idtypefood,r_typefood.name,r_typefood.urlphoto,typefoodxbusiness.isavailable FROM typefoodxbusiness LEFT JOIN r_typefood ON typefoodxbusiness.idtypefood = r_typefood.idtypefood WHERE typefoodxbusiness.idbusiness=$1"
	rows, error_showname := db.Query(q, idbusiness)

	//.Scan(&typeF_x_Business.IDTypeFood, &typeF_x_Business.NameFood, &typeF_x_Business.URLPhoto, &typeF_x_Business.Weight)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	oListPg_TypeFoodXBusiness := []models.Pg_TypeFoodXBusiness{}

	if error_showname != nil {

		defer db.Close()
		return oListPg_TypeFoodXBusiness, error_showname
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		typeF_x_Business := models.Pg_TypeFoodXBusiness{}
		rows.Scan(&typeF_x_Business.IDTypeFood, &typeF_x_Business.NameFood, &typeF_x_Business.URL, &typeF_x_Business.IsAvaiable)
		oListPg_TypeFoodXBusiness = append(oListPg_TypeFoodXBusiness, typeF_x_Business)
	}

	defer db.Close()

	//Si todo esta bien
	return oListPg_TypeFoodXBusiness, nil

}
