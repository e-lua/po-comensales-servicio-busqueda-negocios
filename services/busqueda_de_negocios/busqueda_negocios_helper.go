package busqueda

import "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"

type ResponseJWT struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      JWT    `json:"data"`
}

type JWT struct {
	Phone      int `json:"phone"`
	Country    int `json:"country"`
	IDComensal int ` json:"comensal"`
}

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type ResponseIBusinessCards struct {
	Error     bool          `json:"error"`
	DataError string        `json:"dataError"`
	Data      []interface{} `json:"data"`
}

type ResponseIBusinessCards_SearchedBefore struct {
	Error     bool                         `json:"error"`
	DataError string                       `json:"dataError"`
	Data      BusinessCards_SearchedBefore `json:"data"`
}

type BusinessCards_SearchedBefore struct {
	Quantity int           `json:"quantity"`
	Business []interface{} `json:"business"`
}

type ResponseFilterTypeFoods struct {
	Error     bool                   `json:"error"`
	DataError string                 `json:"dataError"`
	Data      []models.Pg_R_TypeFood `json:"data"`
}

type ResponseFilterPayments struct {
	Error     bool                        `json:"error"`
	DataError string                      `json:"dataError"`
	Data      []models.Pg_R_PaymentMethod `json:"data"`
}

type SearchFilters struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Services  []int   `json:"services"`
	TypeFood  []int   `json:"typefoods"`
	Payment   []int   `json:"payments"`
}
