package main

import (
	"github.com/Caaki/go-gin/initializers"
	"github.com/Caaki/go-gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	initializers.DB.AutoMigrate(&models.Post{}, &models.User{})

}
