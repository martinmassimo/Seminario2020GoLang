package chat

import (
	"github.com/jmoiron/sqlx"
	"github.com/martinmassimo/Seminario2020GoLang/internal/config"
)

// House ...
type Houses struct {
	ID     int64
	name   string
	status string
	rooms  int
	price  float32
}

// Service ...
type Service interface {
	AddHouse(House) error
	FindByID(int) *House
	FindAll() []*House
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddHouse(m House) error {
	return nil
}

func (s service) FindByID(ID int) *House {
	return nil
}

func (s service) FindAll() []*Message {
	var list []*House
	if err := s.db.Select(&list, "SELECT * FROM houses"); err != nil {
		panic(err)
	}
	return list
}

// package main

// import "github.com/gin-gonic/gin"

// func main() {
// 	r := gin.Default()
// 	r.GET("/houses", getHousesHandler)
// 	r.POST("/houses", addHousesHandler)
// 	r.Run()
// }

// func getHousesHandler(c *gin.Context) {

// 	c.JSON(200, gin.H{
// 		"status": "ok",
// 	})
// }
// func addHousesHandler(c *gin.Context) {

// }
