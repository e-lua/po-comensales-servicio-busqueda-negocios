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
	Distance   float64    `json:"distance"`
	Banner     string     `json:"banner"`
	IDBusiness int        `json:"idbusiness"`
	IsOpen     bool       `json:"isopen"`
	Latitude   float32    `json:"latitude"`
	Longitude  float32    `json:"longitude"`
	Name       string     `json:"name"`
	Schedule   string     `json:"schedule"`
	View       int        `json:"view"`
	Services   []ServiceB `json:"services"`
}

type PaymentMethB struct {
	IDBusiness int `json:"idbusiness"`
	IDPayment  int `json:"idpayment"`
}

type ServiceB struct {
	IDBusiness int     `json:"idbusiness"`
	IDService  int     `json:"idservice"`
	Price      float32 `json:"pricing"`
	TypeMoney  int     `json:typemoney`
}

type TypeFoodB struct {
	IDBusiness int    `json:"idbusiness"`
	IDTypeFood int    `json:"idtypefood"`
	Name       string `json:"name"`
}

/*--------------------------MODELOS A EXPORTAR---------------------------*/

type Pg_R_PaymentMethod_ToBusiness struct {
	IDPaymenth  int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	PhoneNumber string `bson:"phonenumber" json:"phonenumber"`
	HasNumber   bool   `json:"hasnumber"`
	IsAvailable bool   `json:"available"`
}

type Pg_R_Schedule_ToBusiness struct {
	IDSchedule int    `json:"idschedule"`
	Starttime  string `json:"starttime"`
	Endtime    string `json:"endtime"`
	Available  bool   `json:"available"`
}

type Pg_PaymentMethod_X_Business_ToBusiness struct {
	IDPaymenth  int
	IDBusiness  int
	IsAvailable bool
}

type Pg_R_Service_ToBusiness struct {
	IDservice   int     `json:"id"`
	Name        string  `json:"name"`
	Pricing     float32 `json:"price"`
	TypeMoney   int     `json:"typemoney"`
	IsAvailable bool    `json:"available"`
}

type Pg_R_TypeFood_ToBusiness struct {
	IDTypefood  int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	IsAvailable bool   `json:"available"`
}

type Pg_BasicData_ToBusiness struct {
	IsOpen   bool   `json:"isopen"`
	Name     string `json:"name"`
	TimeZone string `json:"timezone"`
}
