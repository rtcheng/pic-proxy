package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"github.com/jinadam/pic-proxy/api"
)

var (
	port = flag.Int("port", -1, "specify a port")
)

func main() {
	flag.Parse()

	listener := gateway.ListenAndServe
	portStr := fmt.Sprintf(":%d", *port)

	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
		http.Handle("/", http.FileServer(http.Dir("./public")))
	}

	log.Fatal(listener(portStr, Router()))
}

func Router() *gin.Engine {
	r := gin.New()
	r.GET("/api/proxy/*url", api.Proxy)
	return r
}
