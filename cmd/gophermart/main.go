package main

import (
	"database/sql"
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

	flag.StringVar(&dbStringConnect, "d", "", "DATABASE_DSN")
	flag.StringVar(&serverAddress, "a", "", "SERVER_ADDRESS")
	flag.StringVar(&baseURL, "b", "", "BASE_URL")
	flag.StringVar(&fileStoragePath, "f", "", "FILE_STORAGE_PATH")
	flag.Parse()

	fmt.Println(serverAddress)
	fmt.Println(baseURL)
	fmt.Println(fileStoragePath)
	fmt.Println(dbStringConnect)

	fmt.Println(`+++++++++++++`)

	var DB1 *sql.DB
	var err error
	DB1, err = sql.Open("postgres", dbStringConnect)
	if err != nil {
		fmt.Println(`StaticTest - err`)
	}
	fmt.Println(DB1)

	fmt.Println(`+++++++++++++`)
	httpserver.MainServer()

}
