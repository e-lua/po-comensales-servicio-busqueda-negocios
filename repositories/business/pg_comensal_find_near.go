package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Pg_Comensal_Find_Near_Searched(idcomensal int) (int, []models.Pg_Found_All_Business, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB_Comensal()
	q := "SELECT json_build_object('distance',near.distance,'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness LEFT JOIN near ON near.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness WHERE bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND near.idcomensal=$1 AND b.istest=false GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,near.distance,bsch.endtime,bsch.starttime,bsch.isavailable ORDER BY near.distance"
	rows, error_show := db.Query(ctx, q, idcomensal)

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
