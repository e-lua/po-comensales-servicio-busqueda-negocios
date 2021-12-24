package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	busqueda "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/services/busqueda_de_negocios"
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
	go Consumer_Service()
	go Consumer_Typefood()
	go Consumer_Name()
	go Consumer_Banner()
	go Consumer_Address()
	go Consumer_Open()

	e.GET("/", index)
	//VERSION
	version_1 := e.Group("/v1")

	//V1 FROM V1 TO ...TO ENTITY BUSINESS
	router_business := version_1.Group("/business")
	router_business.GET("/cache", busqueda.BusquedaRouter.GetBusinessCards_SearchedBefore)
	router_business.GET("/open", busqueda.BusquedaRouter.GetBusinessCards_Open)
	router_business.GET("/search", busqueda.BusquedaRouter.GetBusinessCards)

	//V1 FROM V1 TO ...TO FILTERS
	router_filter := version_1.Group("/filter")
	router_filter.GET("/typefood", busqueda.BusquedaRouter.GetFilterTypeFoods)
	router_filter.GET("/payment", busqueda.BusquedaRouter.GetFilterPaymentMethods)

	//V1 FROM V1 TO ...TO FAVORITES
	router_comensal := version_1.Group("/comensal")
	router_comensal.GET("/favorite", busqueda.BusquedaRouter.GetFavorites)
	router_comensal.POST("/favorite/:idbusiness", busqueda.BusquedaRouter.AddFavorites)

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
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/paymenth", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStopPaymenth := make(chan bool)

	go func() {
		for d := range msgs {
			var paymenth models.Mqtt_PaymentMethod
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&paymenth)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			informacion.InformationRouter_pg.UpdatePaymenth(paymenth)

			time.Sleep(5 * time.Second)
		}
	}()

	<-noStopPaymenth
}

func Consumer_Service() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/service", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStopService := make(chan bool)

	go func() {
		for d := range msgs {
			var service models.Mqtt_Service
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&service)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			informacion.InformationRouter_pg.UpdateService(service)

			time.Sleep(5 * time.Second)
		}
	}()

	<-noStopService
}

func Consumer_Typefood() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/typefood", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStopTypeFood := make(chan bool)

	go func() {
		for d := range msgs {
			var typefood models.Mqtt_TypeFood
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&typefood)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			informacion.InformationRouter_pg.UpdateTypeFood(typefood)

			time.Sleep(5 * time.Second)
		}
	}()

	<-noStopTypeFood
}

func Consumer_Name() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/name", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStopName := make(chan bool)

	go func() {
		for d := range msgs {
			var name models.Mqtt_Name
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&name)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			informacion.InformationRouter_pg.UpdateName(name)

			time.Sleep(5 * time.Second)
		}
	}()

	<-noStopName
}

func Consumer_Banner() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/banner", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStopBanner := make(chan bool)

	go func() {
		for d := range msgs {
			var banner models.Mqtt_Banner_Cola
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&banner)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			informacion.InformationRouter_pg.UpdateBanner(banner)

			time.Sleep(5 * time.Second)
		}
	}()

	<-noStopBanner
}

func Consumer_Address() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/address", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStopAddress := make(chan bool)

	go func() {
		for d := range msgs {
			var address models.Mqtt_Addres
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&address)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			informacion.InformationRouter_pg.UpdateAddress(address)

			time.Sleep(5 * time.Second)
		}
	}()

	<-noStopAddress
}

func Consumer_Open() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/timezone", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStopOpen := make(chan bool)

	go func() {
		for d := range msgs {
			var timezone models.Mqtt_TimeZone
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&timezone)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			informacion.InformationRouter_pg.UpdateTimeZone(timezone)

			time.Sleep(5 * time.Second)
		}
	}()

	<-noStopOpen
}
