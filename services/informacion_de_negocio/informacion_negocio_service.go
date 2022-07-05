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

	error_delete_pg := business_repository.Pg_Add_IntialiData(input_mqtt_create)
	if error_delete_pg != nil {
		log.Fatal(error_delete_pg)
	}

	//Insertamos los datos en la BD que consultará el comensal
	error_delete_pg_comensal := business_repository.Pg_Comensal_Add_IntialiData(input_mqtt_create)
	if error_delete_pg_comensal != nil {
		log.Fatal(error_delete_pg_comensal)
	}

	return nil
}

func AddViewInformation_Service(idbusiness int, idcomensal int) (int, bool, string, string) {

	business_repository.Mqtt_ExportView(idbusiness, idcomensal)

	return 200, false, "", "Vista registrada"
}

func UpdatePaymenth_Service(input_mqtt_payment models.Mqtt_PaymentMethod) error {

	error_delete_pg := payment_business_repository.Pg_Delete_Update(input_mqtt_payment)
	if error_delete_pg != nil {
		log.Fatal(error_delete_pg)
	}

	//Insertamos los datos en la BD que consultará el comensal
	error_delete_pg_comensal := payment_business_repository.Pg_Comensal_Delete_Update(input_mqtt_payment)
	if error_delete_pg_comensal != nil {
		log.Fatal(error_delete_pg_comensal)
	}

	return nil
}

func UpdateSchedule_Service(input_mqtt_schedule models.Mqtt_Schedule) error {

	error_delete_pg := day_business_repository.Pg_Delete_Update(input_mqtt_schedule)
	if error_delete_pg != nil {
		log.Fatal(error_delete_pg)
	}

	//Insertamos los datos en la BD que consultará el comensal
	error_delete_pg_comensal := day_business_repository.Pg_Comensal_Delete_Update(input_mqtt_schedule)
	if error_delete_pg_comensal != nil {
		log.Fatal(error_delete_pg_comensal)
	}

	return nil
}

func UpdateService_Service(input_mqtt_service models.Mqtt_Service) error {

	error_delete_pg := service_business_repository.Pg_Delete_Update(input_mqtt_service)
	if error_delete_pg != nil {
		log.Fatal(error_delete_pg)
	}

	//Insertamos los datos en la BD que consultará el comensal
	error_delete_pg_comensal := service_business_repository.Pg_Comensal_Delete_Update(input_mqtt_service)
	if error_delete_pg_comensal != nil {
		log.Fatal(error_delete_pg_comensal)
	}

	return nil
}

func UpdateTypeFood_Service(input_mqtt_typefood models.Mqtt_TypeFood) error {

	error_delete_pg := typefood_business_repository.Pg_Delete_Update(input_mqtt_typefood)
	if error_delete_pg != nil {
		log.Fatal(error_delete_pg)
	}

	//Insertamos los datos en la BD que consultará el comensal
	error_delete_pg_comensal := typefood_business_repository.Pg_Comensal_Delete_Update(input_mqtt_typefood)
	if error_delete_pg_comensal != nil {
		log.Fatal(error_delete_pg_comensal)
	}

	return nil
}

func UpdateName_Service(input_mqtt_name models.Mqtt_Name) error {

	//Insertamos los datos en PG
	error_add_pg := business_repository.Pg_UpdateName(input_mqtt_name)
	if error_add_pg != nil {
		log.Fatal(error_add_pg)
	}

	//Insertamos los datos en la BD que consultará el comensal
	error_delete_pg_comensal := business_repository.Pg_Comensal_UpdateName(input_mqtt_name)
	if error_delete_pg_comensal != nil {
		log.Fatal(error_delete_pg_comensal)
	}

	//Guardamos los datos en la memoria cache
	basic_data_re, error_get_cache := business_repository.Re_Get_BasicData_Business(input_mqtt_name.IdBusiness)
	if error_get_cache != nil {
		log.Fatal(error_get_cache)
	}
	basic_data_re.Basic_Data.Name = input_mqtt_name.Name
	err_add_cache := business_repository.Re_Set_BasicData_Business(input_mqtt_name.IdBusiness, basic_data_re.Basic_Data)
	if err_add_cache != nil {
		log.Fatal(err_add_cache)
	}

	return nil
}

