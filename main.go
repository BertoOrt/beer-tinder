package main

import (
	// "encoding/json"
	// "fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Beer is the model for mongo
type Beer struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string
	Brewery     string
	Style       string
	Alcohol     float64
	Description string
	URL         string
	Current     bool
	Up          int
	Down        int
	Timestamp   int32
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

	r.GET("/current", func(c *gin.Context) {
		var current Beer
		query := bson.M{"current": true}
		col.Find(query).One(&current)
		c.JSON(200, gin.H{
			"data": current,
		})
	})

	r.POST("/new", func(c *gin.Context) {
		name := c.PostForm("name")
		brewery := c.PostForm("brewery")
		style := strings.Trim(c.PostForm("style"), " ")
		url := strings.Trim(c.PostForm("url"), " ")
		description := c.PostForm("description")
		query := bson.M{"current": true}
		change := bson.M{"$set": bson.M{"current": false}}
		alcohol, _ := strconv.ParseFloat(c.PostForm("alcohol"), 64)
		col.UpdateAll(query, change)
		err = col.Insert(&Beer{Name: name, Brewery: brewery, Style: style, Alcohol: alcohol,
			Description: description, URL: url, Up: 0, Down: 0, Current: true, Timestamp: int32(time.Now().Unix())})
		if err != nil {
			panic(err)
		}
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.POST("/delete/:id", func(c *gin.Context) {
		time, _ := strconv.Atoi(c.Param("id"))
		err := col.Remove(bson.M{"timestamp": time})
		if err != nil {
			panic(err)
		}
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.POST("/edit/:id", func(c *gin.Context) {
		name := c.PostForm("name")
		brewery := c.PostForm("brewery")
		style := strings.Trim(c.PostForm("style"), " ")
		url := strings.Trim(c.PostForm("url"), " ")
		description := c.PostForm("description")
		alcohol, _ := strconv.ParseFloat(c.PostForm("alcohol"), 64)
		time, _ := strconv.Atoi(c.Param("id"))
		update := bson.M{
			"name":        name,
			"brewery":     brewery,
			"style":       style,
			"alcohol":     alcohol,
			"description": description,
			"url":         url,
		}
		change := bson.M{"$set": update}
		col.Update(bson.M{"timestamp": time}, change)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.POST("/makeCurrent/:id", func(c *gin.Context) {
		time, _ := strconv.Atoi(c.Param("id"))
		query := bson.M{"current": true}
		change := bson.M{"$set": bson.M{"current": false}}
		col.UpdateAll(query, change)
		col.Update(bson.M{"timestamp": time}, bson.M{"$set": bson.M{"current": true}})
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.POST("/tally/:id", func(c *gin.Context) {
		var current Beer
		score := c.PostForm("score")
		time, _ := strconv.Atoi(c.Param("id"))
		query := bson.M{"timestamp": time}
		col.Find(query).One(&current)
		if score == "Up" {
			col.Update(query, bson.M{"$set": bson.M{"up": current.Up + 1}})
		} else {
			col.Update(query, bson.M{"$set": bson.M{"down": current.Down + 1}})
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("listening...")
	r.Run(":" + port)
}
