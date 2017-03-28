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

	"github.com/mainflux/mainflux-app-manager/db"
	"github.com/mainflux/mainflux-app-manager/models"

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

// createUser function
func createUser(w http.ResponseWriter, r *http.Request) {
	// Set up defaults and pick up new values from user-provided JSON
	u := models.User{}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	/*
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		if len(data) > 0 {
			if err, str := validateUserSchema(data); err {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, str)
				return
			}

			if err := json.Unmarshal(data, &d); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				str := `{"response": "cannot decode body"}`
				io.WriteString(w, str)
				return
			}
		}
	*/

	u.Name = "Default name"

	// Creating UUID Version 4
	uuid := uuid.NewV4()
	u.ID = uuid.String()

	// Timestamp
	t := time.Now().UTC().Format(time.RFC3339)
	u.Created, u.Updated = t, t

	// Init MongoDB
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	// Insert User
	if err := Db.C("users").Insert(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		str := `{"response": "cannot create user"}`
		io.WriteString(w, str)
		return
	}

	// Publish on NATS
	/**
	hdr := r.Header.Get("Authorization")
	msg := `{"type": "device", "id":"` + d.ID + `", "owner": "` + hdr + `"}`
	NatsConn.Publish("core-auth", []byte(msg))
	*/

	// Send RSP
	w.Header().Set("Location", fmt.Sprintf("/users/%s", u.ID))
	w.WriteHeader(http.StatusCreated)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	results := []models.User{}
	if err := Db.C("users").Find(nil).All(&results); err != nil {
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

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	id := bone.GetValue(r, "user_id")

	result := models.User{}
	err := Db.C("users").Find(bson.M{"id": id}).One(&result)
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

// updateUser function
func updateUser(w http.ResponseWriter, r *http.Request) {
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

	// if err, str := validateDeviceSchema(data); err {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	io.WriteString(w, str)
	// 	return
	// }

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

	// Device id
	id := bone.GetValue(r, "device_id")

	colQuerier := bson.M{"id": id}
	change := bson.M{"$set": body}
	if err := Db.C("devices").Update(colQuerier, change); err != nil {
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
