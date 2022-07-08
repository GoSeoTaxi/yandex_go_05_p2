package main

import (
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05_p2/internal/httpserver"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	vars := os.Environ()

	fmt.Println(`+++++++++++++1`)
	fmt.Println(vars)
	fmt.Println(`+++++++++++++2`)

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	fmt.Println(`+++++++++++++`)

	httpserver.MainServer()

}
