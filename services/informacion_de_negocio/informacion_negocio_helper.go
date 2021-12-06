package informacion

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
