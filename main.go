package main

import (
	// "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Beer is the model for mongo
type Beer struct {
	ID          ID `json:"id" bson:"_id,omitempty"`
	Name        string
	Brewery     string
	Style       string
	Alcohol     int
	Description string
	URL         string
	Current     bool
	Up          int
	Down        int
}

func main() {
	lab := os.Getenv("MONGOLAB_URI")
	db := os.Getenv("DATABASE_NAME")

	session, err := mgo.Dial(lab)
	col := session.DB(db).C("beers")
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/public", "public")

	r.GET("/", func(c *gin.Context) {
		var beers []Beer
		var current Beer
		query := bson.M{"current": true}
		col.Find(nil).All(&beers)
		col.Find(query).One(&current)
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"beers":   beers,
			"current": current,
		})
	})

	r.POST("/addBeer", func(c *gin.Context) {
		name := c.PostForm("name")
		brewery := c.PostForm("brewery")
		style := c.PostForm("style")
		url := c.PostForm("url")
		description := c.PostForm("description")
		query := bson.M{"current": true}
		change := bson.M{"$set": bson.M{"current": false}}
		alcohol, _ := strconv.Atoi(c.PostForm("alcohol"))
		col.UpdateAll(query, change)
		err = col.Insert(&Beer{Name: name, Brewery: brewery, Style: style, Alcohol: alcohol, Description: description, URL: url, Up: 0, Down: 0, Current: true})
		if err != nil {
			panic(err)
		}
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.POST("/deleteBeer", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.POST("/editBeer", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.Run(":" + port)
}
