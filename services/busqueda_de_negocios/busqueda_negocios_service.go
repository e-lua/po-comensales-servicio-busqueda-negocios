package busqueda

import (
	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"

	//REPOSITORIES
	business_repository "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/repositories/business"
)

func GetAllBusiness_Service(input_loc models.Location) (int, bool, string, []models.Mo_Business_Cards) {

	//Eliminamos los datos en PG
	all_business, error_find_pg := business_repository.Mo_Find_All(input_loc)
	if error_find_pg != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los negocios cercanos, detalle: " + error_find_pg.Error(), all_business
	}

	return 200, false, "", all_business
}

func GetAllBusiness_2_Service(input_loc models.Location) (int, bool, string, []interface{}) {

	//Eliminamos los datos en PG
	als_business, error_find_pg := business_repository.Pg_Find_All()
	if error_find_pg != nil {
		return 500, true, "Error interno en el servidor al intentar buscar los negocios cercanos, detalle: " + error_find_pg.Error(), als_business
	}

	return 200, false, "", als_business
}
