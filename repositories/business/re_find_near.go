package repositories

import (
	"encoding/json"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_Near_Business(idcomensal int) (models.Re_SetGetCode, error) {

	var negocios models.Re_SetGetCode

	rpta_redis, err := redis.String(models.RedisCN.Get().Do("GET", idcomensal))

	if err != nil {
		return negocios, err
	}

	err_unmarshal := json.Unmarshal([]byte(rpta_redis), &negocios)
	if err != nil {
		return negocios, err_unmarshal
	}

	return negocios, nil
}
