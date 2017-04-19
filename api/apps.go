/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */
package api

import (
	"fmt"
	"time"

	"github.com/mainflux/mainflux-manager/db"
	"github.com/mainflux/mainflux-manager/models"

	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"

	"encoding/json"
	"log"
	"net/http"

	"github.com/go-zoo/bone"

	"gopkg.in/mgo.v2/bson"
)

/** == Functions == */

// createApp function
func createApp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	// Set up defaults and pick up new values from app-provided JSON
	a := models.App{}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	a.Name = "Default name"

	// Creating UUID Version 4
	uuid := uuid.NewV4()
	a.ID = uuid.String()

	// Timestamp
	t := time.Now().UTC().Format(time.RFC3339)
	a.Created, a.Updated = t, t

	// Init MongoDB
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	// Insert app
	if err := Db.C("apps").Insert(a); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		str := `{"response": "cannot create app"}`
		io.WriteString(w, str)
		return
	}

	// Send RSP
	w.Header().Set("Location", fmt.Sprintf("/apps/%s", a.ID))
	w.WriteHeader(http.StatusCreated)
}

func getApps(w http.ResponseWriter, r *http.Request) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	results := []models.App{}
	if err := Db.C("apps").Find(nil).All(&results); err != nil {
		log.Print(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	res, err := json.Marshal(results)
	if err != nil {
		log.Print(err)
	}
	io.WriteString(w, string(res))
}

func getApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	id := bone.GetValue(r, "app_id")

	result := models.App{}
	err := Db.C("apps").Find(bson.M{"id": id}).One(&result)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		str := `{"response": "not found", "id": "` + id + `"}`
		io.WriteString(w, str)
		return
	}

	w.WriteHeader(http.StatusOK)
	res, err := json.Marshal(result)
	if err != nil {
		log.Print(err)
	}
	io.WriteString(w, string(res))
}

// updateApp function
func updateApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	if len(data) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		str := `{"response": "no data provided"}`
		io.WriteString(w, str)
		return
	}

	var body map[string]interface{}
	if err := json.Unmarshal(data, &body); err != nil {
		panic(err)
	}

	// Timestamp
	body["updated"] = time.Now().UTC().Format(time.RFC3339)

	// Init MongoDB
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	// App id
	id := bone.GetValue(r, "app_id")

	colQuerier := bson.M{"id": id}
	change := bson.M{"$set": body}
	if err := Db.C("apps").Update(colQuerier, change); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		str := `{"response": "not updated", "id": "` + id + `"}`
		io.WriteString(w, str)
		return
	}

	w.WriteHeader(http.StatusOK)
	str := `{"response": "updated", "id": "` + id + `"}`
	io.WriteString(w, str)
}

// deleteApp function
func deleteApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	id := bone.GetValue(r, "app_id")

	// Delete app
	if err := Db.C("apps").Remove(bson.M{"id": id}); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		str := `{"response": "not deleted", "id": "` + id + `"}`
		io.WriteString(w, str)
		return
	}

	w.WriteHeader(http.StatusOK)
	str := `{"response": "deleted", "id": "` + id + `"}`
	io.WriteString(w, str)
}
