package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Find_Filter(idcountry int) ([]models.Pg_R_TypeFood, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()
	q := "SELECT rt.idtypefood,rt.name,rt.urlphoto FROM r_countryr_typefood AS rct JOIN r_typefood AS rt ON rct.idtypefood=rt.idtypefood WHERE rct.idcountry=$1"
	rows, error_show := db.Query(ctx, q, idcountry)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListPg_TypeFood []models.Pg_R_TypeFood

	if error_show != nil {
		return oListPg_TypeFood, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var typefoods models.Pg_R_TypeFood
		rows.Scan(&typefoods.IDTypefood, &typefoods.Name, &typefoods.Url)
		oListPg_TypeFood = append(oListPg_TypeFood, typefoods)
	}

	//Si todo esta bien
	return oListPg_TypeFood, nil

}
