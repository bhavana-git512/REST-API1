package controllers

import (
	"encoding/json"
	"manger/model"
	"net/http"

	"github.com/gorilla/mux"
)

func Api(res http.ResponseWriter, req *http.Request) {

	ds := mux.Vars(req)

	copy := ds["name"]

	if copy != "" {

	}

	pradeep := model.Idlibatch{1, 25, "nice"}

	//db.Save(&pradeep)

	json.NewEncoder(res).Encode(pradeep)

}
