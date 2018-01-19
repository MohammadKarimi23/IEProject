package models

type Comment struct {
	Id          int64  `db:"id"`
	MovieId     int64  `db:"movie_id"`
	CreatedAt   int64  `db:"created_at"`
	Author      string `db:"author"`
	CommentText string `db:"comment_text"`
	Rate        uint   `db:"rate"`
}
