package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Find(idbusiness int, idcountry int) ([]models.Pg_R_TypeFood_ToBusiness, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Cambio de Server y BD, ya que no se puede acceder al rol de superusuario para la busqueda por distancia

	db := models.Conectar_Pg_DB_Comensal()

	q := "SELECT r.idtypefood,r.name,r.urlphoto,bt.isavailable from r_typefood AS r LEFT JOIN businessr_typefood AS bt ON r.idtypefood=bt.idtypefood WHERE bt.idbusiness=$1 UNION SELECT r.idtypefood,r.name,r.urlphoto,false from r_typefood AS r LEFT JOIN businessr_typefood AS bt ON r.idtypefood=bt.idtypefood LEFT JOIN r_countryr_typefood AS rr ON rr.idtypefood=r.idtypefood WHERE r.idtypefood NOT IN (SELECT bt.idtypefood FROM businessr_typefood AS bt WHERE bt.idbusiness=$1) AND rr.idcountry=$2"
	rows, error_show := db.Query(ctx, q, idbusiness, idcountry)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListPg_TypeFood []models.Pg_R_TypeFood_ToBusiness

	if error_show != nil {
		return oListPg_TypeFood, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var typefoods models.Pg_R_TypeFood_ToBusiness
		rows.Scan(&typefoods.IDTypefood, &typefoods.Name, &typefoods.Url, &typefoods.IsAvailable)
		oListPg_TypeFood = append(oListPg_TypeFood, typefoods)
	}

	//Si todo esta bien
	return oListPg_TypeFood, nil

}
