package main

import (
	"fmt"

	"awesomeProject/internal/app"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

var Accnt = map[int64]float64{111: 0.0, 112: 100.1, 222: 300}

func main() {
	err := app.ServerStart(fmt.Sprintf("%s:%s", HOST, PORT))
	if err != nil {
		panic(err)
	}

}
