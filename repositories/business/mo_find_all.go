package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Mo_Find_All(location models.Location) ([]models.Mo_Business_Cards, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN_Externo.Database("restoner_externo")
	col := db.Collection("business_cards")

	// Todos los negocios
	var budsiness_card_all []models.Mo_Business_Cards

	//Filtro
	filter := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry":    location,
				"$maxDistance": 100000,
			},
		},
	}

	//Asignamos los datos del cursor
	cursor, err_find_all := col.Find(ctx, filter)
	if err_find_all != nil {
		return budsiness_card_all, err_find_all
	}

	for cursor.Next(ctx) {
		var budsiness_card models.Mo_Business_Cards
		err := cursor.Decode(&budsiness_card)
		if err != nil {
			return budsiness_card_all, err
		}
		budsiness_card_all = append(budsiness_card_all, budsiness_card)
	}

	return budsiness_card_all, nil
}
