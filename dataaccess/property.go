package dataaccess

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/piyushpawar54/mongo-golang-rest/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PropertyController struct {
	session *mgo.Session
}

func NewPropertyController(s *mgo.Session) *PropertyController {

	return &PropertyController{s}
}

// GetProperties
func (up PropertyController) GetProperties(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	properties := []models.Property{}

	if err := up.session.DB("mongo-golang-rest").C("properties").Find(nil).All(&properties); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	pj, err := json.Marshal(properties)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", pj)
}

//GetPropertyById

func (up PropertyController) GetPropertyById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	p := models.Property{}

	if err := up.session.DB("mongo-golang-rest").C("properties").FindId(oid).One(&p); err != nil {
		w.WriteHeader(404)
		return
	}

	pj, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", pj)
}

// FilterProperties
func (up PropertyController) FilterProperties(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	queryParams := r.URL.Query()
	city := queryParams.Get("city")
	bedrooms, _ := strconv.Atoi(queryParams.Get("bedrooms"))

	filter := bson.M{}
	if city != "" {
		filter["city"] = city
	}
	if bedrooms > 0 {
		filter["bedrooms"] = bedrooms
	}

	properties := []models.Property{}

	if err := up.session.DB("mongo-golang-rest").C("properties").Find(filter).All(&properties); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	pj, err := json.Marshal(properties)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", pj)
}

// InsertProperty
func (up PropertyController) InsertProperty(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p := models.Property{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p.ID = bson.NewObjectId()

	if err := up.session.DB("mongo-golang-rest").C("properties").Insert(p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pj, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", pj)
}

// DeletePropety
func (up PropertyController) DeleteProperty(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := up.session.DB("mongo-golang-rest").C("properties").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
