package app

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Database interface {
	CreateItem(*Item) error
	DeleteItem(int) error
	UpdateItem(int, *Item) error
	ListItems() ([]*Item, error)
	ListItemById(int) (*Item, error)
}

type PostgresDB struct {
	db *sql.DB
}

func PostgresConnection() (*PostgresDB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred loading on .env")
	}

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("port=%d user=%s dbname=%s password=%s sslmode=disable", port, user, dbname, password)
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

func (d *PostgresDB) UpdateItem(id int, item *Item) error {
	_, err := d.db.Query("update item set name = $1 where id = $2", item.Name, id)
	return err
}

func (d *PostgresDB) DeleteItem(id int) error {
	_, err := d.db.Query("delete from item where id = $1", id)
	return err
}

func (d *PostgresDB) ListItems() ([]*Item, error) {
	rows, err := d.db.Query("select * from item")
	if err != nil {
		return nil, err
	}

	items := []*Item{}
	for rows.Next() {
		item, err := scanIntoItem(rows)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (d *PostgresDB) ListItemById(id int) (*Item, error) {
	rows, err := d.db.Query("select * from item where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoItem(rows)
	}

	return nil, fmt.Errorf("item %d not found", id)
}

func scanIntoItem(rows *sql.Rows) (*Item, error) {
	item := new(Item)
	err := rows.Scan(
		&item.Id,
		&item.Name)

	return item, err
}
