package main

import (
	"io"
	"os"

	"github.com/bangbugs/go-microservice-1/user-service/common"
	"github.com/bangbugs/go-microservice-1/user-service/controllers"
	"github.com/bangbugs/go-microservice-1/user-service/databases"

	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error

	// Load config file
	err = common.LoadConfig("config/config.json")

	if err != nil {
		return err
	}

	//Initialize user database
	err = databases.Database.Init()
	if err != nil {
		return err
	}

	// Setting gin logger
	if common.Config.EnableGinFileLog {
		f, _ := os.Create("logs/gin.log")
		if common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	} else {
		if !common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter()
		}
	}

	m.router = gin.Default()
	return nil
}

func main() {
	m := Main{}

	// intialize serer
	if m.initServer() != nil {
		return
	}

	defer databases.Database.Close()

	c := controllers.User{}

	// Simple group: v1
	v1 := m.router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.POST("/auth", c.Authenticate)
		}

		user := v1.Group("/users")

		//API needs to be authenticated
		user.Use(jwt.Auth(common.Config.JwtSecretPassword))
		{
			user.POST("", c.AddUser)
			user.GET("/list", c.ListUsers)
			user.GET("/detail/:id", c.GetUserByID)
			user.GET("/", c.GetUserByParams)
			user.DELETE(":id", c.DeleteUserByID)
			user.PUT("", c.UpdateUser)
		}

	}
	m.router.Run(common.Config.Port)
}
