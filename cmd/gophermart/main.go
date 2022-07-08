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

	str := os.Getenv("DATABASE_URI")

	fmt.Println(`+++++++++++++1`)
	fmt.Println(vars)
	fmt.Println(`+++++++++++++2`)

	var serverAddress string
	var baseURL string
	var fileStoragePath string
	var dbStringConnect string

	flag.StringVar(&dbStringConnect, "d", "", "GOPHERMART-DATABASE-URI")
	flag.StringVar(&serverAddress, "a", "", "gophermart-database-uri")
	flag.StringVar(&baseURL, "b", "", "gophermart-port")
	flag.StringVar(&fileStoragePath, "f", "", "FILE_STORAGE_PATH")
	flag.Parse()

	fmt.Println(serverAddress)
	fmt.Println(baseURL)
	fmt.Println(fileStoragePath)
	fmt.Println(dbStringConnect)

	fmt.Println(`+++++++++++++`)
	fmt.Println(`str`)
	str1 := "postgresql://postgres:postgres@postgres/praktikum?sslmode=disable"
	fmt.Println(str1)
	fmt.Println(str)
	fmt.Println(`+++++++++++++`)
	var DB1 *sql.DB
	var err error
	DB1, err = sql.Open("postgres", str1)
	if err != nil {
		fmt.Println(`StaticTest - err`)
	}
	fmt.Println(DB1)

	fmt.Println(`+++++++++++++`)
	httpserver.MainServer()

}
