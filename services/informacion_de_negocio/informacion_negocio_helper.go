package informacion

import "time"

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
	NameBusiness      string  `json:"name"`
	LatitudeBusiness  float32 `json:"latitude"`
	PostalCode        int     `json:"postalCode"`
	LongitudeBusiness float32 `json:"longitude"`
	Fulladdress       string  `json:"fullAddress"`
	ReferenceAddress  string  `json:"referenceAddress"`
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

//ADDRESS
type ResponseAddress struct {
	Error     bool      `json:"error"`
	DataError string    `json:"dataError"`
	Data      B_Address `json:"data"`
}

type B_Address struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Send_View_Information struct {
	IDBusiness int       `bson:"idbusiness" json:"idbusiness"`
	IDComensal int       `bson:"idcomensal" json:"idcomensal"`
	Date       time.Time `bson:"date" json:"date"`
}
