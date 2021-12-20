package models

type Pg_R_TypeFood struct {
	IDTypefood int    `json:"id"`
	Name       string `json:"name"`
	Url        string `json:"url"`
}

type Pg_R_PaymentMethod struct {
	IDPaymenth int    `json:"id"`
	Name       string `json:"name"`
	Url        string `json:"url"`
}

type Pg_Found_All_Business struct {
	Banner     string         `json:"banner"`
	IDBusiness int            `json:"idbusiness"`
	IsOpen     bool           `json:"isopen"`
	Latitude   float32        `json:"latitude"`
	Longitude  float32        `json:"longitude"`
	Name       string         `json:"name"`
	Services   []ServiceB     `json:"services"`
	TypeFood   []TypeFoodB    `json:"typefoods"`
	Payment    []PaymentMethB `json:"paymentmethods"`
}

type PaymentMethB struct {
	IDBusiness int `json:"idbusiness"`
	IDPayment  int `json:"idpayment"`
}

type ServiceB struct {
	IDBusiness int `json:"idbusiness"`
	IDService  int `json:"idservice"`
}

type TypeFoodB struct {
	IDBusiness int    `json:"idbusiness"`
	IDTypeFood int    `json:"idtypefood"`
	Name       string `json:"name"`
}
