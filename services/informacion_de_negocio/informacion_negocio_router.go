package informacion

import (
	"log"

	"github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

var InformationRouter_pg *informationRouter_pg

type informationRouter_pg struct {
}

/*
func GetJWT(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://147.182.232.30:3000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IDComensal
}*/

func (cr *informationRouter_pg) RegisterPaymenth(inputserialize_payment models.Mqtt_PaymentMethod) {

	//Enviamos los datos al servicio
	error_r := RegisterPaymenth_Service(inputserialize_payment)
	if error_r != nil {
		log.Fatal(error_r)
	}
}
