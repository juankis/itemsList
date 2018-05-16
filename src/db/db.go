package db

import "github.com/go-pg/pg"

//Connect return the conexion with the data base
func Connect() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "root",
		Database: "example",
	})
	return db
}
