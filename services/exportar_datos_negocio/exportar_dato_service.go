package exportar

import (

	//MDOELS

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"

	//REPOSITORIES
	business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/business"
	day_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/day_x_business"
	payment_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/paymenth_x_business"
	recover_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/recover_data_business"
	service_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/service_x_business"
	typefood_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/typefood_x_business"
)

func GetBasicData_Service(idbusiness int) (int, bool, string, models.Pg_BasicData_ToBusiness) {

	var basic_data models.Pg_BasicData_ToBusiness

	//Eliminamos los datos en PG
	basic_data, error_find := business_repository.Pg_Find_BasicData(idbusiness)
	if error_find != nil {
		return 500, true, "Error interno en el servidor al intentar buscar la informacion basica del negocio, detalle: " + error_find.Error(), basic_data
	}
	if basic_data.Name == "" {
		basic_data, error_find_2 := business_repository.Pg_Find_BasicData_WithoutData(idbusiness)
		if error_find_2 != nil {
			return 500, true, "Error interno en el servidor al intentar buscar la informacion basica del negocio, detalle: " + error_find_2.Error(), basic_data
		}
	}

	return 200, false, "", basic_data
}

func GetSchedule_Service(idbusiness int) (int, bool, string, []models.Pg_R_Schedule_ToBusiness) {

	//Eliminamos los datos en PG
	schedule, error_find := day_repository.Pg_Find(idbusiness)
	if error_find != nil {
		return 500, true, "Error interno en el servidor al intentar buscar el horario del negocio, detalle: " + error_find.Error(), schedule
	}

	return 200, false, "", schedule
}

func GetPayment_Service(idbusiness int, country int) (int, bool, string, []models.Pg_R_PaymentMethod_ToBusiness) {

	//Eliminamos los datos en PG
	payment, error_find := payment_repository.Pg_Find(idbusiness, country)
	if error_find != nil {
		return 500, true, "Error interno en el servidor al intentar buscar el horario del negocio, detalle: " + error_find.Error(), payment
	}

	return 200, false, "", payment
}

func GetService_Service(idbusiness int, country int) (int, bool, string, []models.Pg_R_Service_ToBusiness) {

	//Eliminamos los datos en PG
	service, error_find := service_repository.Pg_Find(idbusiness, country)
	if error_find != nil {
		return 500, true, "Error interno en el servidor al intentar buscar el horario del negocio, detalle: " + error_find.Error(), service
	}

	return 200, false, "", service
}

func GetTypeFood_Service(idbusiness int, country int) (int, bool, string, []models.Pg_R_TypeFood_ToBusiness) {

	//Eliminamos los datos en PG
	typefood, error_find := typefood_repository.Pg_Find(idbusiness, country)
	if error_find != nil {
		return 500, true, "Error interno en el servidor al intentar buscar el horario del negocio, detalle: " + error_find.Error(), typefood
	}

	return 200, false, "", typefood
}

/*========================= RECUPERAR DATOS DEL NEGOCIO =========================*/

func GetRecoverAll_Service() (int, bool, string, []models.Mo_Business) {

	//Buscamos todos los negocios a recuperar datos
	all_business, error_find := recover_repository.Pg_Recover_All()
	if error_find != nil {
		return 500, true, "Error interno en el servidor al intentar buscar el horario del negocio, detalle: " + error_find.Error(), all_business
	}

	return 200, false, "", all_business
}

func GetRecoverOne_Service(idbusiness int) (int, bool, string, models.Mo_Business) {

	//Buscamos un negocio a recuperar datos
	one_business, error_find := recover_repository.Pg_Recover_One(idbusiness)
	if error_find != nil {
		return 500, true, "Error interno en el servidor al intentar buscar el horario del negocio, detalle: " + error_find.Error(), one_business
	}

	return 200, false, "", one_business
}
