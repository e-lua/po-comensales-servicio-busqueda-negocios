package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Add(idcomensal int, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db_external := models.Conectar_Pg_DB()

	query := `INSERT INTO Favorites(idbusiness,idcomensal,date) VALUES ($1,$2,$3)`
	_, err := db_external.Query(ctx, query, idbusiness, idcomensal, time.Now())

	if err != nil {
		return err
	}

	return nil
}
