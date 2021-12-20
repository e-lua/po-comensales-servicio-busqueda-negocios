package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/jackc/pgx/v4"
)

func Pg_Find_Favorite(idcomensal int) ([]models.Pg_Found_All_Business, error) {

	db := models.Conectar_Pg_DB()

	//Instanciamos una query
	var q string
	var rows pgx.Rows
	var error_show error

	q = "SELECT json_build_object('idbusiness',b.idbusiness,'name',b.name,'banner',b.urlbanner,'latitude',b.latitude,'longitude',b.longitude,'isopen',b.isopen,'services',json_agg(DISTINCT bs),'typefoods',json_agg(DISTINCT bt),'paymentmethods',json_agg(DISTINCT bp)) FROM business AS b LEFT JOIN (select bse.idbusiness,bse.idservice from bussinessr_service as bse) AS bs ON bs.idbusiness=b.idbusiness LEFT JOIN ( select bte.idbusiness,bte.idtypefood,t.name from businessr_typefood as bte join r_typefood as t on bte.idtypefood=t.idtypefood order by t.name asc) AS bt ON bt.idbusiness=b.idbusiness LEFT JOIN ( select bpe.idbusiness,bpe.idpayment from business_r_paymenth as bpe) AS bp ON bs.idbusiness=b.idbusiness LEFT JOIN favorites AS fav ON fav.idbusiness=b.idbusiness WHERE fav.idcomensal=$1 GROUP BY b.idbusiness,b.name,b.urlbanner,b.latitude,b.longitude,b.isopen"
	rows, error_show = db.Query(context.Background(), q, idcomensal)

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
