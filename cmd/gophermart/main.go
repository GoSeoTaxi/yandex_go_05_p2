package main

import (
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05_p2/internal/httpServer"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	fmt.Println(`+++++++++++++1`)
	fmt.Println(os.Getenv)
	fmt.Println(`+++++++++++++2`)

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	fmt.Println(`+++++++++++++`)

	httpServer.MainServer()

}
