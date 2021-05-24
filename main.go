package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	obsws "github.com/christopher-dG/go-obs-websocket"
)

var (
	obsClient obsws.Client
)

func init() {
	obsClient = obsws.Client{Host: "localhost", Port: 4444}
	err := fmt.Errorf("")
	t := ""

	for err != nil {
		err = obsClient.Connect()
		if err != nil {
			log.Println("Unable to connect to OBS! Press enter to try again")
			fmt.Scanln(&t)
		}

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

	log.Println("Opening UI...")
	exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:4445").Start()
	log.Println("Listening on :4445...")
	err := http.ListenAndServe(":4445", nil)
	if err != nil {
		log.Fatal(err)
	}

}
