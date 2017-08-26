package requests

import (
	"simple_microservice/models"
)

type Get struct {
	ID int `json:"person_id"`
}

type Add struct {
	models.Person
}

type Update struct {
	ID int
	models.Person
}

type Delete struct {
	ID int
}
