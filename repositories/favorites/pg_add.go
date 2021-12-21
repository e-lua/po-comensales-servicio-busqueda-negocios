package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Add(idcomensal int, idbusiness int) error {

	db_external := models.Conectar_Pg_DB()

	query := `INSERT INTO Favorites(idbusiness,idcomensl,date) VALUES ($1,$2,$3)`
	_, err := db_external.Query(context.Background(), query, idbusiness, idcomensal, time.Now())

	if err != nil {
		return err
	}

	return nil
}
