package informacion

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

var InformationRouter_pg *informationRouter_pg

type informationRouter_pg struct {
}

/*
func GetJWT(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://147.182.232.30:3000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IDComensal
}*/

func (ir *informationRouter_pg) GetInformationData_Pg(c echo.Context) error {

	//Obtenemos los datos del auth
	/*	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
		if dataerror != "" {
			results := Response{Error: boolerror, DataError: dataerror, Data: ""}
			return c.JSON(status, results)
		}
		if data_idcomensal <= 0 {
			results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
			return c.JSON(400, results)
		}*/

	//Recibimos el id del Business Owner
	idBusiness := c.Param("idbusiness")

	idBusiness_int, _ := strconv.Atoi(idBusiness)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetInformationData_Pg_Service(idBusiness_int)
	results := ResponseCBusinessFullData{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}
