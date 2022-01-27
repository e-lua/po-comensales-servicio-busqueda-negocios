package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_UpdateAddress(urlphoto string, idbusiness int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `UPDATE Business SET urlbanner=$1 WHERE idbusiness=$2`
	if _, err_update := db.Exec(ctx, query, urlphoto, idbusiness); err_update != nil {
		return err_update
	}

	return nil
}
