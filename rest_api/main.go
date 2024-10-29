package main

import {
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
}

type todo struct {
	ID string  `json:"id"`
	Item string   `json:"title"`
	Completed bool   `json:"completed"`
}

todo := []todo{
	{ID: "1", Item: "enter the room", Completed: false},
	{ID: "2", Item: "clean the room", Completed: false},
	{ID: "3", Item: "view the room", Completed: false},
}

func getTodos(context *gin.Context){
    context.IndentedJSON(http.StatusOK, todos)
}

func main()  {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("127.0.0.1:8080")
}