func UpdateLegalIdentity_Service(inputserialize_legalidentity models.Mqtt_LegalIdentity) error {

	//Insertamos los datos en PG
	error_add_pg := business_repository.Pg_UpdateLegalIdentity(inputserialize_legalidentity)
	if error_add_pg != nil {
		log.Fatal(error_add_pg)
	}

	//Guardamos los datos en la memoria cache
	basic_data_re, error_get_cache := business_repository.Re_Get_BasicData_Business(inputserialize_legalidentity.IdBusiness)
	if error_get_cache != nil {
		log.Fatal(error_get_cache)
	}
	basic_data_re.Basic_Data.Legalidentity = inputserialize_legalidentity.LegalIdentity
	basic_data_re.Basic_Data.IVA = inputserialize_legalidentity.IVA
	err_add_cache := business_repository.Re_Set_BasicData_Business(inputserialize_legalidentity.IdBusiness, basic_data_re.Basic_Data)
	if err_add_cache != nil {
		log.Fatal(err_add_cache)
	}

	return nil
}

func UpdateUniqueName_Service(input_mqtt_uniquename models.Mqtt_Uniquename) error {

	//Insertamos los datos en PG
	error_add_pg := business_repository.Pg_UpdateUniqueName(input_mqtt_uniquename)
	if error_add_pg != nil {
		log.Fatal(error_add_pg)
	}

	//Insertamos los datos en la BD que consultará el comensal
	error_add_pg_comensal := business_repository.Pg_Comensal_UpdateUniqueName(input_mqtt_uniquename)
	if error_add_pg_comensal != nil {
		log.Fatal(error_add_pg_comensal)
	}

	//Guardamos los datos en la memoria cache
	basic_data_re, error_get_cache := business_repository.Re_Get_BasicData_Business(input_mqtt_uniquename.IdBusiness)
	if error_get_cache != nil {
		log.Fatal(error_get_cache)
	}
	basic_data_re.Basic_Data.Uniquename = input_mqtt_uniquename.Uniquename
	err_add_cache := business_repository.Re_Set_BasicData_Business(input_mqtt_uniquename.IdBusiness, basic_data_re.Basic_Data)
	if err_add_cache != nil {
		log.Fatal(err_add_cache)
	}

	return nil
}

func UpdateTimeZone_Service(input_mqtt_open models.Mqtt_TimeZone) error {

	//Insertamos los datos en PG
	error_add_pg := business_repository.Pg_UpdateTimeZone(input_mqtt_open)
	if error_add_pg != nil {
		log.Fatal(error_add_pg)
	}

	//Insertamos los datos en la BD que consultará el comensal
	error_add_pg_comensal := business_repository.Pg_Comensal_UpdateTimeZone(input_mqtt_open)
	if error_add_pg_comensal != nil {
		log.Fatal(error_add_pg_comensal)
	}

	//Guardamos los datos en la memoria cache
	basic_data_re, error_get_cache := business_repository.Re_Get_BasicData_Business(input_mqtt_open.IdBusiness)
	if error_get_cache != nil {
		log.Fatal(error_get_cache)
	}
	basic_data_re.Basic_Data.TimeZone = input_mqtt_open.TimeZone
	err_add_cache := business_repository.Re_Set_BasicData_Business(input_mqtt_open.IdBusiness, basic_data_re.Basic_Data)
	if err_add_cache != nil {
		log.Fatal(err_add_cache)
	}

	return nil
}

func UpdateAddress_Service(input_mqtt_address models.Mqtt_Addres) error {

	//Insertamos los datos en PG
	error_add_pg := business_repository.Pg_UpdateAddress(input_mqtt_address)
	if error_add_pg != nil {
		log.Fatal(error_add_pg)
	}

	//Insertamos los datos en PG
	error_add_pg_comensal := business_repository.Pg_Comensal_UpdateAddress(input_mqtt_address)
	if error_add_pg_comensal != nil {
		log.Fatal(error_add_pg_comensal)
	}

	return nil
}

func UpdateBanner_Service(input_mqtt_banner models.Mqtt_Banner_Cola) error {

	//Insertamos los datos en PG
	error_add_pg := business_repository.Pg_UpdateBanner(input_mqtt_banner)
	if error_add_pg != nil {
		log.Fatal(error_add_pg)
	}

	//Insertamos los datos en PG
	error_add_pg_comensal := business_repository.Pg_Comensal_UpdateBanner(input_mqtt_banner)
	if error_add_pg_comensal != nil {
		log.Fatal(error_add_pg_comensal)
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
