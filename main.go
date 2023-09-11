package main

import (
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/jenniekibiri/crud-golang/configs"
	"github.com/jenniekibiri/crud-golang/migrate"

	"github.com/jenniekibiri/crud-golang/routes"
)

func init() {
	configs.LoadEnvs()
	configs.ConnectToDb()
	migrate.Migrate()
}


func main() {
	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	routes.PersonRouter(r)
	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server GO",

		})
	})
	r.Run()
	}

