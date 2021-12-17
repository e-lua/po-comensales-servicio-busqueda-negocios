package models

import "time"

type Pg_DayXBusiness struct {
	IDDia      int  `json:"idday"`
	IDBusiness bool `json:"idbusiness"`
}

type Pg_ServiceXBusiness struct {
	IDService  int  `json:"idservice"`
	IDBusiness bool `json:"idbusiness"`
}

type Pg_ContactxBusiness struct {
	IDContact  int  `json:"idcontact"`
	IDBusiness bool `json:"idbusiness"`
}

type Pg_PaymenthMethXBusiness struct {
	IDPaymenth int  `json:"idpaymenth"`
	IDBusiness bool `json:"idbusiness"`
}

type Pg_TypeFoodXBusiness struct {
	IDTypeFood int  `json:"idtypeFood"`
	IDBusiness bool `json:"idbusiness"`
}

type Pg_Business struct {
	IDBusiness int
	Banner     string
	Name       string
	Latitude   float32
	Longitude  float32
	IsOpen     bool
}

type Pg_BusinessWorker_Mqtt struct {
	IdBusiness  int       `json:"idBusiness"`
	IdWorker    int       `json:"idWorker"`
	IdCountry   int       `json:"country"`
	CodeRedis   int       `json:"code"`
	Name        string    `json:"name"`
	IdRol       int       `json:"rol"`
	LastName    string    `json:"lastName"`
	Phone       int       `json:"phone"`
	Password    string    `json:"password"`
	UpdatedDate time.Time `json:"updateDate"`
}

type Pg_R_Country struct {
	IDCountry int
	Name      string
}

type Pg_R_Day struct {
	IDDay     int    `json:"id"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
}

type Pg_R_Contact struct {
	IDContact int    `json:"id"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
}

type Pg_R_Service struct {
	IdService int     `json:"id"`
	Name      string  `json:"name"`
	Url       string  `json:"url"`
	Price     float32 `json:"price"`
	Available bool    `json:"available"`
}

type Pg_R_PaymentMethod struct {
	IDPaymenth int    `json:"id"`
	Name       string `json:"name"`
	IdCountry  int    `json:"country"`
	Url        string `json:"url"`
	HasNumber  bool   `json:"hasNumber"`
	Available  bool   `json:"available"`
}

type Pg_R_TypeFood struct {
	IdTypeFood int    `json:"id"`
	Name       string `json:"name"`
	IdCountry  int    `json:"country"`
	Url        string `json:"url"`
	Available  bool   `json:"available"`
}
