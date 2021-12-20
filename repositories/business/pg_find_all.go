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
	service_pg, typefood_pg, payment_pg := []int{}, []int{}, []int{}

	//counter-service
	service_counter := 0
	for _, s := range services {
		service_counter = service_counter + s
		service_pg = append(service_pg, s)
	}
	//counter-typefood
	typefood_counter := 0
	for _, t := range typefood {
		typefood_counter = typefood_counter + t
		typefood_pg = append(typefood_pg, t)
	}
	//counter-payment
	payment_counter := 0
	for _, p := range payment {
		payment_counter = payment_counter + p
		payment_pg = append(payment_pg, p)
	}

	//Agregamos un contador para la consulta
	counter := 0
	if service_counter > 0 {
		counter = counter + 1
	}
	if typefood_counter > 0 {
		counter = counter + 10
	}
	if payment_counter > 0 {
		counter = counter + 20
	}

	//Buscamos la consulta ideal
	switch counter {
	case 1:
		q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bs.idservice IN (select * from unnest($3::int[])) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($4, $5)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($6, $7)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, service_pg, latitude, longitude, latitude, longitude)
	case 10:
		q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bt.idtypefood IN (select * from unnest($3::int[])) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($4, $5)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($6, $7)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, typefood_pg, latitude, longitude, latitude, longitude)
	case 20:
		q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bp.idpayment IN (select * from unnest($3::int[])) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($4, $5)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($6, $7)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, payment_pg, latitude, longitude, latitude, longitude)
	case 11:
		q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bs.idservice IN (select * from unnest($3::int[])) AND bt.idtypefood IN (select * from unnest($4::int[])) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($5, $6)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($7, $8)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, services, typefood, latitude, longitude, latitude, longitude)
	case 21:
		q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bp.idpayment IN (select * from unnest($3::int[])) AND bs.idservice IN (select * from unnest($4::int[])) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($5, $6)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($7, $8)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, payment, services, latitude, longitude, latitude, longitude)
	case 31:
		q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bp.idpayment IN (select * from unnest($3::int[])) AND bt.idtypefood IN (select * from unnest($4::int[])) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($5, $6)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($7, $8)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, payment, typefood, latitude, longitude, latitude, longitude)
	default:
		q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness WHERE earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)) ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($3, $4)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, latitude, longitude)
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

	query := `INSERT INTO Near(idcomensal,nearbusiness) (select * from unnest($1::int[], $2::json[]))`
	if _, err := db.Exec(context.Background(), query, idcomensales, business); err != nil {
		return err
	}

	return nil
}
