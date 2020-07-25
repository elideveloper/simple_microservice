package service

import (
	//"fmt"
	"errors"

	"github.com/elideveloper/simple_microservice/models"
	"github.com/elideveloper/simple_microservice/requests"
	"github.com/elideveloper/simple_microservice/responses"
)

type ExampleServer interface {
	Get(req *requests.Get) (*responses.Get, error)
	Add(req *requests.Add) (*responses.Add, error)
	Delete(req *requests.Delete) (*responses.Delete, error)
	Update(req *requests.Update) (*responses.Update, error)
}

type exampleService struct {
	usersMap map[int]*models.User
	incID    int
}

func NewService() ExampleServer {
	return &exampleService{make(map[int]*models.User), 1}
}

func (svc *exampleService) Add(req *requests.Add) (*responses.Add, error) {
	svc.usersMap[svc.incID] = &req.User
	svc.incID++
	return &responses.Add{true}, nil
}

func (svc exampleService) Get(req *requests.Get) (*responses.Get, error) {
	if p, exists := svc.usersMap[req.ID]; exists {
		return &responses.Get{*p}, nil
	} else {
		return nil, errors.New("No such user.")
	}
}

func (svc exampleService) Update(req *requests.Update) (*responses.Update, error) {
	if _, exists := svc.usersMap[req.ID]; exists {
		svc.usersMap[req.ID] = &req.User
	} else {
		return nil, errors.New("No such user.")
	}
	return &responses.Update{true}, nil
}

func (svc *exampleService) Delete(req *requests.Delete) (*responses.Delete, error) {
	if _, exists := svc.usersMap[req.ID]; exists {
		delete(svc.usersMap, req.ID)
	} else {
		return nil, errors.New("No such user.")
	}
	return &responses.Delete{true}, nil
}
