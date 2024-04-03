package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Caaki/go-gin/initializers"
	"github.com/Caaki/go-gin/models"
	"github.com/gin-gonic/gin"
	"time"
)

func PostHandler(c *gin.Engine) {

	v1 := c.Group("/posts")
	{
		v1.POST("", PostsCreate)
		v1.GET("", AllPosts)
		v1.GET("/:id", PostByIndex)
		v1.PUT("/update/:id", PostUpdate)
		v1.DELETE("/:id", PostDelete)
	}

}

func PostsCreate(c *gin.Context) {
	//TODO: OVAKO UZIMAMO VREDNOSTI IZ TELA ZAHTEVA

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//==========================================================

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func AllPosts(c *gin.Context) {

	token := (c.Request.Header["Authorization"])[0]

	fmt.Println(base64.StdEncoding.DecodeString(token))

	_, err := initializers.RedisClient.Get(context.Background(), token).Result()
	if err != nil {
		c.JSON(400, fmt.Sprintf("Token invalid or expiered!!!"))
		return
	}

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostByIndex(c *gin.Context) {

	var post models.Post
	id := c.Param("id")
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//======== ISTO RADE OBA ALI OVO GORE VRATI LEPO VREDNOST====
	//post.Title = body.Title
	//post.Body = body.Body
	//initializers.DB.Updates(post)

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostDelete(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"message": "Post was deleted",
	})
}

func UserHandlers(c *gin.Engine) {

	v1 := c.Group("/users")
	{
		v1.POST("/register", UserCreate)
		v1.DELETE("/:id", PostDelete)
		v1.POST("/login", LoginHandler)
	}

}

func UserCreate(c *gin.Context) {
	//TODO: OVAKO UZIMAMO VREDNOSTI IZ TELA ZAHTEVA

	var user struct {
		Username string
		Password string
		Role     string
	}
	c.Bind(&user)

	//==========================================================

	newUser := models.User{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}

	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"User": newUser,
	})
}

func LoginHandler(c *gin.Context) {

	var loginParams struct {
		Username string
		Password string
	}

	c.Bind(&loginParams)

	var user models.User
	initializers.DB.Where(
		"username=? AND password=?",
		loginParams.Username,
		loginParams.Password).Find(&user)

	fmt.Println(user)

	if user == (models.User{}) {
		c.JSON(400, "Wrong credentials!")
		return
	}

	jsonString, err := json.Marshal(user)
	if err != nil {
		c.JSON(500, "Failed to marshal data")
	}

	err = initializers.RedisClient.Set(
		context.Background(),
		base64.StdEncoding.EncodeToString(jsonString),
		"Active",
		time.Duration(15*time.Second)).Err()
	if err != nil {
		c.JSON(500, fmt.Sprintf("Failed to set value in the redis instance %s", err.Error()))
	}

	c.JSON(200, gin.H{
		"token": jsonString,
	})

}
