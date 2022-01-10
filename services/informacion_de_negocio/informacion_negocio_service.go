package informacion

import (

	//MDOELS
	"log"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"

	//REPOSITORIES
	business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/business"
	day_business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/day_x_business"
	payment_business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/paymenth_x_business"
	service_business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/service_x_business"
	typefood_business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/typefood_x_business"
)

func CreateBusiness_Service(input_mqtt_create models.Mqtt_CreateInitialData) error {

	if input_mqtt_create.Country > 2 {
		//Eliminamos los datos en PG
		error_delete_pg := business_repository.Pg_Add_IntialiData(input_mqtt_create)
		if error_delete_pg != nil {
			log.Fatal(error_delete_pg)
		}
	}

	return nil
}

func UpdatePaymenth_Service(input_mqtt_payment models.Mqtt_PaymentMethod) error {

	//Eliminamos los datos en PG
	error_delete_pg := payment_business_repository.Pg_Delete_Update(input_mqtt_payment)
	if error_delete_pg != nil {
		log.Fatal(error_delete_pg)
	}

	return nil
}

func UpdateSchedule_Service(input_mqtt_schedule models.Mqtt_Schedule) error {

	//Eliminamos los datos en PG
	error_delete_pg := day_business_repository.Pg_Delete_Update(input_mqtt_schedule)
	if error_delete_pg != nil {
		log.Fatal(error_delete_pg)
	}

	return nil
}

func UpdateService_Service(input_mqtt_service models.Mqtt_Service) error {

	//Eliminamos los datos en PG
	error_delete_pg := service_business_repository.Pg_Delete_Update(input_mqtt_service)
	if error_delete_pg != nil {
		log.Fatal(error_delete_pg)
	}

	return nil
}

func UpdateTypeFood_Service(input_mqtt_typefood models.Mqtt_TypeFood) error {

	//Eliminamos los datos en PG
	error_delete_pg := typefood_business_repository.Pg_Delete_Update(input_mqtt_typefood)
	if error_delete_pg != nil {
		log.Fatal(error_delete_pg)
	}

	return nil
}

func UpdateName_Service(input_mqtt_name models.Mqtt_Name) error {

	//Insertamos los datos en PG

	error_add_pg := business_repository.Pg_UpdateName(input_mqtt_name)
	if error_add_pg != nil {
		log.Fatal(error_add_pg)
	}

	return nil
}

func UpdateTimeZone_Service(input_mqtt_open models.Mqtt_TimeZone) error {

	//Insertamos los datos en PG

	error_add_pg := business_repository.Pg_UpdateTimeZone(input_mqtt_open)
	if error_add_pg != nil {
		log.Fatal(error_add_pg)
	}

	return nil
}

func UpdateAddress_Service(input_mqtt_address models.Mqtt_Addres) error {

	//Insertamos los datos en PG

	error_add_pg := business_repository.Pg_UpdateAddress(input_mqtt_address)
	if error_add_pg != nil {
		log.Fatal(error_add_pg)
	}

	return nil
}

func UpdateBanner_Service(input_mqtt_banner models.Mqtt_Banner_Cola) error {

	//Insertamos los datos en PG

	error_add_pg := business_repository.Pg_UpdateBanner(input_mqtt_banner)
	if error_add_pg != nil {
		log.Fatal(error_add_pg)
	}

	return nil
}

func FindAddress_Service(inputObjectIdBusiness int) (int, bool, string, B_Address) {

	var business_address B_Address

	latitude, longitude, error_address := business_repository.Pg_Find_Address(inputObjectIdBusiness)
	if error_address != nil {
		return 500, true, "Error interno en el servidor al intentar buscar la latitud y longitude, detalle: " + error_address.Error(), business_address
	}

	business_address.Latitude = latitude
	business_address.Longitude = longitude

	return 200, false, "", business_address
}
