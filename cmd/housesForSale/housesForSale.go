package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/martinmassimo/Seminario2020GoLang/internal/config"
	"github.com/martinmassimo/Seminario2020GoLang/internal/database"
	"github.com/martinmassimo/Seminario2020GoLang/internal/service/houses"
)

func main() {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg := config.LoadConfig(*configFile)

	db, err := database.NewDatabase(cfg)
	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := houses.New(db, cfg)
	httpService := houses.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS houses (
		id integer primary key autoincrement,
		name   varchar,
		status varchar,
		rooms  int,
		price  float);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	insertHouse := `INSERT INTO houses (name, status, rooms, price) VALUES (?,?,?,?)`
	name := "Las Ceibos"
	status := "For Sale"
	rooms := 1
	price := 70.500
	db.MustExec(insertHouse, name, status, rooms, price)
	
	name = "Las Ceibos"
	status = "For Sale"
	rooms = 1
	price = 70.500
	db.MustExec(insertHouse, name, status, rooms, price)


	name = "Los Alamos"
	status = "For Sale"
	rooms = 3
	price = 100.000
	db.MustExec(insertHouse, name, status, rooms, price)

	name = "Los Aroldos"
	status = "For Sale"
	rooms = 2
	price = 73.500
	db.MustExec(insertHouse, name, status, rooms, price)
	
	name = "Los Naranjos"
	status = "For Sale"
	rooms = 4
	price = 150.600
	db.MustExec(insertHouse, name, status, rooms, price)

	name = "Las Acacias"
	status = "For Sale"
	rooms = 5
	price = 200.00
	db.MustExec(insertHouse, name, status, rooms, price)
	return nil
}
