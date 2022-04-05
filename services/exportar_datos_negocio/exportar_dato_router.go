package exportar

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

var ExportarRouter *exportarRouter

type exportarRouter struct {
}

func (er *exportarRouter) GetBasicData(c echo.Context) error {

	idbusiness_string := c.Request().URL.Query().Get("idbusiness")

	idbusiness, _ := strconv.Atoi(idbusiness_string)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetBasicData_Service(idbusiness)
	results := ResponseBasicData{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (er *exportarRouter) GetSchedule(c echo.Context) error {

	idbusiness_string := c.Request().URL.Query().Get("idbusiness")

	idbusiness, _ := strconv.Atoi(idbusiness_string)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetSchedule_Service(idbusiness)
	results := ResponseSchedule{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (er *exportarRouter) GetPayment(c echo.Context) error {

	idbusiness_string := c.Request().URL.Query().Get("idbusiness")
	country_string := c.Request().URL.Query().Get("country")

	idbusiness, _ := strconv.Atoi(idbusiness_string)
	country, _ := strconv.Atoi(country_string)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetPayment_Service(idbusiness, country)
	results := ResponsePayment{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (er *exportarRouter) GetService(c echo.Context) error {

	idbusiness_string := c.Request().URL.Query().Get("idbusiness")
	country_string := c.Request().URL.Query().Get("country")

	idbusiness, _ := strconv.Atoi(idbusiness_string)
	country, _ := strconv.Atoi(country_string)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetService_Service(idbusiness, country)
	results := ResponseService{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (er *exportarRouter) GetTypeFood(c echo.Context) error {

	idbusiness_string := c.Request().URL.Query().Get("idbusiness")
	country_string := c.Request().URL.Query().Get("country")

	idbusiness, _ := strconv.Atoi(idbusiness_string)
	country, _ := strconv.Atoi(country_string)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetTypeFood_Service(idbusiness, country)
	results := ResponseTypeFood{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*========================= RECUPERAR DATOS DEL NEGOCIO =========================*/

func (er *exportarRouter) GetRecoverAll(c echo.Context) error {

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetRecoverAll_Service()
	results := ResponseRecoverAll{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

func (er *exportarRouter) GetRecoverOne(c echo.Context) error {

	idbusiness_string := c.Request().URL.Query().Get("idbusiness")
	idbusiness, _ := strconv.Atoi(idbusiness_string)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetRecoverOne_Service(idbusiness)
	results := ResponseRecoverOne{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)

}

/*----------------------OBTENER TODOS LOS DATOS NEGOCIOS PARA NOTIFICARLOS----------------------*/

func (er *exportarRouter) SearchToNotify() {

	//Enviamos los datos al servicio
	status, _, dataerror, _ := SearchToNotify_Service()
	log.Println(strconv.Itoa(status) + " " + dataerror)
}
