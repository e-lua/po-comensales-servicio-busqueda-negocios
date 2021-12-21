package busqueda

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/labstack/echo/v4"
)

var BusquedaRouter *busquedaRouter

type busquedaRouter struct {
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://143.110.145.136:3000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IDComensal
}

func GetJWT_Country(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://143.110.145.136:3000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.Country
}

/*----------------------INICIA EL ROUTER----------------------*/

func (br *busquedaRouter) GetBusinessCards_SearchedBefore(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetBusinessCards_SearchedBefore_Service(data_idcomensal)
	results := ResponseIBusinessCards_SearchedBefore{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (br *busquedaRouter) GetBusinessCards(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	var search_filters SearchFilters

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&search_filters)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar la longitud y latitud del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetBusinessCards_Service(search_filters, data_idcomensal)
	results := ResponseIBusinessCards{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (br *busquedaRouter) GetBusinessCards_Open(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	var search_filters SearchFilters

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&search_filters)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar la longitud y latitud del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetBusinessCards_Open_Service(search_filters, data_idcomensal)
	results := ResponseIBusinessCards{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (br *busquedaRouter) GetFavorites(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetFavorites_Service(data_idcomensal)
	results := ResponseIBusinessCards{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (br *busquedaRouter) GetInformationOneBusiness(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id del Business Owner
	idbusiness := c.Param("idbusiness")

	//Enviamos los datos al servicio de anfitriones para obtener los datos completos
	respuesta, _ := http.Get("http://137.184.74.10:5800/v1/business/comensal/bnss/" + idbusiness)
	var get_respuesta models.Mo_Business
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	return c.JSON(status, get_respuesta)
}

/*----------------------FILTROS----------------------*/

func (br *busquedaRouter) GetFilterTypeFoods(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcountry := GetJWT_Country(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetFilterTypeFoods_Service(data_idcountry)
	results := ResponseFilterTypeFoods{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (br *busquedaRouter) GetFilterPaymentMethods(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcountry := GetJWT_Country(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetFilterPaymentMethods_Service(data_idcountry)
	results := ResponseFilterPayments{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------AGREGAR DATOS----------------------*/

func (br *busquedaRouter) AddFavorites(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id del negocio
	idbusiness := c.Param("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddFavorites_Service(data_idcomensal, idbusiness_int)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}
