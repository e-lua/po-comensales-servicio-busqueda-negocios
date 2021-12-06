package models

type Mqtt_Business struct {
	Name           string              `bson:"name" json:"name"`
	DeliveryRange  string              `bson:"deliveryrange" json:"deliveryrange"`
	IsOpen         bool                `bson:"isopen" json:"isopen"`
	Contact        []Mqtt_Contact      `bson:"contact" json:"contact"`
	DailySchedule  []Mqtt_Day          `bson:"schedule" json:"schedule"`
	Address        Mqtt_Address        `bson:"address" json:"address"`
	Banner         []Mqtt_Banner       `bson:"banners" json:"banners"`
	TypeOfFood     []Mqtt_TypeFood     `bson:"typeoffood" json:"typeoffood"`
	Services       []Mqtt_Service      `bson:"services" json:"services"`
	PaymentMethods []Mqtt_PaymenthMeth `bson:"paymentmethods" json:"paymentmethods"`
}

type Mqtt_Banner struct {
	IdBanner string `bson:"id" json:"id"`
	UrlImage string `bson:"url" json:"url"`
}

type Mqtt_Address struct {
	Latitude         float32 `bson:"latitude" json:"latitude"`
	Longitude        float32 `bson:"longitude" json:"longitude"`
	FullAddress      string  `bson:"fulladdress" json:"fulladdress"`
	PostalCode       int     `bson:"postalcode" json:"postalcode"`
	State            string  `bson:"state" json:"state"`
	City             string  `bson:"city" json:"city"`
	ReferenceAddress string  `bson:"referenceaddress" json:"referenceaddress"`
}

type Mqtt_Day struct {
	IDDia      int    `bson:"id" json:"id"`
	Name       string `bson:"name" json:"name"`
	StarTime   string `bson:"starttime" json:"starttime"`
	EndTime    string `bson:"endtime" json:"endtime"`
	IsAvaiable bool   `bson:"available" json:"available"`
}

type Mqtt_TypeFood struct {
	IDTypeFood int    `bson:"id" json:"id"`
	Name       string `bson:"name" json:"name"`
	UrlImage   string `bson:"url" json:"url"`
}

type Mqtt_Service struct {
	IDService  int     `bson:"id" json:"id"`
	Name       string  `bson:"name" json:"name"`
	Price      float32 `bson:"price" json:"price"`
	Url        string  `bson:"url" json:"url"`
	TypeMoney  int     `bson:"typemoney" json:"typemoney"`
	IsAvaiable bool    `bson:"available" json:"available"`
}

type Mqtt_PaymenthMeth struct {
	IDPaymenth  int    `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	PhoneNumber string `bson:"phonenumber" json:"phonenumber"`
	Url         string `bson:"url" json:"url"`
	HasNumber   bool   `bson:"hasnumber" json:"hasnumber"`
	IsAvaiable  bool   `bson:"available" json:"available"`
}

type Mqtt_Contact struct {
	IDContact   int    `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	DataContact string `bson:"data" json:"data"`
}

type Mqtt_PaymentMethod struct {
	IdBusiness     int    `json:"idbusiness"`
	Idbusiness_pg  []int  `json:"idbusiness_pg"`
	Idpaymenth_pg  []int  `json:"idpaymenth_pg"`
	Isavailable_pg []bool `json:"isavailable_pg"`
}
