package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func sendFile(filename string, w http.ResponseWriter) {
	f, err := os.ReadFile("./static/" + filename)
	if err != nil {
		http.Error(w, "Server failed to open file.", http.StatusInternalServerError)
		return
	}
	w.Write(f)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Println("client requested '" + r.URL.Path + "' html")
	if r.Method != "GET" {
		http.Error(w, "Request Method not supported", http.StatusBadRequest)
		return
	}
	if r.URL.Path == "/" {
		sendFile("index.html", w)
	} else if r.URL.Path == "/zawarudo" {
		sendFile("zawarudo.html", w)
	} else {
		sendFile("notFound.html", w)
		http.Error(w, "", http.StatusNotFound)
	}
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/css")
	w.WriteHeader(http.StatusOK)
	fmt.Println("client requested '" + r.URL.Path + "' css")
	if r.Method != "GET" {
		http.Error(w, "Request Method not supported", http.StatusBadRequest)
		return
	}
	if r.URL.Path == "/css/style.css" {
		sendFile("css/style.css", w)
	} else {
		sendFile("notFound.html", w)
		http.Error(w, "", http.StatusNotFound)
	}
}

func iconHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "image/x-icon")
	w.WriteHeader(http.StatusOK)
	fmt.Println("client requested '" + r.URL.Path + "' ico")
	if r.Method != "GET" {
		http.Error(w, "Request Method not supported", http.StatusBadRequest)
		return
	}
	if r.URL.Path == "/favicon.ico" {
		sendFile("favicon.ico", w)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/css/", cssHandler)
	mux.HandleFunc("/", htmlHandler)
	mux.HandleFunc("/favicon.ico", iconHandler)

	fmt.Printf("starting server at port 80\n")
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}
}
