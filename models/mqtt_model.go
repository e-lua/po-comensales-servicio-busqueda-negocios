package models

import "time"

type Mqtt_PaymentMethod struct {
	IdBusiness     int      `json:"idbusiness"`
	Idbusiness_pg  []int    `json:"idbusiness_pg"`
	Idpaymenth_pg  []int    `json:"idpaymenth_pg"`
	PhoneNumber    []string `json:"phonenumber_pg"`
	Isavailable_pg []bool   `json:"isavailable_pg"`
}

type Mqtt_Schedule struct {
	IdBusiness     int      `json:"idbusiness"`
	Idbusiness_pg  []int    `json:"idbusiness_pg"`
	Idschedule_pg  []int    `json:"idschedule_pg"`
	Isavailable_pg []bool   `json:"isavailable_pg"`
	Starttime_pg   []string `json:"starttime_pg"`
	Endtime_pg     []string `json:"endtime_pg"`
}

type Mqtt_Service struct {
	IdBusiness     int       `json:"idbusiness"`
	Idbusiness_pg  []int     `json:"idbusiness_pg"`
	Idservice_pg   []int     `json:"idservice_pg"`
	Pricing_pg     []float32 `json:"pricing_pg"`
	TypeMoney_pg   []int     `json:"typemoney_pg"`
	Isavailable_pg []bool    `json:"isavailable_pg"`
}

type Mqtt_TypeFood struct {
	IdBusiness     int    `json:"idbusiness"`
	Idbusiness_pg  []int  `json:"idbusiness_pg"`
	Idtypefood_pg  []int  `json:"Idtypefood_pg"`
	Isavailable_pg []bool `json:"isavailable_pg"`
}

type Mqtt_Addres struct {
	IdBusiness int     `json:"idbusiness"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type Mqtt_TimeZone struct {
	IdBusiness int    `json:"idbusiness"`
	TimeZone   string `json:"timezone"`
}

type Mqtt_Name struct {
	IdBusiness int    `json:"idbusiness"`
	Name       string `json:"name"`
}

type Mqtt_Uniquename struct {
	IdBusiness int    `json:"idbusiness"`
	Uniquename string `json:"uniquename"`
}

type Mqtt_LegalIdentity struct {
	IdBusiness      int     `json:"idbusiness"`
	LegalIdentity   string  `json:"legalidentity"`
	IVA             float32 `json:"iva"`
	Typesuscription int     `json:"typesuscription"`
}

type Mqtt_Banner_Cola struct {
	IdBusiness                int    `bson:"idbusiness" json:"idbusiness"`
	IdBanner_Category_Element int    `json:"idbCE"`
	IdType                    int    `json:"idtype"`
	Url                       string `bson:"url" json:"url"`
}

type Mqtt_CreateInitialData struct {
	IDBusiness   int  `json:"idbusiness"`
	Country      int  `json:"country"`
	IsSubsidiary bool `json:"issubsidiary"`
	SubsidiaryOf int  `json:"subsidiaryof"`
}

type Mqtt_View_Information struct {
	IDBusiness int       `bson:"idbusiness" json:"idbusiness"`
	IDComensal int       `bson:"idcomensal" json:"idcomensal"`
	Date       time.Time `bson:"date" json:"date"`
}
