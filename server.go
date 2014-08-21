package main

import (
	"fmt"
	"strings"
	"net/http"
	"log"
)

var (
	config    = readConfig()
)

func main() {
	http.HandleFunc("/router_ip", serveRouterIp)
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	fmt.Printf("Server listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}

func serveRouterIp(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Request to router_ip for %s\n", req.RemoteAddr)

	// allow cross domain AJAX requests
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json")

	var remoteIp string
	if strings.Contains(req.RemoteAddr, ":") {
		remoteIp = strings.Split(req.RemoteAddr, ":")[0]
	} else {
		remoteIp = req.RemoteAddr
	}

	json, err := JsonMap{"router_ip": remoteIp }.String(); if err != nil{
		fmt.Println(err.Error)
	}

	fmt.Fprint(res, json)
}
