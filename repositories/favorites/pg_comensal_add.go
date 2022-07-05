package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Comensal_Add(idcomensal int, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB_Comensal()

	query := `INSERT INTO Favorites(idbusiness,idcomensal,date) VALUES ($1,$2,$3)`
	_, err := db.Query(ctx, query, idbusiness, idcomensal, time.Now())

	if err != nil {
		return err
	}

	return nil
}
