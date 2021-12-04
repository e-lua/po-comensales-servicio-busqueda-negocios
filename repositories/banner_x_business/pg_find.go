package repositories

import (
	models "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/models"
)

func Pg_Find(idbusiness int) ([]models.Pg_BannerXBusiness, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT BannerxBusiness.idbanner,BannerxBusiness.urlphoto FROM BannerxBusiness JOIN r_banner ON BannerxBusiness.idbanner = r_banner.idbanner WHERE BannerxBusiness.idbusiness=$1"
	rows, error_banner := db.Query(q, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	oListPg_BannerXBusiness := []models.Pg_BannerXBusiness{}

	if error_banner != nil {
		defer db.Close()
		return oListPg_BannerXBusiness, error_banner
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		banner_x_Business := models.Pg_BannerXBusiness{}
		rows.Scan(&banner_x_Business.IDBanner, &banner_x_Business.URL)
		oListPg_BannerXBusiness = append(oListPg_BannerXBusiness, banner_x_Business)
	}

	//Si todo esta bien
	return oListPg_BannerXBusiness, nil
}
