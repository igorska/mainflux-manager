/**
 * Copyright (c) 2017 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package main

import (
	//"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/mainflux/mainflux-app-manager/api"
	"github.com/mainflux/mainflux-app-manager/db"
	//"github.com/nats-io/go-nats"
	//"log"
	"net/http"
	"os"
)

const (
	help string = `
Usage: mainflux-influxdb [options]
Options:
	-a, --host	Host address
	-p, --port	Port
	-h, --help	Prints this message end exits`
)

type (
	Opts struct {
		HTTPHost      string
		HTTPPort      string
		MongoHost     string
		MongoPort     string
		MongoDatabase string
		Help          bool
	}
)

func main() {
	opts := Opts{}

	flag.StringVar(&opts.HTTPHost, "a", "0.0.0.0", "HTTP host.")
	flag.StringVar(&opts.HTTPPort, "p", "9090", "HTTP port.")
	flag.StringVar(&opts.MongoHost, "m", "0.0.0.0", "MongoDB host.")
	flag.StringVar(&opts.MongoPort, "q", "27017", "MongoDB port.")
	flag.StringVar(&opts.MongoDatabase, "d", "mainflux", "MongoDB database.")
	flag.BoolVar(&opts.Help, "h", false, "Show help.")
	flag.BoolVar(&opts.Help, "help", false, "Show help.")

	flag.Parse()

	if opts.Help {
		fmt.Printf("%s\n", help)
		os.Exit(0)
	}

	// MongoDb
	db.InitMongo(opts.MongoHost, opts.MongoPort, opts.MongoDatabase)

	// Print banner
	color.Cyan(banner)

	// Serve HTTP
	httpAddr := fmt.Sprintf("%s:%s", opts.HTTPHost, opts.HTTPPort)
	http.ListenAndServe(httpAddr, api.HTTPServer())
}

var banner = `
+-+-+-+-+-+-+-+-+ +-+-+-+ +-+-+-+-+-+-+-+
|M|a|i|n|f|l|u|x| |A|p|p| |M|a|n|a|g|e|r|
+-+-+-+-+-+-+-+-+ +-+-+-+ +-+-+-+-+-+-+-+

        == Industrial IoT System ==
       
       Made with <3 by Mainflux Team
[w] http://mainflux.io
[t] @mainflux

`
