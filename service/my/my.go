package my

import (
	"go-dependency-injection/service/all"
)

type MyService interface {
	Hello(name string) string
}

type MyServiceImpl struct {
	allService all.AllService
}

func New(allService all.AllService) *MyServiceImpl {
	return &MyServiceImpl{
		allService: allService,
	}
}

func (s *MyServiceImpl) Hello(name string) string {
	return s.allService.Welcome() + " " + name
}
