package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/wcisco17/youtube-api-go/websocket"
)
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}

func stats(w http.ResponseWriter, r *http.Request) {
    ws, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }
    go websocket.Writer(ws)
}

func setupRoutes() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/stats", stats)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    fmt.Println("YouTube Subscriber Monitor")
    setupRoutes()
}