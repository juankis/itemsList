package models

type Item struct {
	tableName   struct{} `sql:"items"`
	Id          int      `sql:",pk"`
	Title       string
	Description string
	Picture     string
	Position    int
}
