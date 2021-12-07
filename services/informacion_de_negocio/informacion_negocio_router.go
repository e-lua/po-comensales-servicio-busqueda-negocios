package informacion

import (
	"log"

	"github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

var InformationRouter_pg *informationRouter_pg

type informationRouter_pg struct {
}

/*----------------------UPDATE MQTT DATA----------------------*/

func (cr *informationRouter_pg) UpdatePaymenth(inputserialize_payment models.Mqtt_PaymentMethod) {
	//Enviamos los datos al servicio
	error_r := UpdatePaymenth_Service(inputserialize_payment)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) UpdateService(inputserialize_service models.Mqtt_Service) {
	//Enviamos los datos al servicio
	error_r := UpdateService_Service(inputserialize_service)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) UpdateTypeFood(inputserialize_typefood models.Mqtt_TypeFood) {
	//Enviamos los datos al servicio
	error_r := UpdateTypeFood_Service(inputserialize_typefood)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) UpdateName(inputserialize_name models.Mqtt_Name) {
	//Enviamos los datos al servicio
	error_r := UpdateName_Service(inputserialize_name)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) UpdateOpen(inputserialize_open models.Mqtt_IsOpen) {
	//Enviamos los datos al servicio
	error_r := UpdateOpen_Service(inputserialize_open)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) UpdateBanner(inputserialize_banner models.Mqtt_Banner_Cola) {
	//Enviamos los datos al servicio
	error_r := UpdateBanner_Service(inputserialize_banner)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) UpdateAddress(inputserialize_address models.Mqtt_Addres) {
	//Enviamos los datos al servicio
	error_r := UpdateAddress_Service(inputserialize_address)
	if error_r != nil {
		log.Fatal(error_r)
	}
}
