package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

//En caso de hackeo
func Pg_Recover_All() ([]models.Mo_Business, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT idbusiness,name,createddate,timezone,view,uniquename FROM business"
	rows, error_show := db.Query(context.Background(), q)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListBusiness []models.Mo_Business

	if error_show != nil {
		return oListBusiness, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var business models.Mo_Business
		rows.Scan(&business.IdBusiness, &business.Name, &business.CreatedDate, &business.TimeZone, &business.View, &business.Uniquename)
		oListBusiness = append(oListBusiness, business)
	}

	//Si todo esta bien
	return oListBusiness, nil
}
