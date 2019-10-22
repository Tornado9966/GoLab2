package menu

import (
	"database/sql"
	//"fmt"
)

type Dish struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Price   float64  `json:"price"`
}

type OrderForm struct {
	Ids []int64 `json:"ids"`
	TableNum int64 `json:"table_number"`
}

type Order struct {
	Table int64
	Dishes []Dish
	GenSum float64
	GenSumWithoutVAT float64
	Tips float64
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListDishes() ([]*Dish, error) {
	rows, err := s.Db.Query("SELECT id, name, price FROM menu_table LIMIT 200")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Dish
	for rows.Next() {
		var d Dish
		if err := rows.Scan(&d.Id, &d.Name, &d.Price); err != nil {
			return nil, err
		}
		res = append(res, &d)
	}
	if res == nil {
		res = make([]*Dish, 0)
	}
	return res, nil
}

func (s *Store) CreateOrder(ids []int64, table_number int64) (*Order, error) {
	var res Order
	var dishes []Dish
	for _, value := range ids{
		dish, err := s.Db.Query("SELECT id, name, price FROM menu_table WHERE id=$1", value)
	if err != nil {
		return nil, err
	}

		var d Dish
		for dish.Next(){
		if err := dish.Scan(&d.Id, &d.Name, &d.Price); err != nil {
			return nil, err
		}
		dishes = append(dishes, d)
		res.GenSum += d.Price
	}
	}
	
	if dishes == nil {
		dishes = make([]Dish, 0)
	}
	res.Dishes = dishes;
	res.Table = table_number;
	res.GenSumWithoutVAT = 0.8 * res.GenSum
	res.Tips = 0.1 * res.GenSum
	
	return &res, nil
}
