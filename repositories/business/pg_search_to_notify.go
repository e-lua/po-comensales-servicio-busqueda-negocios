package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_SearchToNotify() ([]int, int, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Cambio de Server y BD, ya que no se puede acceder al rol de superusuario para la busqueda por distancia

	db := models.Conectar_Pg_DB_Comensal()

	q := "SELECT idbusiness,name FROM business WHERE latitude is not null OR uniquename <>'@'"
	rows, error_shown := db.Query(ctx, q)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListBusiness []int
	quantity := 0

	if error_shown != nil {

		return oListBusiness, quantity, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oBusiness int
		rows.Scan(&oBusiness)
		oListBusiness = append(oListBusiness, oBusiness)
		quantity = quantity + 1
	}

	//Si todo esta bien
	return oListBusiness, quantity, nil

}
