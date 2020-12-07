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
	FindByID(int) (Houses, error)
	FindAll() ([]*Houses, error)
	AddHouse(Houses) (int64, error)
	DeleteByID(int) (int64,error)
	SetSoldByID(int) (int64,error)
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddHouse(h Houses) (int64, error) {
	insertHouse := `INSERT INTO houses (name, status, rooms, price) VALUES (?,?,?,?)`
	name := h.Name
	status := h.Status
	rooms := h.Rooms
	price := h.Price
	id, err := s.db.MustExec(insertHouse, name, status, rooms, price).LastInsertId()
	return id, err
}

func (s service) FindByID(ID int) (Houses,error){	
	var house Houses
	houses := []Houses{}
    err := s.db.Select(&houses, "SELECT * FROM houses WHERE id=$1", ID)
	if err != nil {	return house, err}
	if (len(houses)>0){
		house = houses[0]
	}
	return house, err
}

func (s service) FindAll() ([]*Houses,error) {
	var list []*Houses
	err := s.db.Select(&list, "SELECT * FROM houses")
	return list, err
}

func (s service) DeleteByID(ID int) (int64, error) {
	q := "DELETE FROM houses WHERE id=?"
	stmt, err := s.db.Prepare(q)
	if err != nil {	return 0, err }
	res, err := stmt.Exec(strconv.Itoa(ID))
	if err != nil {	return 0,err }
	affect, err := res.RowsAffected()
	if err != nil {	return affect, err }
	fmt.Println(affect)
	return affect, err
}

func (s service) SetSoldByID(ID int)  (int64,error) {
	q := "UPDATE houses SET status='Sold' WHERE id=? AND status='For Sale'"
	stmt, err := s.db.Prepare(q)
	if err != nil {	return 0, err }
	res, err := stmt.Exec(strconv.Itoa(ID))
	if err != nil {	return 0, err }
	affect, err := res.RowsAffected()
	if err != nil {	return 0, err }
	fmt.Println(affect)
	return affect, err
}
