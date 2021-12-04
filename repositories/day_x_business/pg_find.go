package repositories

import (
	models "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/models"
)

func Pg_Find(idbusiness int) ([]models.Pg_DayXBusiness, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT dayxbusiness.idday,r_day.name,dayxbusiness.starttime,dayxbusiness.endtime,dayxbusiness.isavailable FROM dayxbusiness LEFT JOIN r_day ON dayxbusiness.idday = r_day.idday WHERE dayxbusiness.idbusiness=$1"
	rows, error_shown := db.Query(q, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	oListPg_DayXBusiness := []models.Pg_DayXBusiness{}

	if error_shown != nil {
		defer db.Close()
		return oListPg_DayXBusiness, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		day_x_Business := models.Pg_DayXBusiness{}
		rows.Scan(&day_x_Business.IDDia, &day_x_Business.Name, &day_x_Business.StarTime, &day_x_Business.EndTime, &day_x_Business.IsAvaiable)
		oListPg_DayXBusiness = append(oListPg_DayXBusiness, day_x_Business)
	}

	defer db.Close()

	//Si todo esta bien
	return oListPg_DayXBusiness, nil

}
