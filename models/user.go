/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package models

type (
	User struct {
		ID   string `json:"id"`
		Name string `json:"name"`

		Created string `json:"created"`
		Updated string `json:"updated"`
	}
)
