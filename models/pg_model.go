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
