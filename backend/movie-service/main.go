package main

import (
	"io"
	"os"

	"github.com/bangbugs/go-microservice-1/movie-service/common"
	"github.com/bangbugs/go-microservice-1/movie-service/controllers"
	"github.com/bangbugs/go-microservice-1/movie-service/databases"

	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error
	err = common.LoadConfig("config/config.json")
	if err != nil {
		return err
	}

	// Initialize mongo database
	err = databases.Database.Init()
	if err != nil {
		return err
	}

	// Setting Gin Logger
	if common.Config.EnableGinFileLog {
		f, _ := os.Create("logd/gin.log")
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

	// Initialize server
	if m.initServer() != nil {
		return
	}

	defer databases.Database.Close()

	c := controllers.Movie{}

	v1 := m.router.Group("/api/v1")
	{
		v1.POST("/login", c.Login)
		v1.GET("/movies/list", c.ListMovies)

		// Protected Routes
		v1.Use(jwt.Auth(common.Config.JwtSecretPassword))
		v1.POST("/movies", c.AddMovie)
	}

	m.router.Run(common.Config.Port)
}
