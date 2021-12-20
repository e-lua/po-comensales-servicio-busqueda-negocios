package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Find_Near_Searched(idcomensal int) (int, []models.Pg_Found_All_Business, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT nearbusiness FROM Near WHERE idcomensal=$1"
	rows, error_show := db.Query(context.Background(), q, idcomensal)

	rowsaffect := 0
	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListBusiness []models.Pg_Found_All_Business

	if error_show != nil {
		return rowsaffect, oListBusiness, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var interfac models.Pg_Found_All_Business
		rows.Scan(&interfac)
		rowsaffect = rowsaffect + 1
		oListBusiness = append(oListBusiness, interfac)
	}

	//Si todo esta bien
	return rowsaffect, oListBusiness, nil

}
