package busqueda

import (
	//REPOSITORIES
	"github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/business"
	paymenth_x_business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/paymenth_x_business"
	typefood_x_business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/typefood_x_business"
)

func GetBusinessCards_SearchedBefore_Service(input_data_idcomensal int) (int, bool, string, interface{}) {

	//Buscamos si ya he visitado negocios antes
	all_business_searched, _ := business_repository.Re_Get_Near_Business(input_data_idcomensal)

	return 200, false, "", all_business_searched
}

func GetBusinessCards_Service(input_search_filters SearchFilters, input_data_idcomensal int) (int, bool, string, []interface{}) {

	//Buscamos los negocios
	business_cards, error_find_pg := business_repository.Pg_Find_All(input_search_filters.Latitude, input_search_filters.Longitude, input_search_filters.Services, input_search_filters.TypeFood, input_search_filters.Payment, input_data_idcomensal)
	if error_find_pg != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los negocios cercanos, detalle: " + error_find_pg.Error(), business_cards
	}
	return 200, false, "", business_cards

}

func GetBusinessCards_Open_Service(input_search_filters SearchFilters, input_data_idcomensal int) (int, bool, string, []interface{}) {

	//Buscamos los negocios
	business_cards, error_find_pg := business_repository.Pg_Find_Open(input_search_filters.Latitude, input_search_filters.Longitude)
	if error_find_pg != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los negocios cercanos abiertos, detalle: " + error_find_pg.Error(), business_cards
	}
	return 200, false, "", business_cards

}

func GetBusinessCards_Favorite_Service(input_data_idcomensal int) (int, bool, string, []interface{}) {

	//Buscamos los negocios
	business_cards, error_find_pg := business_repository.Pg_Find_Favorite(input_data_idcomensal)
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
