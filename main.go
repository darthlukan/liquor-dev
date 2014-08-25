package main

import (
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"time"
)

var (
	projectRoot string
	templates   *template.Template
	goPath      = os.Getenv("GOPATH")
	messages    = web_responders.NewMessageMap()
)

func main() {
	log.Println("Starting server...")

	if goPath == "" {
		projectRoot = "."
	} else {
		projectRoot = path.Join(goPath, "src", "github.com", "darthlukan", "liquor-dev")
	}

	templates = template.Must(template.ParseGlob(projectRoot + "/html/*"))
	goweb.Map("/", indexHandler)
	goweb.Map("/ping", pingHandler)
	goweb.Map("/html/***", htmlFileHandler)
	goweb.MapStatic("/bower_components", path.Join(projectRoot, "bower_components"))
	goweb.MapStatic("/css", path.Join(projectRoot, "css"))
	goweb.MapStatic("/js", path.Join(projectRoot, "js"))
	goweb.MapStatic("/img", path.Join(projectRoot, "img"))

	address := ":3000"
	if port := os.Getenv("PORT"); port != "" {
		address = ":" + port
	}
	server := &http.Server{
		Addr:           address,
		Handler:        &LoggedHandler{goweb.DefaultHttpHandler()},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	listener, listenErr := net.Listen("tcp", address)
	if listenErr != nil {
		log.Panicf("Could not listen for TCP on %s: %s", address, listenErr)
	}
	log.Println("Server loaded, check localhost" + address)
	server.Serve(listener)
}

type LoggedHandler struct {
	baseHandler http.Handler
}

func (handler *LoggedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	go log.Printf("%s Request for %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
	handler.baseHandler.ServeHTTP(w, r)
}
