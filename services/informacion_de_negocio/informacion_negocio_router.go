package informacion

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	//REPOSITORIES
	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

var InformationRouter_pg *informationRouter_pg

type informationRouter_pg struct {
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://c-registro-authenticacion.restoner-api.fun:80/v1/trylogin?jwt=" + jwt)
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
		log.Println(error_r)
	}
}

/*----------------------UPDATE MQTT DATA----------------------*/

func (cr *informationRouter_pg) Manual_CreateBusiness(c echo.Context) error {

	var createbusiness models.Mqtt_CreateInitialData

	err := c.Bind(&createbusiness)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el payment del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	error_r := CreateBusiness_Service(createbusiness)
	if error_r != nil {
		return nil
	}
	return nil
}

func (cr *informationRouter_pg) Manual_UpdatePaymenth(c echo.Context) error {

	var inputserialize_payment models.Mqtt_PaymentMethod

	err := c.Bind(&inputserialize_payment)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el payment del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	error_r := UpdatePaymenth_Service(inputserialize_payment)
	if error_r != nil {
		log.Println("ERROR UPDATE PAYMENT --> ", error_r)
		return nil
	}

	return nil
}

func (cr *informationRouter_pg) Manual_UpdateSchedule(c echo.Context) error {

	var inputserialize_schedule models.Mqtt_Schedule

	err := c.Bind(&inputserialize_schedule)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el schedule del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	error_r := UpdateSchedule_Service(inputserialize_schedule)
	if error_r != nil {
		log.Println("ERROR UPDATE SCHEDULE --> ", error_r)
		return nil
	}

	return nil
}

func (cr *informationRouter_pg) Manual_UpdateService(c echo.Context) error {

	var inputserialize_service models.Mqtt_Service

	err := c.Bind(&inputserialize_service)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el service del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	error_r := UpdateService_Service(inputserialize_service)
	if error_r != nil {
		log.Println("ERROR UPDATE SERVICE --> ", error_r)
		return nil
	}

	return nil
}

func (cr *informationRouter_pg) Manual_UpdateTypeFood(c echo.Context) error {

	var inputserialize_typefood models.Mqtt_TypeFood

	err := c.Bind(&inputserialize_typefood)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el typefood del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	error_r := UpdateTypeFood_Service(inputserialize_typefood)
	if error_r != nil {
		log.Println("ERROR UPDATE NAME --> ", error_r)
		return nil
	}

	return nil

}

func (cr *informationRouter_pg) Manual_UpdateName(c echo.Context) error {

	var inputserialize_name models.Mqtt_Name

	err := c.Bind(&inputserialize_name)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el name del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	error_r := UpdateName_Service(inputserialize_name)
	if error_r != nil {
		log.Println("ERROR UPDATE NAME --> ", error_r)
		return nil
	}

	return nil

}

func (cr *informationRouter_pg) UpdateLegalIdentity(inputserialize_legalidentity_multiple []models.Mqtt_LegalIdentity) {
	//Enviamos los datos al servicio
	error_r := UpdateLegalIdentity_Service(inputserialize_legalidentity_multiple)
	if error_r != nil {
		log.Println(error_r)
	}
}

func (cr *informationRouter_pg) Manual_UpdateUniqueName(c echo.Context) error {

	var inputserialize_uniquename models.Mqtt_Uniquename

	err := c.Bind(&inputserialize_uniquename)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el uniquename del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	error_r := UpdateUniqueName_Service(inputserialize_uniquename)
	if error_r != nil {
		log.Println("ERROR UPDATE UNIQUENAME --> ", error_r)
		return nil
	}

	return nil
}

func (cr *informationRouter_pg) Manual_UpdateTimeZone(c echo.Context) error {

	var inputserialize_timezone models.Mqtt_TimeZone

	err := c.Bind(&inputserialize_timezone)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el timezone del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	error_r := UpdateTimeZone_Service(inputserialize_timezone)
	if error_r != nil {
		log.Println("ERROR UPDATE TIMEZONE --> ", error_r)
		return nil
	}

	return nil
}

func (cr *informationRouter_pg) Manual_UpdateAddress(c echo.Context) error {

	var inputserialize_address models.Mqtt_Addres

	err := c.Bind(&inputserialize_address)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el address del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	error_r := UpdateAddress_Service(inputserialize_address)
	if error_r != nil {
		log.Println("ERROR UPDATE ADDRESS --> ", error_r)
		return nil
	}

	return nil
}

/*-----------------------------------------------------------------------------------------------------*/

func (cr *informationRouter_pg) Manual_UpdateBanner(c echo.Context) error {

	var inputserialize_banner models.Mqtt_Banner_Cola

	err := c.Bind(&inputserialize_banner)
	if err != nil {
		results := Response{Error: true, DataError: "Se debe enviar el banner del negocio, revise la estructura o los valores", Data: ""}
		return c.JSON(400, results)
	}

	//Enviamos los datos al servicio
	error_r := UpdateBanner_Service(inputserialize_banner)
	if error_r != nil {
		log.Println("ERROR UPDATE BANNER --> ", error_r)
		return nil
	}

	return nil
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
