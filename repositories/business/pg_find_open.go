package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/jackc/pgx/v4"
)

func Pg_Find_Open(latitude float64, longitude float64, services []int, typefood []int, payment []int) ([]models.Pg_Found_All_Business, error) {

	db := models.Conectar_Pg_DB()

	//Instanciamos una query
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

	//Buscamos la consulta
	switch counter {
	case 1:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE (CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END)=true AND  bsch.idschedule=EXTRACT(ISODOW FROM NOW()) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bs.idservice IN (select * from unnest($3::int[])) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($4, $5)),bsch.endtime,bsch.starttime,bsch.available ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($6, $7)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, service_pg, latitude, longitude, latitude, longitude)
	case 10:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE (CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END)=true AND  bsch.idschedule=EXTRACT(ISODOW FROM NOW()) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bt.idtypefood IN (select * from unnest($3::int[])) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($4, $5)),bsch.endtime,bsch.starttime,bsch.available ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($6, $7)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, typefood_pg, latitude, longitude, latitude, longitude)
	case 20:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE (CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END)=true AND  bsch.idschedule=EXTRACT(ISODOW FROM NOW()) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bp.idpayment IN (select * from unnest($3::int[])) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($4, $5)),bsch.endtime,bsch.starttime,bsch.available ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($6, $7)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, payment_pg, latitude, longitude, latitude, longitude)
	case 11:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE (CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END)=true AND  bsch.idschedule=EXTRACT(ISODOW FROM NOW()) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bs.idservice IN (select * from unnest($3::int[])) AND bt.idtypefood IN (select * from unnest($4::int[])) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($5, $6)),bsch.endtime,bsch.starttime,bsch.available ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($7, $8)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, services, typefood, latitude, longitude, latitude, longitude)
	case 21:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE (CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END)=true AND  bsch.idschedule=EXTRACT(ISODOW FROM NOW()) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bp.idpayment IN (select * from unnest($3::int[])) AND bs.idservice IN (select * from unnest($4::int[])) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($5, $6)),bsch.endtime,bsch.starttime,bsch.available ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($7, $8)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, payment, services, latitude, longitude, latitude, longitude)
	case 31:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE (CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END)=true AND  bsch.idschedule=EXTRACT(ISODOW FROM NOW()) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 AND  bp.idpayment IN (select * from unnest($3::int[])) AND bt.idtypefood IN (select * from unnest($4::int[])) GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($5, $6)),bsch.endtime,bsch.starttime,bsch.available ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($7, $8)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, payment, typefood, latitude, longitude, latitude, longitude)
	default:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE (CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END)=true AND  bsch.idschedule=EXTRACT(ISODOW FROM NOW()) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<10000 GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)),bsch.endtime,bsch.starttime,bsch.available ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($3, $4)) ASC"
		rows, error_show = db.Query(context.Background(), q, latitude, longitude, latitude, longitude)
	}

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListaInterface []models.Pg_Found_All_Business

	if error_show != nil {
		return oListaInterface, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var interfac models.Pg_Found_All_Business
		rows.Scan(&interfac)
		oListaInterface = append(oListaInterface, interfac)
	}

	//Si todo esta bien
	return oListaInterface, nil
}
