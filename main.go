package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

type Player struct {
	ID    string
	Piece int
	Board [10][20]int
	Game  *Game
}

type Game struct {
	CurrentID string
	Players   *[]Player
}

type Games struct {
	Lock    sync.Mutex
	Games   *[]Game
	Players *[]Player
}

type Client struct {
	Conn *websocket.Conn
	ID   string
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/create-game", CreateGameHandler)
	
	server := &http.Server{
		Handler: router,
		Addr: "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Println("Running on :8080")
	log.Fatal(server.ListenAndServe())
}