package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database interface {
	CreateItem(*Item) error
	DeleteItem(int) error
	UpdateItem(*Item) error
	ListItems() ([]*Item, error)
	ListItemById(int) (*Item, error)
}

type PostgresDB struct {
	db *sql.DB
}

func PostgresConnection() (*PostgresDB, error) {
	connStr := "user=postgres dbname=postgres password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{
		db: db,
	}, nil
}

func (d *PostgresDB) Init() error {
	return d.createItemTable()
}

func (d *PostgresDB) createItemTable() error {
	query := `create table if not exists item(
		id serial primary key,
		name varchar(50)
	)`

	_, err := d.db.Exec(query)
	return err
}

func (d *PostgresDB) CreateItem(item *Item) error {
	query := `INSERT INTO item (name) values($1)`
	resp, err := d.db.Query(query, item.Name)
	
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (d *PostgresDB) UpdateItem(*Item) error {
	return nil
}

func (d *PostgresDB) DeleteItem(id int) error {
	return nil
}

func (d *PostgresDB) ListItems()([]*Item, error) {
	rows, err := d.db.Query("SELECT * FROM item")
	if err != nil {
		return nil, err
	}

	items := []*Item{}
	for rows.Next() {
		item := new(Item)
		err := rows.Scan(
			&item.Id,
			&item.Name)
	
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (d *PostgresDB) ListItemById(id int)(*Item, error) {
	return nil, nil
}