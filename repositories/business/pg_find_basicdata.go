package repositories

import (
	models "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/models"
)

func Pg_Find_BasicData(idbusiness int) (string, float32, float32, int, string, string, string, error) {

	var name string
	var latitude float32
	var longitude float32
	var fulladdress string
	var referenceadd string
	var postalcode int
	var deliveryrange string

	db := models.Conectar_Pg_DB()
	q := "SELECT name,latitude,longitude,postalcode,fulladdress,referenceaddress,deliveryrangedescrip FROM Business WHERE idbusiness=$1"
	error_showname := db.QueryRow(q, idbusiness).Scan(&name, &latitude, &longitude, &postalcode, &fulladdress, &referenceadd, &deliveryrange)

	if error_showname != nil {
		defer db.Close()
		return name, latitude, longitude, postalcode, fulladdress, referenceadd, deliveryrange, error_showname
	}

	defer db.Close()

	//Si todo esta bien
	return name, latitude, longitude, postalcode, fulladdress, referenceadd, deliveryrange, nil

}
