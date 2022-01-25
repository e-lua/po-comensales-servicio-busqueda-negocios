package exportar

import "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"

//BASICDATA
type ResponseBasicData struct {
	Error     bool                           `json:"error"`
	DataError string                         `json:"dataError"`
	Data      models.Pg_BasicData_ToBusiness `json:"data"`
}

//SCHEDULE
type ResponseSchedule struct {
	Error     bool                              `json:"error"`
	DataError string                            `json:"dataError"`
	Data      []models.Pg_R_Schedule_ToBusiness `json:"data"`
}

//PAYMENTH
type ResponsePayment struct {
	Error     bool                                   `json:"error"`
	DataError string                                 `json:"dataError"`
	Data      []models.Pg_R_PaymentMethod_ToBusiness `json:"data"`
}

//SERVICE
type ResponseService struct {
	Error     bool                             `json:"error"`
	DataError string                           `json:"dataError"`
	Data      []models.Pg_R_Service_ToBusiness `json:"data"`
}

//TYPEFOOD
type ResponseTypeFood struct {
	Error     bool                              `json:"error"`
	DataError string                            `json:"dataError"`
	Data      []models.Pg_R_TypeFood_ToBusiness `json:"data"`
}

//RECOVER-ALL
type ResponseRecoverAll struct {
	Error     bool                 `json:"error"`
	DataError string               `json:"dataError"`
	Data      []models.Mo_Business `json:"data"`
}

//RECOVER-ONE
type ResponseRecoverOne struct {
	Error     bool               `json:"error"`
	DataError string             `json:"dataError"`
	Data      models.Mo_Business `json:"data"`
}
