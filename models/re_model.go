package models

type Re_SetGetCode struct {
	IdBusiness int                     `json:"idbusiness"`
	Basic_Data Pg_BasicData_ToBusiness `json:"basicdata"`
}
