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

	//Cambio de Server y BD, ya que no se puede acceder al rol de superusuario para la busqueda por distancia
	db := models.Conectar_Pg_DB_Comensal()

	//Agregamos el Business
	query := `INSERT INTO Business(idbusiness,idcountry,createddate,issubsidiary,subsidiaryof) VALUES ($1,$2,$3,$4,$5)`
	_, err_add_business := db.Query(ctx, query, anfitrionpg.IDBusiness, anfitrionpg.Country, time.Now(), anfitrionpg.IsSubsidiary, anfitrionpg.SubsidiaryOf)

	if err_add_business != nil {
		return err_add_business
	}

	return nil
}
