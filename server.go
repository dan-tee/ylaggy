package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	indexHtml []byte
	favicon   []byte
	config    = readConfig()
)

func main() {
	cacheIndexHtml()
	cacheFavicon()

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/favicon.ico", serveFavicon)
	http.HandleFunc("/router_ip", serveRouterIp)

	fmt.Printf("Server listening on port %d\n", config.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	if err != nil {
		panic(err)
	}
}

func cacheIndexHtml() {
	indexHtml, _ = ioutil.ReadFile("index.html")
}

func cacheFavicon() {
	favicon, _ = ioutil.ReadFile("traffic_light.png")
}

func serveIndex(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Request to Index")
	res.Write(indexHtml)
}

func serveFavicon(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Request to favicon.ico")
	res.Write(favicon)
}

func serveRouterIp(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Request to router_ip for %s\n", req.RemoteAddr)

	// allow cross domain AJAX requests
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json")

	json, err := JsonMap{"router_ip": req.RemoteAddr}.String(); if err != nil{
		fmt.Println(err.Error)
	}

	fmt.Fprint(res, json)
}
