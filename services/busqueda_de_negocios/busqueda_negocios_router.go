package busqueda

import (
	"encoding/json"
	"net/http"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/labstack/echo/v4"
)

var BusquedaRouter_pg *busquedaRouter_pg

type busquedaRouter_pg struct {
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

/*----------------------INICIA EL ROUTER----------------------*/

func (br *busquedaRouter_pg) GetAllBusiness(c echo.Context) error {

	var input_location BusinessAll

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&input_location)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar la longitud y latitud del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	var location models.Location
	location.GeoJSONType = "Point"
	location.Coordinates = []float64{input_location.Latitude, input_location.Longitude}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetAllBusiness_Service(location)
	results := ResponseBusinessAll{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (br *busquedaRouter_pg) GetAllBusiness_2(c echo.Context) error {

	var input_location BusinessAll

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&input_location)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar la longitud y latitud del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	var location models.Location
	location.GeoJSONType = "Point"
	location.Coordinates = []float64{input_location.Latitude, input_location.Longitude}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetAllBusiness_2_Service(location)
	results := ResponseInterface{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (br *busquedaRouter_pg) GetInformationOneBusiness(c echo.Context) error {

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
	respuesta, _ := http.Get("http://137.184.74.10:3000/business/comensal/bnss/" + idbusiness)
	var get_respuesta models.ResponseBusiness
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	return c.JSON(status, get_respuesta)
}
