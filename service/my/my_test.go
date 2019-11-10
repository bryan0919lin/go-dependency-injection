package my

import (
	"testing"
)

func TestHello(t *testing.T) {
	allService := NewAllServiceMock()
	myService := New(allService)

	msg := myService.Hello("Bret")
	expected := "Welcome Bret"

	if msg != expected {
		t.Errorf("Actual %s; expected %s", msg, expected)
	}
}
