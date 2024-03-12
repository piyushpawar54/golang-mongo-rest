package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/piyushpawar54/mongo-golang-rest/dataaccess"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	up := dataaccess.NewPropertyController(getSession())
	//Routes
	r.GET("/Properties", up.GetProperties)
	r.GET("/Properties/:id", up.GetPropertyById)
	r.GET("/properties/CityWise", up.FilterProperties)
	r.POST("/Properties", up.InsertProperty)
	r.DELETE("/Properties/:id", up.DeleteProperty)
}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost:27017/?directConnection=true")
	if err != nil {
		println("Error is here: ")
		panic(err)
	}
	return s
}
