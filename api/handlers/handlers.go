package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	informacion "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/services/informacion_de_negocio"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Consumidor-MQTT
	go Consumer_Paymenth()

	e.GET("/", index)
	//VERSION
	//version_1 := e.Group("/v1")

	//V1 FROM V1 TO ...TO ENTITY BUSINESS
	//router_business := version_1.Group("/business")
	//router_business.GET("/:idbusiness", informacion.InformationRouter_pg.GetInformationData_Pg)

	//Abrimos el puerto
	PORT := os.Getenv("PORT")
	//Si dice que existe PORT
	if PORT == "" {
		PORT = "6200"
	}

	//cors son los permisos que se le da a la API
	//para que sea accesibl esde cualquier lugar
	handler := cors.AllowAll().Handler(e)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

func index(c echo.Context) error {
	return c.JSON(401, "Acceso no autorizado")
}

func Consumer_Paymenth() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		defer ch.Close()
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/paymenth", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	go func() {
		for d := range msgs {
			var anfitrion models.Mqtt_PaymentMethod
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&anfitrion)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			informacion.InformationRouter_pg.RegisterPaymenth(anfitrion)

			time.Sleep(10 * time.Second)
		}
	}()

}
