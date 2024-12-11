package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()

	r.GET("/data", getalbums)
	r.POST("/data", postalbums)
	// r.GET("/data/:name", specificdata)
	r.GET("/data/:name", matchingdata)
	r.Run(":4000")
}

var albums = []album{
	{Id: 1, Name: "tharun"},
	{Id: 2, Name: "varun"},
	{Id: 3, Name: "Arun"},
}

// Get request to get the data
func getalbums(cxt *gin.Context) {
	cxt.IndentedJSON(http.StatusOK, albums)
}

// Post request to post the data
func postalbums(cxt *gin.Context) {
	var newalbum album
	if err := cxt.BindJSON(&newalbum); err != nil {
		return
	}
	albums = append(albums, newalbum)
	cxt.IndentedJSON(http.StatusOK, newalbum)
}

// Get the specific data using name (it is return only first data of the json)
// func specificdata(cxt *gin.Context) {
// 	name := cxt.Param("name")

// 	for _, value := range albums {
// 		if value.Name == name {
// 			cxt.IndentedJSON(http.StatusOK, value)
// 			return
// 		}
// 	}
// 	cxt.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

// Get the specific data using name (it is return all data matching of the name in json)

func matchingdata(cxt *gin.Context) {
	name := cxt.Param("name")

	var matchingalbums []album

	for _, value := range albums {
		if value.Name == name {
			matchingalbums = append(matchingalbums, value)
		}
	}
	if len(matchingalbums) == 0 {
		cxt.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	} else {
		cxt.IndentedJSON(http.StatusOK, matchingalbums)
	}
}
