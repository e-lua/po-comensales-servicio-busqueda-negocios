package models

type Mqtt_PaymentMethod struct {
	IdBusiness     int    `json:"idbusiness"`
	Idbusiness_pg  []int  `json:"idbusiness_pg"`
	Idpaymenth_pg  []int  `json:"idpaymenth_pg"`
	Isavailable_pg []bool `json:"isavailable_pg"`
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

type Mqtt_IsOpen struct {
	IdBusiness int  `json:"idbusiness"`
	IsaOpen    bool `json:"isopen"`
}

type Mqtt_Name struct {
	IdBusiness int    `json:"idbusiness"`
	Name       string `json:"name"`
}

type Mqtt_Banner_Cola struct {
	IdBusiness                int    `bson:"idbusiness" json:"idbusiness"`
	IdBanner_Category_Element int    `json:"idbCE"`
	IdType                    int    `json:"idtype"`
	Url                       string `bson:"url" json:"url"`
}
