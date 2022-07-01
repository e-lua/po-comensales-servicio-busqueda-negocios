package repositories

import (
	"encoding/json"
	"strconv"

	models "github.com/Aphofisis/po-comensales-servicio-busqueda-negocios/models"
)

func Re_Set_BasicData_Business(idbusiness int, basic_Data models.Pg_BasicData_ToBusiness) error {

	var basicdata models.Re_SetGetCode
	basicdata.Basic_Data = basic_Data
	basicdata.IdBusiness = idbusiness

	uJson, err_marshal := json.Marshal(basicdata.Basic_Data)
	if err_marshal != nil {
		return err_marshal
	}

	_, err_do := models.RedisCN.Get().Do("SET", strconv.Itoa(basicdata.IdBusiness), uJson, "EX", 2000)
	if err_do != nil {
		return err_do
	}

	return nil
}
