package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/jackc/pgx/v4"
)

func Pg_Comensal_Find_All(latitude float64, longitude float64, services []int, typefood []int, payment []int, idcomensal int) ([]models.Pg_Found_All_Business, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Instanciamos una query
	var idcomensales []int
	var q string
	var rows pgx.Rows
	var error_show error
	cantidad := 0
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

	db := models.Conectar_Pg_DB_Comensal()

	//Buscamos la consulta
	switch counter {
	case 1:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<2000 AND  bs.idservice IN (select * from unnest($3::int[])) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=false GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($4, $5)),bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($6, $7)) ASC"
		rows, error_show = db.Query(ctx, q, latitude, longitude, service_pg, latitude, longitude, latitude, longitude)
	case 10:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<2000 AND  bt.idtypefood IN (select * from unnest($3::int[])) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=false GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($4, $5)),bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($6, $7)) ASC"
		rows, error_show = db.Query(ctx, q, latitude, longitude, typefood_pg, latitude, longitude, latitude, longitude)
	case 20:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<2000 AND  bp.idpayment IN (select * from unnest($3::int[])) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=false GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($4, $5)),bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($6, $7)) ASC"
		rows, error_show = db.Query(ctx, q, latitude, longitude, payment_pg, latitude, longitude, latitude, longitude)
	case 11:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<2000 AND  bs.idservice IN (select * from unnest($3::int[])) AND bt.idtypefood IN (select * from unnest($4::int[])) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=false GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($5, $6)),bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($7, $8)) ASC"
		rows, error_show = db.Query(ctx, q, latitude, longitude, services, typefood, latitude, longitude, latitude, longitude)
	case 21:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<2000 AND  bp.idpayment IN (select * from unnest($3::int[])) AND bs.idservice IN (select * from unnest($4::int[])) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=false GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($5, $6)),bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($7, $8)) ASC"
		rows, error_show = db.Query(ctx, q, latitude, longitude, payment, services, latitude, longitude, latitude, longitude)
	case 31:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<2000 AND  bp.idpayment IN (select * from unnest($3::int[])) AND bt.idtypefood IN (select * from unnest($4::int[])) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=false GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($5, $6)),bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($7, $8)) ASC"
		rows, error_show = db.Query(ctx, q, latitude, longitude, payment, typefood, latitude, longitude, latitude, longitude)
	default:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<2000 AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=false GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)),bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($3, $4)) ASC"
		rows, error_show = db.Query(ctx, q, latitude, longitude, latitude, longitude)
	}

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	oListaInterface := []models.Pg_Found_All_Business{}
	oListaBusiness := []int{}
	oListDistance := []float64{}

	if error_show != nil {
		return oListaInterface, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		interfac := models.Pg_Found_All_Business{}
		cantidad = cantidad + 1
		rows.Scan(&interfac)
		oListaBusiness = append(oListaBusiness, interfac.IDBusiness)
		oListDistance = append(oListDistance, interfac.Distance)
		oListaInterface = append(oListaInterface, interfac)
	}

	//Agregamos los id de comensal
	for i := 0; i < cantidad; i++ {
		idcomensales = append(idcomensales, idcomensal)
	}

	//Insertamos los datos en cache
	delete(idcomensal)
	error_insert := insertFoundBusiness(idcomensales, oListaBusiness, oListDistance)
	if error_insert != nil {
		return oListaInterface, error_insert
	}

	//Si todo esta bien
	return oListaInterface, nil
}

func delete(idcomensal int) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Cambio de Server y BD, ya que no se puede acceder al rol de superusuario para la busqueda por distancia
	db := models.Conectar_Pg_DB_Comensal()

	//Eliminamos los datos
	q := `DELETE FROM Near WHERE idcomensal=$1`
	_, err := db.Exec(ctx, q, idcomensal)
	if err != nil {
		return err
	}

	return nil
}

func insertFoundBusiness(idcomensales []int, idbusiness []int, distance []float64) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Cambio de Server y BD, ya que no se puede acceder al rol de superusuario para la busqueda por distancia
	db := models.Conectar_Pg_DB_Comensal()

	query := `INSERT INTO Near(idcomensal,idbusiness,distance) (select * from unnest($1::int[], $2::int[], $3::real[]))`
	if _, err := db.Exec(ctx, query, idcomensales, idbusiness, distance); err != nil {
		return err
	}

	return nil
}

func Pg_Comensal_Find_All_Test(latitude float64, longitude float64, services []int, typefood []int, payment []int, idcomensal int) ([]models.Pg_Found_All_Business, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	//Instanciamos una query
	var q string
	var rows pgx.Rows
	var error_show error
	cantidad := 0
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

	db := models.Conectar_Pg_DB_Comensal()

	//Buscamos la consulta
	switch counter {
	case 1:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<2000 AND  bs.idservice IN (select * from unnest($3::int[])) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=true GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($4, $5)),bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($6, $7)) ASC"
		rows, error_show = db.Query(ctx, q, latitude, longitude, service_pg, latitude, longitude, latitude, longitude)
	case 10:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<2000 AND  bt.idtypefood IN (select * from unnest($3::int[])) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=true GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($4, $5)),bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($6, $7)) ASC"
		rows, error_show = db.Query(ctx, q, latitude, longitude, typefood_pg, latitude, longitude, latitude, longitude)
	case 20:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<2000 AND  bp.idpayment IN (select * from unnest($3::int[])) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=true GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($4, $5)),bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($6, $7)) ASC"
		rows, error_show = db.Query(ctx, q, latitude, longitude, payment_pg, latitude, longitude, latitude, longitude)
	case 11:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<2000 AND  bs.idservice IN (select * from unnest($3::int[])) AND bt.idtypefood IN (select * from unnest($4::int[])) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=true GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($5, $6)),bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($7, $8)) ASC"
		rows, error_show = db.Query(ctx, q, latitude, longitude, services, typefood, latitude, longitude, latitude, longitude)
	case 21:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<2000 AND  bp.idpayment IN (select * from unnest($3::int[])) AND bs.idservice IN (select * from unnest($4::int[])) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=true GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($5, $6)),bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($7, $8)) ASC"
		rows, error_show = db.Query(ctx, q, latitude, longitude, payment, services, latitude, longitude, latitude, longitude)
	case 31:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<2000 AND  bp.idpayment IN (select * from unnest($3::int[])) AND bt.idtypefood IN (select * from unnest($4::int[])) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=true GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($5, $6)),bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($7, $8)) ASC"
		rows, error_show = db.Query(ctx, q, latitude, longitude, payment, typefood, latitude, longitude, latitude, longitude)
	default:
		q = "SELECT json_build_object('distance',earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2)),'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND  earth_distance(ll_to_earth(latitude, longitude), ll_to_earth($1, $2))<2000 AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=true GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($1, $2)),bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth($3, $4)) ASC"
		rows, error_show = db.Query(ctx, q, latitude, longitude, latitude, longitude)
	}

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	oListaInterface := []models.Pg_Found_All_Business{}

	if error_show != nil {
		return oListaInterface, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		interfac := models.Pg_Found_All_Business{}
		cantidad = cantidad + 1
		rows.Scan(&interfac)
		oListaInterface = append(oListaInterface, interfac)
	}

	//Si todo esta bien
	return oListaInterface, nil
}
