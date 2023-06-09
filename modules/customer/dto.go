package customer

import (
	"crmservice/dto"
	"crmservice/entities"
	"time"
)

type CustomerParam struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Avatar     string `json:"avatar"`
	Created_at time.Time
	Updated_at time.Time
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data CustomerParam `json:"data"`
}

type FindCustomer struct {
	dto.ResponseMeta
	Data entities.Customer `json:"data"`
}
