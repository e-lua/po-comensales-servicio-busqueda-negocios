package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/jackc/pgx/v4"
)

func Pg_Find_All(latitude float64, longitude float64, services []int, typefood []int, payment []int, idcomensal int) ([]interface{}, error) {

	db := models.Conectar_Pg_DB()

	//Instanciamos una query
	var idcomensales []int
	var q string
	var rows pgx.Rows
	var error_show error

	//Agregamos un contador para la consulta
	counter := 0
	if services != nil {
		counter = counter + 1
	}
	if typefood != nil {
		counter = counter + 10
	}
	if payment != nil {
		counter = counter + 20
	}

	//Buscamos la consulta ideal
	switch counter {
	case 1:
		q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bs.idservice IN ($3::int[]) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, services)
	case 10:
		q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bt.idtypefood IN ($3::int[]) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, typefood)
	case 20:
		q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bp.idpayment IN ($3::int[]) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, payment)
	case 11:
		q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bs.idservice IN ($3::int[]) AND bt.idtypefood IN ($4::int[]) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, services, typefood)
	case 21:
		q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bp.idpayment IN ($3::int[]) AND bs.idservice IN ($4::int[]) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, payment, services)
	case 30:
		q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bp.idpayment IN ($3::int[]) AND bt.idtypefood IN ($4::int[]) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, payment, typefood)
	default:
		q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude)
	}

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListaInterface []interface{}

	if error_show != nil {
		return oListaInterface, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var interfac interface{}
		rows.Scan(&interfac)
		//idcomensales = append(idcomensales, idcomensal)
		oListaInterface = append(oListaInterface, interfac)
	}

	insertFoundBusiness(idcomensales, oListaInterface)

	//Si todo esta bien
	return oListaInterface, nil
}

func insertFoundBusiness(idcomensales []int, business []interface{}) error {
	db := models.Conectar_Pg_DB()

	query := `INSERT INTO Near(idcomensal,nearbusiness) (select * from unnest($1::int[], $2::int[]))`
	if _, err := db.Exec(context.Background(), query, idcomensales, business); err != nil {
		return err
	}

	return nil
}
