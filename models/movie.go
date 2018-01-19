package main

type Movie struct {
	Id            int64  `db:"id"`
	CreatedAt     int64  `db:"created_at"`
	Title         string `db:"title"`
	OriginalTitle string `db:"original_title"`
	Rate          uint   `db:"rate"`
	Year          int64  `db:"year"`
	Length        string `db:"length"`
	Language      string `db:"language"`
	Country       string `db:"country"`
	Description   string `db:"description"`
	Director      string `db:director"`
}
