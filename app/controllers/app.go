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
	//success, err := c.Txn.Delete(&models.Movie{Id: 1})
	//if err != nil || success == 0 {
	//	return c.RenderText("Failed to remove BidItem")
	//}
	//return c.RenderText("Deleted %v", 1)
	//	err := c.Txn.SelectOne(movie,
	//		`SELECT * FROM movie WHERE id = ?`, id)
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
