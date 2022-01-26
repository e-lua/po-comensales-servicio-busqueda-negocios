package busqueda

import (
	//REPOSITORIES
	"github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/business"
	favorite_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/favorites"
	paymenth_x_business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/paymenth_x_business"
	typefood_x_business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/typefood_x_business"
)

/*----------------------TRAEMOS LOS DATOS----------------------*/

func GetBusinessCards_SearchedBefore_Service(input_data_idcomensal int) (int, bool, string, BusinessCards_SearchedBefore) {

	var business_cards BusinessCards_SearchedBefore

	//Buscamos si ya he visitado negocios antes
	quantity, all_business_searched, _ := business_repository.Pg_Find_Near_Searched(input_data_idcomensal)

	business_cards.Quantity = quantity
	business_cards.Business = all_business_searched

	return 200, false, "", business_cards
}

func GetBusinessCards_Service(latitude float64, longitude float64, services []int, typefood []int, payment []int, input_data_idcomensal int) (int, bool, string, []models.Pg_Found_All_Business) {

	//Buscamos los negocios
	business_cards, error_find_pg := business_repository.Pg_Find_All(latitude, longitude, services, typefood, payment, input_data_idcomensal)
	if error_find_pg != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los negocios cercanos, detalle: " + error_find_pg.Error(), business_cards
	}
	return 200, false, "", business_cards

}

func GetBusinessCards_Open_Service(latitude float64, longitude float64, services []int, typefood []int, payment []int) (int, bool, string, []models.Pg_Found_All_Business) {

	//Buscamos los negocios
	business_cards, error_find_pg := business_repository.Pg_Find_Open(latitude, longitude, services, typefood, payment)
	if error_find_pg != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los negocios cercanos abiertos, detalle: " + error_find_pg.Error(), business_cards
	}
	return 200, false, "", business_cards

}

func GetBusinessCardsByName_Service(name string, tipo int) (int, bool, string, []models.Pg_Found_All_Business) {

	var business_cards []models.Pg_Found_All_Business
	var error_find_pg error

	if tipo == 1 {
		//Buscamos los negocios con @
		business_cards, error_find_pg = business_repository.Pg_SearchByUniqueName(name)
		if error_find_pg != nil {
			return 500, true, "Error interno en el servidor al intentar buscar los negocios con el nombre detallado, detalle: " + error_find_pg.Error(), business_cards
		}

	} else {

		//Buscamos los negocios sin @
		business_cards, error_find_pg = business_repository.Pg_SearchByName(name)
		if error_find_pg != nil {
			return 500, true, "Error interno en el servidor al intentar buscar los negocios con el nombre detallado, detalle: " + error_find_pg.Error(), business_cards
		}
	}

	return 200, false, "", business_cards

}

func GetFavorites_Service(input_data_idcomensal int) (int, bool, string, []models.Pg_Found_All_Business) {

	//Buscamos los negocios
	business_cards, error_find_pg := favorite_repository.Pg_Find(input_data_idcomensal)
	if error_find_pg != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los negocios favoritos, detalle: " + error_find_pg.Error(), business_cards
	}
	return 200, false, "", business_cards

}

func GetFilterTypeFoods_Service(idcountry int) (int, bool, string, []models.Pg_R_TypeFood) {

	//Buscamos los negocios
	filter_typefood, error_find_pg := typefood_x_business_repository.Pg_Find_Filter(idcountry)
	if error_find_pg != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los filtros disponibles en los metodos de pago, detalle: " + error_find_pg.Error(), filter_typefood
	}

	return 200, false, "", filter_typefood
}

func GetFilterPaymentMethods_Service(idcountry int) (int, bool, string, []models.Pg_R_PaymentMethod) {

	//Buscamos los negocios
	filter_payments, error_find_pg := paymenth_x_business_repository.Pg_Find_Filter(idcountry)
	if error_find_pg != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los filtros disponibles en los metodos de pago, detalle: " + error_find_pg.Error(), filter_payments
	}

	return 200, false, "", filter_payments
}

/*----------------------INSERTAMOS LOS DATOS----------------------*/

func AddFavorites_Service(idcomensal int, idbusiness int) (int, bool, string, string) {

	//Buscamos los negocios
	error_add := favorite_repository.Pg_Add(idcomensal, idbusiness)
	if error_add != nil {
		return 500, true, "Error interno en el servidor al intentar agregar el negocio como favorito, detalle: " + error_add.Error(), ""
	}

	return 200, false, "", "Favorito agregado"
}

func GetUniqueNames_Service(uniquename_string string) (int, bool, string, string) {

	//Buscamos los negocios
	uniquenames, error_add := business_repository.Pg_Find_Uniquename(uniquename_string)
	if error_add != nil {
		return 500, true, "Error interno en el servidor al intentar agregar el negocio como favorito, detalle: " + error_add.Error(), uniquenames
	}

	return 200, false, "", uniquenames
}

/*=============================== INICIO TEST===============================*/

func GetBusinessCards_Test_Service(latitude float64, longitude float64, services []int, typefood []int, payment []int, input_data_idcomensal int) (int, bool, string, []models.Pg_Found_All_Business) {

	//Buscamos los negocios
	business_cards, error_find_pg := business_repository.Pg_Find_All_Test(latitude, longitude, services, typefood, payment, input_data_idcomensal)
	if error_find_pg != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los negocios cercanos, detalle: " + error_find_pg.Error(), business_cards
	}
	return 200, false, "", business_cards
}

func GetBusinessCardsByName_Test_Service(name string) (int, bool, string, []models.Pg_Found_All_Business) {

	//Buscamos los negocios
	business_cards, error_find_pg := business_repository.Pg_SearchByName_Test(name)
	if error_find_pg != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los negocios con el nombre detallado, detalle: " + error_find_pg.Error(), business_cards
	}
	return 200, false, "", business_cards

}

/*=============================== FIN TEST===============================*/
