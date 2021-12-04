package informacion

import (
	banner_x_busines_repository "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/repositories/banner_x_business"
	business_repository "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/repositories/business"
	contact_x_business_repository "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/repositories/contact_x_business"
	schedule_x_business_repository "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/repositories/day_x_business"
	payment_x_business_repository "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/repositories/paymenth_x_business"
	service_x_business_repository "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/repositories/service_x_business"
	typefood_x_business_repository "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/repositories/typefood_x_business"
)

func GetInformationData_Pg_Service(inputidbusiness int) (int, bool, string, ResponseWithStructBusiness) {

	//Instanciamos el modelo de respuesta
	var oRespondeWithStructure ResponseWithStructBusiness

	//Validamos la existencia del objectId en Business_Pg
	name, latitude, longitude, postalcode, fulladdress, referenceadd, deliveryrange, _ := business_repository.Pg_Find_BasicData(inputidbusiness)

	//===============ASIGNAMOS LOS VALORES A LA VARIABLE INSTANCIADA===============

	oRespondeWithStructure.NameBusiness = name
	oRespondeWithStructure.LatitudeBusiness = latitude
	oRespondeWithStructure.LongitudeBusiness = longitude
	oRespondeWithStructure.Fulladdress = fulladdress
	oRespondeWithStructure.ReferenceAddress = referenceadd
	oRespondeWithStructure.PostalCode = postalcode
	oRespondeWithStructure.DeliveryRange = deliveryrange

	//===============ASIGNAMOS LOS VALORES A LA VARIABLE INSTANCIADA

	//Treamos la portada
	type_food_x_business, _ := typefood_x_business_repository.Pg_Find(inputidbusiness)
	//Asignamos la portada
	oRespondeWithStructure.TypeOfFood = type_food_x_business

	//Treamos el tipo de comida
	banner_x_business, _ := banner_x_busines_repository.Pg_Find(inputidbusiness)
	//Asignamos el tipo de comida
	oRespondeWithStructure.Banner = banner_x_business

	//Traemos los servicios
	service_x_business, _ := service_x_business_repository.Pg_Find(inputidbusiness)
	//Asignamos los servicios
	oRespondeWithStructure.Services = service_x_business

	//Traemos los metodos de pago
	payment_x_business, _ := payment_x_business_repository.Pg_Find(inputidbusiness)
	//Asignamos los metodos de pago
	oRespondeWithStructure.PaymentMethods = payment_x_business

	//Traemos los horarios
	day_x_business, _ := schedule_x_business_repository.Pg_Find(inputidbusiness)
	//Asignamos los horarios
	oRespondeWithStructure.DailySchedule = day_x_business

	//Traemos los numeros de contacto
	contact, _ := contact_x_business_repository.Pg_Find(inputidbusiness)
	//Asignamos los numeros de contacto
	oRespondeWithStructure.PhoneContact = contact

	return 200, false, "", oRespondeWithStructure
}
