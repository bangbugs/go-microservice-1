package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/bangbugs/go-microservice-1/movie-service/common"
	"github.com/bangbugs/go-microservice-1/movie-service/daos"
	"github.com/bangbugs/go-microservice-1/movie-service/models"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Movie struct {
	movieDAO daos.Movie
}

func (m *Movie) Login(ctx *gin.Context) {
	username := ctx.PostForm("user")
	password := ctx.PostForm("password")

	formData := url.Values{
		"user":     {username},
		"password": {password},
	}

	var authAddr string = common.Config.AuthAddr + "/api/v1/admin/auth"
	resp, err := http.PostForm(authAddr, formData)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
		return
	}

	if resp.StatusCode == http.StatusOK {
		var token models.Token
		json.NewDecoder(resp.Body).Decode(&token)
		ctx.JSON(http.StatusOK, token)
	} else {
		var e models.Error
		json.NewDecoder(resp.Body).Decode(&e)
		ctx.JSON(resp.StatusCode, e)
	}
}

func (m *Movie) AddMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.BindJSON(&movie); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	movie.ID = bson.NewObjectId()
	err := m.movieDAO.Insert(movie)
	if err == nil {
		ctx.JSON(http.StatusOK, models.Message{Message: "Movie Created Successfully"})
	} else {
		ctx.JSON(http.StatusForbidden, models.Error{Code: common.StatusCodeUnknown, Message: err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

func (m *Movie) ListMovies(ctx *gin.Context) {
	var movies []models.Movie
	var err error

	movies, err = m.movieDAO.GetAll()

	if err == nil {
		ctx.JSON(http.StatusOK, movies)
	} else {
		ctx.JSON(http.StatusNotFound, models.Error{Code: common.StatusCodeUnknown, Message: "Cannot retrive movie info"})
		log.Debug("[ERROR]: ", err)
	}

}
