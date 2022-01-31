package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Add_IntialiData(anfitrionpg models.Mqtt_CreateInitialData) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	//Agregamos el Business
	query := `INSERT INTO Business(idbusiness,idcountry,createddate) VALUES ($1,$2,$3)`
	_, err_add_business := db.Query(ctx, query, anfitrionpg.IDBusiness, anfitrionpg.Country, time.Now())

	if err_add_business != nil {
		return err_add_business
	}

	return nil
}
