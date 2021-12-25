package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Find_Near_Searched(idcomensal int) (int, []models.Pg_Found_All_Business, error) {

	db := models.Conectar_Pg_DB()
	q := "SELECT json_build_object('distance',near.distance,'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.available=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness LEFT JOIN near ON near.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM NOW()) AND near.idcomensal=$1 GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen,near.distance,bsch.endtime,bsch.starttime,bsch.available ORDER BY near.distance"
	rows, error_show := db.Query(context.Background(), q, idcomensal)

	rowsaffect := 0
	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListBusiness []models.Pg_Found_All_Business

	if error_show != nil {
		return rowsaffect, oListBusiness, error_show
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var interfac models.Pg_Found_All_Business
		rows.Scan(&interfac)
		rowsaffect = rowsaffect + 1
		oListBusiness = append(oListBusiness, interfac)
	}

	//Si todo esta bien
	return rowsaffect, oListBusiness, nil

}
