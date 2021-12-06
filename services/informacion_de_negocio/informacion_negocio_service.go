package informacion

import (

	//MDOELS
	"log"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"

	//REPOSITORIES
	payment_business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/paymenth_x_business"
)

func RegisterPaymenth_Service(input_mqtt_payment models.Mqtt_PaymentMethod) error {

	//Insertamos los datos en PG
	error_add_pg := payment_business_repository.Pg_Update(input_mqtt_payment)
	if error_add_pg != nil {
		log.Fatal(error_add_pg)
	}
	return nil
}
