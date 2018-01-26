package controllers

import (
	//	"database/sql"
	"github.com/moolica/IEProject/app/models"
	//"github.com/moolica/IEProject/app/routes"

	"github.com/revel/revel"
	"time"
)

type Application struct {
	GorpController
}

func (c Application) Download() revel.Result {
	return c.Render()
}

func (c Application) Submit() revel.Result {
	var movie models.Movie
	c.Params.BindJSON(&movie)
	movie.CreatedAt = time.Now().UnixNano()

	if err := c.Txn.Insert(&movie); err != nil {
		return c.RenderText(
			"Error inserting record into database!")
	} else {
		return c.RenderJSON(movie)
	}
}

func (c Application) Index() revel.Result {
	count, err := c.Txn.SelectStr("select title from Movie where id=?", 2)
	if err != nil {
		return c.RenderText(err.Error())
	}
	return c.Render(count)
	//	id := 1
	/*	var movie *models.Movie
		_, err := c.Txn.Select(&movie,
			c.Db.SqlStatementBuilder.
				Select("Title").
				From("Movie").Where("Id = ?", 1))
		if movie == nil {
			return c.NotFound("Movie not found")
		}
		if err != nil {
			panic(err)
		}
		return c.Render(movie.Title)
	*/
}

func (c Application) GetMovieDetails(id int) revel.Result {
	movie := new(models.Movie)
	err := c.Txn.SelectOne(movie,
		`SELECT * FROM Movie WHERE id = ?`, id)
	if err != nil {
		return c.RenderText("Error.  Item probably doesn't exist.")
	}
	return c.RenderJSON(movie)

}

func (c Application) GetRecentMovies(limit int) revel.Result {
	movies, err := c.Txn.Select(models.Movie{},
		`SELECT * FROM Movie order by created_at desc limit ?`, limit)
	if err != nil {
		return c.RenderText(
			err.Error())
	}
	return c.RenderJSON(movies)
}

func (c Application) SubmitComment(id int) revel.Result {
	var comment models.Comment
	c.Params.BindJSON(&comment)
	comment.CreatedAt = time.Now().UnixNano()
	comment.MovieId = int64(id)
	if err := c.Txn.Insert(&comment); err != nil {
		return c.RenderText(
			"Error inserting record into database!")
	} else {
		return c.RenderJSON(comment)
	}
}

func (c Application) GetComments(id int) revel.Result {
	comments, err := c.Txn.Select(models.Comment{},
		`SELECT * FROM Comment WHERE movie_id = ? order by created_at desc`, id)
	if err != nil {
		return c.RenderText(
			err.Error())
	}
	return c.RenderJSON(comments)

}

func (c Application) getMovieById(id int) *models.Movie {
	m, err := c.Txn.Get(models.Movie{}, id)
	if err != nil {
		panic(err)
	}
	if m == nil {
		return nil
	}
	return m.(*models.Movie)
}

func (c Application) Search(title string) revel.Result {
	movies, err := c.Txn.Select(models.Movie{},
		`SELECT * FROM Movie where title = ? order by created_at desc`, title)
	if err != nil {
		return c.RenderText(
			err.Error())
	}
	return c.RenderJSON(movies)
}
