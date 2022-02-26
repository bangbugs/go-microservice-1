package controllers

import (
	"net/http"

	"github.com/bangbugs/go-microservice-1/user-service/common"
	"github.com/bangbugs/go-microservice-1/user-service/daos"
	"github.com/bangbugs/go-microservice-1/user-service/models"
	"github.com/bangbugs/go-microservice-1/user-service/utils"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type User struct {
	utils   utils.Utils
	userDAO daos.User
}

func (u *User) Authenticate(ctx *gin.Context) {
	username := ctx.PostForm("user")
	password := ctx.PostForm("password")

	log.Debug("username: " + username + "password: " + password)

	var err error
	_, err = u.userDAO.Login(username, password)

	if err == nil {
		var tokensString string

		// generate token string
		tokensString, err := u.utils.GenerateJWT(username, "")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
			log.Debug("[ERROR]: ", err)
			return
		}

		token := models.Token{tokensString}
		ctx.JSON(http.StatusOK, token)
	} else {
		ctx.JSON(http.StatusUnauthorized, models.Error{common.StatusCodeUnknown, err.Error()})
	}
}

func (u *User) AddUser(ctx *gin.Context) {
	var addUser models.AddUser
	if err := ctx.ShouldBindJSON(&addUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	if err := addUser.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	user := models.User{bson.NewObjectId(), addUser.Name, addUser.Password}
	err := u.userDAO.Insert(user)
	if err == nil {
		ctx.JSON(http.StatusOK, models.Message{"Successfullt registered a user"})
		log.Debug("Registered a new user = " + user.Name + ", password = " + user.Password)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

func (u *User) ListUsers(ctx *gin.Context) {
	var users []models.User
	var err error
	users, err = u.userDAO.GetAll()

	if err == nil {
		ctx.JSON(http.StatusOK, users)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

func (u *User) GetUserByID(ctx *gin.Context) {
	var user models.User
	var err error

	id := ctx.Params.ByName("id")
	user, err = u.userDAO.GetByID(id)

	if err == nil {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

func (u *User) GetUserByParams(ctx *gin.Context) {
	var user models.User
	var err error

	id := ctx.Request.URL.Query()["id"][0]
	user, err = u.userDAO.GetByID(id)

	if err == nil {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

func (u *User) DeleteUserByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	err := u.userDAO.DeleteByID(id)

	if err == nil {
		ctx.JSON(http.StatusOK, models.Error{common.StatusCodeUnknown, err.Error()})
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

func (u *User) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	err := u.userDAO.Update(user)
	if err == nil {
		ctx.JSON(http.StatusOK, models.Message{"Successfully updated user"})
		log.Debug("Registered a new user = " + user.Name + ", password = " + user.Password)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}
