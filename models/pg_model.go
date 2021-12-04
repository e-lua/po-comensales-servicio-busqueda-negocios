package models

import "time"

type Pg_DayXBusiness struct {
	IDDia      int    `json:"id"`
	Name       string `json:"name"`
	StarTime   string `json:"startTime"`
	EndTime    string `json:"endTime"`
	IsAvaiable bool   `json:"available"`
}

type Pg_ServiceXBusiness struct {
	IDService  int     `bson:"id" json:"id"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	Url        string  `json:"url"`
	TypeMoney  int     `json:"typeMoney"`
	IsAvaiable bool    `json:"available"`
}

type Pg_ContactxBusiness struct {
	IDContact   int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsAvaiable  bool   `json:"available"`
}

type Pg_PaymenthMethXBusiness struct {
	IDPaymenth  int    `bson:"id" json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Url         string `json:"url"`
	HasNumber   bool   `json:"hasNumber"`
	IsAvaiable  bool   `json:"available"`
}

type Pg_TypeFoodXBusiness struct {
	IDTypeFood int    `bson:"id" json:"id"`
	NameFood   string `json:"name"`
	URL        string `json:"url"`
	IsAvaiable bool   `json:"available"`
}

type Pg_BannerXBusiness struct {
	IDBanner int    `json:"id"`
	URL      string `json:"url"`
}

type Pg_Business struct {
	IDBusiness int
	Banner     string

	LegalIdentity    string
	LegalBusiness    bool
	Name             string
	ReferenceAddress string
	Latitude         float32
	Longitude        float32
	IsOpen           bool
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
	IDDay int    `json:"id"`
	Name  string `json:"name"`
}

type Pg_R_Contact struct {
	IDContact int    `json:"id"`
	Name      string `json:"name"`
}

type Pg_R_Service struct {
	IdService int     `json:"id"`
	Name      string  `json:"name"`
	Url       string  `json:"url"`
	Price     float32 `json:"price"`
}

type Pg_R_PaymentMethod struct {
	IDPaymenth int    `json:"id"`
	Name       string `json:"name"`
	IdCountry  int    `json:"country"`
	Url        string `json:"url"`
	HasNumber  bool   `json:"hasNumber"`
}

type Pg_R_TypeFood struct {
	IdTypeFood int    `json:"id"`
	Name       string `json:"name"`
	IdCountry  int    `json:"country"`
	Url        string `json:"url"`
}

type Pg_R_Typebsslegalindt struct {
	IDTypebsslegalindt int    `json:"id"`
	Name               string `json:"name"`
	Length             int    `json:"length"`
}

type Pg_ToBanner_Mqtt struct {
	IdBusiness                int    `bson:"idBusiness" json:"idBusiness"`
	IdBanner_Category_Element int    `json:"idBCE"`
	IdType                    int    `json:"idType"`
	Url                       string `bson:"url" json:"url"`
}
