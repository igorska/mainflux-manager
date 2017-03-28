/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package api

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/go-zoo/bone"
)

// HTTPServer function
func HTTPServer() http.Handler {
	mux := bone.New()

	// Status
	mux.Get("/status", http.HandlerFunc(getStatus))

	// Users
	mux.Post("/users", http.HandlerFunc(createUser))
	mux.Get("/users", http.HandlerFunc(getUsers))

	mux.Get("/users/:user_id", http.HandlerFunc(getUser))
	mux.Put("/users/:user_id", http.HandlerFunc(updateUser))

	mux.Delete("/users/:user_id", http.HandlerFunc(deleteUser))

	// Apps
	//mux.Post("/apps", http.HandlerFunc(createApp))
	//mux.Get("/apps", http.HandlerFunc(getApps))

	//mux.Get("/apps/:app_id", http.HandlerFunc(getApp))
	//mux.Put("/apps/:app_id", http.HandlerFunc(updateApp))

	//mux.Delete("/apps/:app_id", http.HandlerFunc(deleteApp))

	n := negroni.Classic()
	n.UseHandler(mux)
	return n
}
