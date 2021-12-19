package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/jackc/pgx/v4"
)

func Pg_Find_Open(latitude float64, longitude float64) ([]interface{}, error) {

	db := models.Conectar_Pg_DB()

	//Instanciamos una query
	//var idcomensales []int
	var q string
	var rows pgx.Rows
	var error_show error

	q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  b.isopen=true GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ASC"
	rows, error_show = db.Query(context.Background(), q, latitude, longitude)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListaInterface []interface{}

	if error_show != nil {
		return oListaInterface, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var interfac interface{}
		rows.Scan(&interfac)
		oListaInterface = append(oListaInterface, interfac)
	}

	//Si todo esta bien
	return oListaInterface, nil
}
