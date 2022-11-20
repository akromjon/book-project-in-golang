package models

import (
	"book/config"
	"log"
	"net/url"
)

var table string = "books"

type Book struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func (b Book) All() interface{} {

	db := config.NewDatabaseConnection()

	rows, err := db.Connection.Query("SELECT * FROM " + table)

	db.Close()

	if err != nil {
		log.Panic(err)
	}

	books := []Book{}

	for rows.Next() {

		rows.Scan(&b.ID, &b.Name, &b.Author)

		books = append(books, b)

	}

	return books
}

func (b Book) Find(id string) interface{} {

	db := config.NewDatabaseConnection()

	rows, err := db.Connection.Query("SELECT * FROM "+table+" WHERE id=?", id)

	db.Close()

	if err != nil {
		log.Panic(err)
	}

	for rows.Next() {

		rows.Scan(&b.ID, &b.Name, &b.Author)

	}

	return b
}

func (b Book) Create(u url.Values) bool {

	db := config.NewDatabaseConnection()

	_, err := db.Connection.Query("INSERT INTO "+table+" (name,author) values(?,?) ", u["name"][0], u["author"][0])

	db.Close()

	if err != nil {

		log.Panic(err)

		return false
	}

	return true

}

func (b Book) Update(id string, u url.Values) bool {

	db := config.NewDatabaseConnection()

	_, err := db.Connection.Query("UPDATE "+table+" SET name=?, author=? WHERE id=?", u["name"][0], u["author"][0], id)

	db.Close()

	if err != nil {

		log.Panic(err)

		return false
	}

	return true

}

func (b Book) Delete(id string) bool {

	db := config.NewDatabaseConnection()

	_, err := db.Connection.Query("DELETE FROM "+table+" WHERE id=?", id)

	if err != nil {

		log.Panic(err)

		return false
	}

	db.Close()

	return true
}
