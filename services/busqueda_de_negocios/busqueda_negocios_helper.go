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

type ResponseJsonPostgreSQL struct {
	Error     bool                     `json:"error"`
	DataError string                   `json:"dataError"`
	Data      []models.Json_Postgresql `json:"data"`
}
