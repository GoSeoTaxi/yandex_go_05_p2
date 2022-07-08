package main

import (
	"flag"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05_p2/internal/httpserver"
	"os"
)

func main() {

	vars := os.Environ()

	fmt.Println(`+++++++++++++1`)
	fmt.Println(vars)
	fmt.Println(`+++++++++++++2`)

	var serverAddress string
	var baseURL string
	var fileStoragePath string
	var dbStringConnect string

	flag.StringVar(&dbStringConnect, "d", "user=postgres password=qwerty dbname=goyp sslmode=disable", "DATABASE_DSN")
	flag.StringVar(&serverAddress, "a", "localhost:8080", "SERVER_ADDRESS")
	flag.StringVar(&baseURL, "b", ":8080", "BASE_URL")
	flag.StringVar(&fileStoragePath, "f", "", "FILE_STORAGE_PATH")
	flag.Parse()

	fmt.Println(`+++++++++++++`)

	httpserver.MainServer()

}
