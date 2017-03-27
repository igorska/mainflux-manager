/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */
package api

/*
import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/mainflux/mainflux-core/db"
	"github.com/mainflux/mainflux-core/models"

	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"

	"io"
	"io/ioutil"
	"net/http"

	"github.com/go-zoo/bone"
)

/** == Functions == */

/*
// createDevice function
func createUser(w http.ResponseWriter, r *http.Request) {
	// Set up defaults and pick up new values from user-provided JSON
	d := models.User{}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	if len(data) > 0 {
		if err, str := validateDeviceSchema(data); err {
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

	// Creating UUID Version 4
	uuid := uuid.NewV4()
	d.ID = uuid.String()

	// Timestamp
	t := time.Now().UTC().Format(time.RFC3339)
	d.Created, d.Updated = t, t

	// Init MongoDB
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	// Insert Device
	if err := Db.C("devices").Insert(d); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		str := `{"response": "cannot create device"}`
		io.WriteString(w, str)
		return
	}

	// Publish on NATS
	hdr := r.Header.Get("Authorization")
	msg := `{"type": "device", "id":"` + d.ID + `", "owner": "` + hdr + `"}`
	NatsConn.Publish("core-auth", []byte(msg))

	// Send RSP
	w.Header().Set("Location", fmt.Sprintf("/devices/%s", d.ID))
	w.WriteHeader(http.StatusCreated)
}
*/
