package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/sotomskir/mastermind-server/models"
	"github.com/sotomskir/mastermind-server/utils"
	"net/http"
	"strconv"
)

var GetSshKeys = func(w http.ResponseWriter, r *http.Request) {
	keys := models.GetSshKeys()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(keys)
}

var GetSshKey = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 16)
	key := models.GetSshKey(id)
	if key == nil {
		utils.Error(w, "Cannot load key", errors.New("not found"), http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(key)
}

var SaveSshKeys = func(w http.ResponseWriter, r *http.Request) {
	var key models.SshKey
	err := json.NewDecoder(r.Body).Decode(&key)
	if nil != err {
		utils.Error(w, "Cannot decode key", err, http.StatusInternalServerError)
		return
	}
	if err := models.SaveSshKey(&key); nil != err {
		utils.Error(w, "Cannot save key", err, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(key)
}

var DeleteSshKeys = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 16)
	key := models.GetSshKey(id)
	if key == nil {
		utils.Error(w, "Cannot load key", errors.New("not found"), http.StatusNotFound)
		return
	}
	err := models.DeleteSshKey(key)
	if err != nil {
		utils.Error(w, "Cannot delete key", err, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
}
