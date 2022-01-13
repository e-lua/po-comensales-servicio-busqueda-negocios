package repositories

import (
	"bytes"
	"encoding/json"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

func Pg_ExportView(idbusiness int, idcomensal int) error {

	//Serializamos el MQTT
	var serialize_view models.Mqtt_View_Information
	serialize_view.IDBusiness = idbusiness
	serialize_view.IDComensal = idcomensal
	serialize_view.Date = time.Now()

	//Comenzamos el envio al MQTT

	go func() {
		//Comienza el proceso de MQTT
		ch, error_conection := models.MqttCN.Channel()
		if error_conection != nil {
			log.Error(error_conection)
		}

		bytes, error_serializar := serialize_viewinformation(serialize_view)
		if error_serializar != nil {
			log.Error(error_serializar)
		}

		error_publish := ch.Publish("", "comensal/viewinformation", false, false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         bytes,
			})
		if error_publish != nil {
			log.Error(error_publish)
		}

	}()

	return nil
}

//SERIALIZADORA
func serialize_viewinformation(serialize_view models.Mqtt_View_Information) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(serialize_view)
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}
