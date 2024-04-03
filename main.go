package main

import (
	"context"
	"fmt"
	"github.com/Caaki/go-gin/handlers"
	"github.com/Caaki/go-gin/initializers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.InitializeRedis()
}

func main() {

	port := os.Getenv("PORT")
	router := gin.Default()

	handlers.PostHandler(router)
	handlers.UserHandlers(router)

	ping, err := initializers.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ping)

	//err = initializers.RedisClient.Set(context.Background(), "name", "Nikola", 0).Err()
	//if err != nil {
	//	fmt.Printf("Failed to set value in the redis instance %s", err.Error())
	//}
	//
	//val, err := initializers.RedisClient.Get(context.Background(), "name").Result()
	//if err != nil {
	//	fmt.Println("Error retrieving value", err.Error())
	//}

	//fmt.Println(val)
	err = router.Run(`localhost:` + port)
	if err != nil {
		log.Fatal(err)
	}

}
