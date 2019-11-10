package my

import (
	"testing"

	"go-dependency-injection/service/all"
)

func TestHello(t *testing.T) {
	allService := all.New()
	myService := New(allService)

	msg := myService.Hello("Bret")
	expected := "Hello Bret"

	if msg != expected {
		t.Errorf("Actual %s; expected %s", msg, expected)
	}
}
