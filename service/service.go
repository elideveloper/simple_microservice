package service

import (
	//"fmt"
	"errors"
	"simple_microservice/models"
	"simple_microservice/requests"
	"simple_microservice/responses"
)

type ExampleServer interface {
	Get(req *requests.Get) (*responses.Get, error)
	Add(req *requests.Add) (*responses.Add, error)
	Delete(req *requests.Delete) (*responses.Delete, error)
	Update(req *requests.Update) (*responses.Update, error)
}

type exampleService struct {
	usersMap map[int]*models.Person
	incID    int
}

func NewService() ExampleServer {
	return &exampleService{make(map[int]*models.Person), 1}
}

func (svc *exampleService) Add(req *requests.Add) (*responses.Add, error) {
	svc.usersMap[svc.incID] = &req.Person
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
		svc.usersMap[req.ID] = &req.Person
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
