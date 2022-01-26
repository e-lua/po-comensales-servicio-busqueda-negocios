package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_SearchByName(name string) ([]models.Pg_Found_All_Business, int, error) {

	//variable para contar las respuestas
	rows_found := 0

	db := models.Conectar_Pg_DB()
	q := "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp),'schedule',CONCAT(bsch.starttime,' - ',bsch.endtime),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=false AND replace(lower(b.name), ' ', '') like replace(lower($1), ' ', '') GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY b.view DESC"
	rows, error_show := db.Query(context.Background(), q, name)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListaInterface []models.Pg_Found_All_Business

	if error_show != nil {
		return oListaInterface, rows_found, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var interfac models.Pg_Found_All_Business
		rows.Scan(&interfac)
		oListaInterface = append(oListaInterface, interfac)
		rows_found = rows_found + 1
	}

	//Si todo esta bien
	return oListaInterface, rows_found, nil
}

func Pg_SearchByName_Test(name string) ([]models.Pg_Found_All_Business, int, error) {

	//variable para contar las respuestas
	rows_found := 0

	db := models.Conectar_Pg_DB()
	q := "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp),'schedule',CONCAT(bsch.starttime,' - ',bsch.endtime),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=true AND replace(lower(b.name), ' ', '') like replace(lower($1), ' ', '') GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY b.view DESC"
	rows, error_show := db.Query(context.Background(), q, name)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListaInterface []models.Pg_Found_All_Business

	if error_show != nil {
		return oListaInterface, rows_found, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var interfac models.Pg_Found_All_Business
		rows.Scan(&interfac)
		rows_found = rows_found + 1
		oListaInterface = append(oListaInterface, interfac)
	}

	//Si todo esta bien
	return oListaInterface, rows_found, nil
}
