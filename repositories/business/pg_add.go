package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Add_IntialiData(anfitrionpg models.Mqtt_CreateInitialData) error {

	db := models.Conectar_Pg_DB()

	//Agregamos el Business

	_, err_add_business := db.Exec(context.Background(), "INSERT INTO Business(idbusiness,idcountry,createdDate,isopen) VALUES ($1,$2,$3,$4) RETURNING idbusiness", anfitrionpg.IDBusiness, anfitrionpg.Country, time.Now(), false)
	if err_add_business != nil {
		return err_add_business
	}

	return nil
}
