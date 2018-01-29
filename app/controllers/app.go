package controllers

import (
	//	"database/sql"
	"github.com/moolica/IEProject/app/models"
	//"github.com/moolica/IEProject/app/routes"

	"github.com/revel/revel"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"strconv"
	"time"
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

type Application struct {
	GorpController
}

func (c Application) Download() revel.Result {
	return c.Render()
}

func (c Application) Profile() revel.Result {
	return c.Render()
}

func (c Application) Submit() revel.Result {

	movie := new(models.Movie)
	err := c.Txn.SelectOne(&movie,
		`SELECT * FROM Movie order by id desc limit 1`)
	if err != nil {
		return c.RenderText(
			err.Error())
	}
	return c.RenderText(strconv.Itoa(int(movie.Id)))

	//	var movie models.Movie
	//	c.Params.BindJSON(&movie)
	//	movie.CreatedAt = time.Now().UnixNano()
	//
	//	if err := c.Txn.Insert(&movie); err != nil {
	//		return c.RenderText(
	//			"Error inserting record into database!")
	//	} else {
	//		return c.RenderJSON(movie)
	//	}
}

func (c Application) Index() revel.Result {
	movies, err := c.Txn.Select(models.Movie{},
		`SELECT * FROM Movie order by created_at desc limit ?`, 5)
	if err != nil {
		return c.RenderText(
			err.Error())
	}
	return c.Render(movies)
}

func (c Application) GetMovieDetails(id int) revel.Result {
	movie := new(models.Movie)
	err := c.Txn.SelectOne(movie,
		`SELECT * FROM Movie WHERE id = ?`, id)
	if err != nil {
		return c.RenderText("Error.  Item probably doesn't exist.")
	}
	return c.Render(movie)
}

func (c Application) GetRecentMovies(limit int) revel.Result {
	movies, err := c.Txn.Select(models.Movie{},
		`SELECT * FROM Movie order by created_at desc limit ?`, limit)
	if err != nil {
		return c.RenderText(
			err.Error())
	}
	return c.Render(movies)
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

func (c Application) UploadMovie() revel.Result {
	return c.Render()
}

func (c *Application) HandleUpload(movieId int, movieName string, movieLength int, movieYear int, movieCountry string, movieDetails string, movieDirector string, movieAuthor string, movieStars string, movieCategory string, movieCover []byte) revel.Result { //	var movie models.Movie
	//	movie.CreatedAt = time.Now().UnixNano()
	//
	//	if err := c.Txn.Insert(&movie); err != nil {
	//		return c.RenderText(
	//			"Error inserting record into database!")
	//	} else {
	//		return c.RenderJSON(movie)
	//	}

	movie := models.Movie{
		CreatedAt:   time.Now().UnixNano(),
		Title:       movieName,
		Year:        int64(movieYear),
		Length:      strconv.Itoa(movieLength),
		Director:    movieDirector,
		Description: movieDetails,
	}

	if err := c.Txn.Insert(&movie); err != nil {
		return c.RenderText(
			"Error inserting record into database!")
	}

	err := c.saveImage(movieCover)
	if err != nil {
		return c.RenderText("Rid")
	}
	return c.RenderText("Eyvallll")
}

func (c Application) Comments() revel.Result {
	return c.Render()
}

func (c Application) saveImage(img []byte) (err error) {
	movie := new(models.Movie)
	err = c.Txn.SelectOne(&movie,
		`SELECT * FROM Movie order by id desc limit 1`)

	id := strconv.Itoa(int(movie.Id))

	err = ioutil.WriteFile("/Users/moolica/workspace/go/src/github.com/moolica/IEProject/public/posters/"+id+".jpg", img, 0644)
	return
}
