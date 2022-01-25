package informacion

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/labstack/echo/v4"
)

var InformationRouter_pg *informationRouter_pg

type informationRouter_pg struct {
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

/*----------------------CRATE BUSINESS - MQTT DATA----------------------*/

func (cr *informationRouter_pg) CreateBusiness(inputserialize_create models.Mqtt_CreateInitialData) {
	//Enviamos los datos al servicio
	error_r := CreateBusiness_Service(inputserialize_create)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

/*----------------------UPDATE MQTT DATA----------------------*/

func (cr *informationRouter_pg) UpdatePaymenth(inputserialize_payment models.Mqtt_PaymentMethod) {
	//Enviamos los datos al servicio
	error_r := UpdatePaymenth_Service(inputserialize_payment)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) UpdateSchedule(inputserialize_payment models.Mqtt_Schedule) {
	//Enviamos los datos al servicio
	error_r := UpdateSchedule_Service(inputserialize_payment)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) UpdateService(inputserialize_service models.Mqtt_Service) {
	//Enviamos los datos al servicio
	error_r := UpdateService_Service(inputserialize_service)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) UpdateTypeFood(inputserialize_typefood models.Mqtt_TypeFood) {
	//Enviamos los datos al servicio
	error_r := UpdateTypeFood_Service(inputserialize_typefood)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) UpdateName(inputserialize_name models.Mqtt_Name) {
	//Enviamos los datos al servicio
	error_r := UpdateName_Service(inputserialize_name)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) UpdateUniqueName(inputserialize_uniquename models.Mqtt_Uniquename) {
	//Enviamos los datos al servicio
	error_r := UpdateUniqueName_Service(inputserialize_uniquename)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) UpdateTimeZone(inputserialize_open models.Mqtt_TimeZone) {
	//Enviamos los datos al servicio
	error_r := UpdateTimeZone_Service(inputserialize_open)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) UpdateBanner(inputserialize_banner models.Mqtt_Banner_Cola) {
	//Enviamos los datos al servicio
	error_r := UpdateBanner_Service(inputserialize_banner)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) UpdateAddress(inputserialize_address models.Mqtt_Addres) {
	//Enviamos los datos al servicio
	error_r := UpdateAddress_Service(inputserialize_address)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *informationRouter_pg) GetAddress(c echo.Context) error {

	idbusiness_string := c.Request().URL.Query().Get("idbusiness")

	idbusiness_int, _ := strconv.Atoi(idbusiness_string)

	//Validamos los valores enviados
	if idbusiness_int < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := FindAddress_Service(idbusiness_int)
	results := ResponseAddress{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *informationRouter_pg) AddViewInformation(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id del negocio
	idbusiness := c.Param("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Validamos los valores enviados
	if idbusiness_int < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddViewInformation_Service(idbusiness_int, data_idcomensal)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
