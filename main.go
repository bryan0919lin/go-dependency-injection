package main

import (
	"fmt"

	"go-dependency-injection/service/all"
	"go-dependency-injection/service/my"
)

func main() {
	allService := all.New()
	myService := my.New(allService)
	fmt.Println(myService.Hello("Bryan"))
}
