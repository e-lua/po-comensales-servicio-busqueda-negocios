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
	respuesta, _ := http.Get("http://c-registro-authenticacion.restoner-api.fun:3000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IDComensal
}

func GetJWT_Country(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://c-registro-authenticacion.restoner-api.fun:3000/v1/trylogin?jwt=" + jwt)
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
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
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
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	latitude_string := c.Request().URL.Query().Get("latitude")
	longitude_string := c.Request().URL.Query().Get("longitude")
	services_string := c.Request().URL.Query().Get("services")
	typefoods_string := c.Request().URL.Query().Get("typefoods")
	payments_string := c.Request().URL.Query().Get("payments")

	//variable de array int
	var services []int
	var typefoods []int
	var payments []int

	//Convertimos los tipos de datos
	latitude, _ := strconv.ParseFloat(latitude_string, 64)
	longitude, _ := strconv.ParseFloat(longitude_string, 64)
	err_services := json.Unmarshal([]byte(services_string), &services)
	if err_services != nil {
		results := Response{Error: true, DataError: "Lista de servicios en formato incorrecto, detalles: " + err_services.Error(), Data: ""}
		return c.JSON(400, results)
	}
	err_typefoods := json.Unmarshal([]byte(typefoods_string), &typefoods)
	if err_typefoods != nil {
		results := Response{Error: true, DataError: "Lista de servicios en formato incorrecto, detalles: " + err_typefoods.Error(), Data: ""}
		return c.JSON(400, results)
	}
	err_payments := json.Unmarshal([]byte(payments_string), &payments)
	if err_payments != nil {
		results := Response{Error: true, DataError: "Lista de servicios en formato incorrecto, detalles: " + err_payments.Error(), Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetBusinessCards_Service(latitude, longitude, services, typefoods, payments, data_idcomensal)
	results := ResponseIBusinessCards{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (br *busquedaRouter) GetBusinessCardsByName(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Variable para indicar el tipo de busqueda
	tipo := 0

	//Recibimos el nombre
	name := c.Request().URL.Query().Get("name")

	if name[0] == 64 {
		tipo = tipo + 1
	} else {
		tipo = tipo + 2
	}

	name_string := "%" + name + "%"

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetBusinessCardsByName_Service(name_string, tipo)
	results := ResponseIBusinessCards{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (br *busquedaRouter) GetBusinessCards_Open(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	latitude_string := c.Request().URL.Query().Get("latitude")
	longitude_string := c.Request().URL.Query().Get("longitude")
	services_string := c.Request().URL.Query().Get("services")
	typefoods_string := c.Request().URL.Query().Get("typefoods")
	payments_string := c.Request().URL.Query().Get("payments")

	//variable de array int
	var services []int
	var typefoods []int
	var payments []int

	//Convertimos los tipos de datos
	latitude, _ := strconv.ParseFloat(latitude_string, 64)
	longitude, _ := strconv.ParseFloat(longitude_string, 64)
	err_services := json.Unmarshal([]byte(services_string), &services)
	if err_services != nil {
		results := Response{Error: true, DataError: "Lista de servicios en formato incorrecto, detalles: " + err_services.Error(), Data: ""}
		return c.JSON(400, results)
	}
	err_typefoods := json.Unmarshal([]byte(typefoods_string), &typefoods)
	if err_typefoods != nil {
		results := Response{Error: true, DataError: "Lista de servicios en formato incorrecto, detalles: " + err_typefoods.Error(), Data: ""}
		return c.JSON(400, results)
	}
	err_payments := json.Unmarshal([]byte(payments_string), &payments)
	if err_payments != nil {
		results := Response{Error: true, DataError: "Lista de servicios en formato incorrecto, detalles: " + err_payments.Error(), Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetBusinessCards_Open_Service(latitude, longitude, services, typefoods, payments)
	results := ResponseIBusinessCards{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (br *busquedaRouter) GetFavorites(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
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
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id del Business Owner
	idbusiness := c.Param("idbusiness")

	//Enviamos los datos al servicio de anfitriones para obtener los datos completos
	respuesta, _ := http.Get("http://a-informacion.restoner-api.fun:5800/v1/business/comensal/bnss/" + idbusiness)
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
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
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

func (br *busquedaRouter) GetUniqueNames(c echo.Context) error {

	uniquename_string := c.Request().URL.Query().Get("uniquename")
	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetUniqueNames_Service(uniquename_string)
	results := Response_Uniquename{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*=============================== INICIO TEST===============================*/

func (br *busquedaRouter) GetBusinessCards_Test(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	latitude_string := c.Request().URL.Query().Get("latitude")
	longitude_string := c.Request().URL.Query().Get("longitude")
	services_string := c.Request().URL.Query().Get("services")
	typefoods_string := c.Request().URL.Query().Get("typefoods")
	payments_string := c.Request().URL.Query().Get("payments")

	//variable de array int
	var services []int
	var typefoods []int
	var payments []int

	//Convertimos los tipos de datos
	latitude, _ := strconv.ParseFloat(latitude_string, 64)
	longitude, _ := strconv.ParseFloat(longitude_string, 64)
	err_services := json.Unmarshal([]byte(services_string), &services)
	if err_services != nil {
		results := Response{Error: true, DataError: "Lista de servicios en formato incorrecto, detalles: " + err_services.Error(), Data: ""}
		return c.JSON(400, results)
	}
	err_typefoods := json.Unmarshal([]byte(typefoods_string), &typefoods)
	if err_typefoods != nil {
		results := Response{Error: true, DataError: "Lista de servicios en formato incorrecto, detalles: " + err_typefoods.Error(), Data: ""}
		return c.JSON(400, results)
	}
	err_payments := json.Unmarshal([]byte(payments_string), &payments)
	if err_payments != nil {
		results := Response{Error: true, DataError: "Lista de servicios en formato incorrecto, detalles: " + err_payments.Error(), Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetBusinessCards_Test_Service(latitude, longitude, services, typefoods, payments, data_idcomensal)
	results := ResponseIBusinessCards{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (br *busquedaRouter) GetBusinessCardsByName_Test(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el nombre
	name := c.Request().URL.Query().Get("name")
	name_string := "%" + name + "%"

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetBusinessCardsByName_Test_Service(name_string)
	results := ResponseIBusinessCards{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*=============================== FIN TEST===============================*/
