package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

//En caso de hackeo
func Pg_Recover_One(idbusiness int) (models.Mo_Business, error) {

	var business models.Mo_Business

	db := models.Conectar_Pg_DB()
	q := "SELECT idbusiness,name,createddate,timezone,view,uniquename FROM business WHERE idbusiness=$1"
	error_show := db.QueryRow(context.Background(), q, idbusiness).Scan(&business.IdBusiness, &business.Name, &business.CreatedDate, &business.TimeZone, &business.View, &business.Uniquename)

	if error_show != nil {
		return business, error_show
	}

	//Si todo esta bien
	return business, nil
}
