package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

//En caso de hackeo
func Pg_Comensal_Recover_One(idbusiness int) (models.Mo_Business, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var business models.Mo_Business

	db := models.Conectar_Pg_DB_Comensal()
	q := "SELECT idbusiness,name,createddate,timezone,view,uniquename FROM business WHERE idbusiness=$1"
	error_show := db.QueryRow(ctx, q, idbusiness).Scan(&business.IdBusiness, &business.Name, &business.CreatedDate, &business.TimeZone, &business.View, &business.Uniquename)

	if error_show != nil {
		return business, error_show
	}

	//Si todo esta bien
	return business, nil
}
