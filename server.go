package main

import (
	"encoding/json"
	"fmt"
	"github.com/aeden/traceroute"
	"io/ioutil"
	"net/http"
	"os"
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
	http.HandleFunc("/traceroute", serveTraceRoute)

	fmt.Printf("Server listening on port %d\n", os.Getenv("PORT"))
	err := http.ListenAndServe(fmt.Sprintf(":%d", os.Getenv("PORT")), nil)
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

func serveTraceRoute(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Request to traceroute for %s\n", req.RemoteAddr)
	options := new(traceroute.TracerouteOptions)
	options.SetTimeoutMs(config.Timeout)
	options.SetRetries(1)

	traceRes, err := traceroute.Traceroute(req.RemoteAddr, options)

	if err != nil {
		panic(err)
	}
	json, err := json.Marshal(traceRes)
	if err != nil {
		panic(err)
	}
	fmt.Println(traceRes.DestinationAddress)
	fmt.Printf("%s\n", json)
	res.Write(json)
}
