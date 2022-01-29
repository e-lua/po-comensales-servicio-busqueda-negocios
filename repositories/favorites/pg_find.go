package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/jackc/pgx/v4"
)

func Pg_Find(idcomensal int) ([]models.Pg_Found_All_Business, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	//Instanciamos una query
	var q string
	var rows pgx.Rows
	var error_show error

	q = "SELECT json_build_object('distance',0,'idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',CASE WHEN now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))< concat(bsch.endtime,b.timezone)::time with time zone  AND now()::time at time zone CONCAT('UTC',(b.timezone::integer*-1)::varchar(3))> concat(bsch.starttime,b.timezone)::time with time zone AND bsch.isavailable=true THEN true ELSE false END,'services',json_agg(DISTINCT bs),'schedule',CONCAT(to_char(bsch.starttime::time,'HH12:MI AM'),' - ',to_char(bsch.endtime::time,'HH12:MI AM')),'view',b.view) FROM business AS b LEFT JOIN (select bse.idservice,bse.pricing,bse.typemoney from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness JOIN businessschedule as bsch ON b.idbusiness=bsch.idbusiness LEFT JOIN favorites AS fav ON fav.idbusiness=b.idbusiness WHERE fav.idcomensal=$1 AND bsch.idschedule=EXTRACT(ISODOW FROM (NOW()::timestamp at time zone CONCAT('UTC',(b.timezone::integer)::varchar(3)))) AND urlbanner<>'' AND bs.idbusiness IS NOT NULL AND bt.idtypefood IS NOT NULL AND bp.idpayment IS NOT NULL and bsch.idbusiness IS NOT NULL AND b.name IS NOT NULL AND b.istest=false GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,earth_distance(ll_to_earth(b.latitude, b.longitude), ll_to_earth(-14.084628, -75.728557)),bsch.endtime,bsch.starttime,bsch.isavailable"
	rows, error_show = db.Query(ctx, q, idcomensal)

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
