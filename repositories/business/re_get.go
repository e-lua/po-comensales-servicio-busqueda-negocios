package repositories

import (
	"encoding/json"
	"math/rand"
	"strconv"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_BasicData_Business(idbusiness int) (models.Re_SetGetCode, error) {

	var basicdata models.Re_SetGetCode
	var reply string
	var err error

	random := rand.Intn(4)
	if random%2 == 0 {
		reply, err = redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(idbusiness)))
	} else {
		reply, err = redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(idbusiness)))
	}

	if err != nil {
		return basicdata, err
	}

	err = json.Unmarshal([]byte(reply), &basicdata)

	if err != nil {
		return basicdata, err
	}

	return basicdata, nil
}
