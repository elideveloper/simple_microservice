package responses

import (
	"github.com/elideveloper/simple_microservice/models"
)

// universal format for responses
type General struct {
	Response interface{} `json:"response"`
	Err      string      `json:"error"`
}

type Get struct {
	models.User
}

type Add struct {
	Success bool `json:"success"`
}

type Update struct {
	Success bool `json:"success"`
}

type Delete struct {
	Success bool `json:"success"`
}
