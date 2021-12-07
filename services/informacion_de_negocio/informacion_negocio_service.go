package informacion

import (

	//MDOELS
	"log"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"

	//REPOSITORIES
	business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/business"
	payment_business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/paymenth_x_business"
	service_business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/service_x_business"
	typefood_business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/typefood_x_business"
)

func UpdatePaymenth_Service(input_mqtt_payment models.Mqtt_PaymentMethod) error {

	//Eliminamos los datos en PG
	error_delete_pg := payment_business_repository.Pg_Delete(input_mqtt_payment.IdBusiness)
	if error_delete_pg != nil {
		log.Fatal(error_delete_pg)
	}

	time.Sleep(1 * time.Second)

	//Insertamos los datos en PG
	go func() {
		error_add_pg := payment_business_repository.Pg_Update(input_mqtt_payment)
		if error_add_pg != nil {
			log.Fatal(error_add_pg)
		}
	}()

	time.Sleep(2 * time.Second)

	return nil
}

func UpdateService_Service(input_mqtt_service models.Mqtt_Service) error {

	//Eliminamos los datos en PG
	error_delete_pg := service_business_repository.Pg_Delete(input_mqtt_service.IdBusiness)
	if error_delete_pg != nil {
		log.Fatal(error_delete_pg)
	}

	time.Sleep(1 * time.Second)

	//Insertamos los datos en PG
	go func() {
		error_add_pg := service_business_repository.Pg_Update(input_mqtt_service)
		if error_add_pg != nil {
			log.Fatal(error_add_pg)
		}
	}()

	time.Sleep(2 * time.Second)

	return nil
}

func UpdateTypeFood_Service(input_mqtt_typefood models.Mqtt_TypeFood) error {

	//Eliminamos los datos en PG
	error_delete_pg := typefood_business_repository.Pg_Delete(input_mqtt_typefood.IdBusiness)
	if error_delete_pg != nil {
		log.Fatal(error_delete_pg)
	}

	time.Sleep(1 * time.Second)

	//Insertamos los datos en PG
	go func() {
		error_add_pg := typefood_business_repository.Pg_Update(input_mqtt_typefood)
		if error_add_pg != nil {
			log.Fatal(error_add_pg)
		}
	}()

	time.Sleep(2 * time.Second)

	return nil
}

func UpdateName_Service(input_mqtt_name models.Mqtt_Name) error {

	//Insertamos los datos en PG
	go func() {
		error_add_pg := business_repository.Pg_UpdateName(input_mqtt_name)
		if error_add_pg != nil {
			log.Fatal(error_add_pg)
		}
	}()

	time.Sleep(2 * time.Second)

	return nil
}

func UpdateOpen_Service(input_mqtt_open models.Mqtt_IsOpen) error {

	time.Sleep(1 * time.Second)

	//Insertamos los datos en PG
	go func() {
		error_add_pg := business_repository.Pg_UpdateIsOpen(input_mqtt_open)
		if error_add_pg != nil {
			log.Fatal(error_add_pg)
		}
	}()

	time.Sleep(2 * time.Second)

	return nil
}

func UpdateAddress_Service(input_mqtt_address models.Mqtt_Addres) error {

	time.Sleep(1 * time.Second)

	//Insertamos los datos en PG
	go func() {
		error_add_pg := business_repository.Pg_UpdateAddress(input_mqtt_address)
		if error_add_pg != nil {
			log.Fatal(error_add_pg)
		}
	}()

	time.Sleep(2 * time.Second)

	return nil
}

func UpdateBanner_Service(input_mqtt_banner models.Mqtt_Banner_Cola) error {

	time.Sleep(1 * time.Second)

	//Insertamos los datos en PG
	go func() {
		error_add_pg := business_repository.Pg_UpdateBanner(input_mqtt_banner)
		if error_add_pg != nil {
			log.Fatal(error_add_pg)
		}
	}()

	time.Sleep(2 * time.Second)

	return nil
}
