package busqueda

import "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"

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

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type ResponseBusinessAll struct {
	Error     bool                       `json:"error"`
	DataError string                     `json:"dataError"`
	Data      []models.Mo_Business_Cards `json:"data"`
}

type ResponseInterface struct {
	Error     bool          `json:"error"`
	DataError string        `json:"dataError"`
	Data      []interface{} `json:"data"`
}

type BusinessAll struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
