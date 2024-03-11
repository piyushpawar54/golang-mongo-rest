package dataaccess

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (up PropertyController) GetPropertyById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	p := models.Property{}

	if err := up.session.DB("mondo-golang-rest").C("properties").FindId(oid).One(&p); err != nil {
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
