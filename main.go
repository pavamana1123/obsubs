package main

import (
	"log"
	"net/http"
	"time"

	obsws "github.com/christopher-dG/go-obs-websocket"
)

var (
	obsClient obsws.Client
)

func init() {
	obsClient = obsws.Client{Host: "localhost", Port: 4444}
	if err := obsClient.Connect(); err != nil {
		log.Fatal("Unable to connect to OBS!")
	}
	obsws.SetReceiveTimeout(time.Second * 2)
}

func main() {

	defer func() {
		obsClient.Disconnect()
		log.Println("OBS Client disconnected!")
	}()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/obs/text", updateText)

	log.Println("Listening on :4000...")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
