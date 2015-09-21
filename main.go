package main

import (
	// "encoding/json"
	// "fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Beer is the model for mongo
type Beer struct {
	Name    string
	Brewery string
	Style   string
	Alcohol int
	current bool
}

//Query is used to find the current beer
type Query struct {
	current bool
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
		col.Find(nil).All(&beers)
		current := col.Find(Query{current: true})
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"beers":   beers,
			"current": current,
		})
	})

	r.POST("/addBeer", func(c *gin.Context) {
		name := c.PostForm("name")
		brewery := c.PostForm("brewery")
		style := c.PostForm("style")
		alcohol, _ := strconv.Atoi(c.PostForm("alcohol"))
		err = col.Insert(&Beer{Name: name, Brewery: brewery, Style: style, Alcohol: alcohol, current: true})
		if err != nil {
			panic(err)
		}
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.Run(":" + port)

}
