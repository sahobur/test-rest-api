package main

import (
	"fmt"

	"awesomeProject/internal/app"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {
	var Accnt = &map[int64]float64{111: 0.0, 112: 100.1, 222: 300}
	err := app.ServerStart(fmt.Sprintf("%s:%s", HOST, PORT), Accnt)
	if err != nil {
		panic(err)
	}

}
