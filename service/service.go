package service

import (
	//"fmt"
	"errors"

	"github.com/elideveloper/simple_microservice/models"
	"github.com/elideveloper/simple_microservice/requests"
	"github.com/elideveloper/simple_microservice/responses"
)

var ErrUserNotFound error = errors.New("No such user")

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
	return &exampleService{make(map[int]*models.User), 0}
}

func (svc *exampleService) Add(req *requests.Add) (*responses.Add, error) {
	svc.incID++
	req.Id = svc.incID
	svc.usersMap[svc.incID] = &req.User

	return &responses.Add{req.User}, nil
}

func (svc exampleService) Get(req *requests.Get) (*responses.Get, error) {
	if p, exists := svc.usersMap[req.ID]; exists {
		return &responses.Get{*p}, nil
	}

	return nil, ErrUserNotFound
}

func (svc exampleService) Update(req *requests.Update) (*responses.Update, error) {
	if _, exists := svc.usersMap[req.ID]; exists {
		svc.usersMap[req.ID] = &req.User
	} else {
		return nil, ErrUserNotFound
	}
	return &responses.Update{true}, nil
}

func (svc *exampleService) Delete(req *requests.Delete) (*responses.Delete, error) {
	if _, exists := svc.usersMap[req.ID]; exists {
		delete(svc.usersMap, req.ID)
	} else {
		return nil, ErrUserNotFound
	}
	return &responses.Delete{true}, nil
}
