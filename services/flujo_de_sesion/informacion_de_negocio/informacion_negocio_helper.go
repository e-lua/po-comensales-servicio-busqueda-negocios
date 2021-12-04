package informacion

import "github.com/Aphofisis/po-comensal-servicio-busqueda-negocios/models"

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

//BUSINESS FULL DATA

type ResponseCBusinessFullData struct {
	Error     bool                       `json:"error"`
	DataError string                     `json:"dataError"`
	Data      ResponseWithStructBusiness `json:"data"`
}

type ResponseWithStructBusiness struct {
	NameBusiness      string                            `json:"name"`
	LatitudeBusiness  float32                           `json:"latitude"`
	PostalCode        int                               `json:"postalCode"`
	LongitudeBusiness float32                           `json:"longitude"`
	Fulladdress       string                            `json:"fullAddress"`
	ReferenceAddress  string                            `json:"referenceAddress"`
	Banner            []models.Pg_BannerXBusiness       `json:"banner"`
	TypeOfFood        []models.Pg_TypeFoodXBusiness     `json:"typeOfFood"`
	Services          []models.Pg_ServiceXBusiness      `json:"services"`
	DeliveryRange     string                            `json:"deliveryRange"`
	PaymentMethods    []models.Pg_PaymenthMethXBusiness `json:"paymentMethods"`
	DailySchedule     []models.Pg_DayXBusiness          `json:"schedule"`
	PhoneContact      []models.Pg_ContactxBusiness      `json:"contact"`
}

type ResponseJWT struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      JWT    `json:"data"`
}

type JWT struct {
	Phone      int    `json:"phone"`
	Country    int    `json:"country"`
	IDComensal int    ` json:"comensal"`
	Name       string ` json:"name"`
	LastName   string ` json:"lastName"`
}
