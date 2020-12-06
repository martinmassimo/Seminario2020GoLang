package houses

import (
	"github.com/jmoiron/sqlx"
	"github.com/martinmassimo/Seminario2020GoLang/internal/config"
	"strconv"
	"fmt"
)

// House ...
type Houses struct {
	Id     int     `db:"id" json:"Id"` 
	Name   string  `db:"name" json:"Name"`
	Status string  `db:"status" json:"Status"`
	Rooms  int     `db:"rooms" json:"Rooms"`
	Price  float32 `db:"price" json:"Price"`
}

// Service ...
type Service interface {
	FindByID(int) Houses
	FindAll() []*Houses
	AddHouse(Houses) error
	DeleteByID(int) error
	SetSoldByID(int) error
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddHouse(h Houses) error {
	insertHouse := `INSERT INTO houses (name, status, rooms, price) VALUES (?,?,?,?)`
	name := h.Name
	status := h.Status
	rooms := h.Rooms
	price := h.Price
	s.db.MustExec(insertHouse, name, status, rooms, price)
	return nil
}

func (s service) FindByID(ID int) Houses{	
	var house Houses
	houses := []Houses{}
    err := s.db.Select(&houses, "SELECT * FROM houses WHERE id=$1", ID)
	checkErr(err)
	house = houses[0]
	return house
}

func (s service) FindAll() []*Houses {
	var list []*Houses
	if err := s.db.Select(&list, "SELECT * FROM houses"); err != nil {
		panic(err)
	}
	return list
}

func (s service) DeleteByID(ID int) error {
	q := "DELETE FROM houses WHERE id=?"
	stmt, err := s.db.Prepare(q)
	checkErr(err)
	res, err := stmt.Exec(strconv.Itoa(ID))
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	return nil
}

func (s service) SetSoldByID(ID int) error {
	q := "UPDATE houses SET status='Sold' WHERE id=? AND status='For Sale'"
	stmt, err := s.db.Prepare(q)
	checkErr(err)
	res, err := stmt.Exec(strconv.Itoa(ID))
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}