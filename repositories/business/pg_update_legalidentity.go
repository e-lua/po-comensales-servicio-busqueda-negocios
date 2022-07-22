package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_UpdateLegalIdentity(input_mqtt_legalidentity models.Mqtt_LegalIdentity) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Cambio de Server y BD, ya que no se puede acceder al rol de superusuario para la busqueda por distancia
	db := models.Conectar_Pg_DB_Comensal()

	query := `UPDATE Business SET legalidentity=$1,iva=$2,typesuscription=$3 WHERE idbusiness=$4`
	if _, err_update := db.Exec(ctx, query, input_mqtt_legalidentity.LegalIdentity, input_mqtt_legalidentity.IVA, input_mqtt_legalidentity.Typesuscription, input_mqtt_legalidentity.IdBusiness); err_update != nil {
		return err_update
	}

	return nil
}

func Pg_UpdateLegalIdentity_Multiple(input_mqtt_legalidentity_multiple []models.Mqtt_LegalIdentity) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	idbusiness_pg, legalidentity_pg, iva_pg, typesuscription_pg := []int{}, []string{}, []float32{}, []int{}

	for _, input_mqtt_legalidentity := range input_mqtt_legalidentity_multiple {
		idbusiness_pg = append(idbusiness_pg, input_mqtt_legalidentity.IdBusiness)
		legalidentity_pg = append(legalidentity_pg, input_mqtt_legalidentity.LegalIdentity)
		iva_pg = append(iva_pg, input_mqtt_legalidentity.IVA)
		typesuscription_pg = append(typesuscription_pg, input_mqtt_legalidentity.Typesuscription)
	}

	//Cambio de Server y BD, ya que no se puede acceder al rol de superusuario para la busqueda por distancia
	db := models.Conectar_Pg_DB_Comensal()

	query := `UPDATE Business SET legalidentity=ex.legalid,iva=ex.iv,typesuscription=ex.typesusc FROM (select * from  unnest($1::varchar(50)[],$2::decimal(10,2)[],$3::int[],$4::int[])) as ex(legalid,iv,typesusc,idbss) WHERE idbusiness=ex.idbss`
	if _, err_update := db.Exec(ctx, query, legalidentity_pg, iva_pg, typesuscription_pg, idbusiness_pg); err_update != nil {
		return err_update
	}

	return nil
}
