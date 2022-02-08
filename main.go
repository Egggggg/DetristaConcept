package main

import (
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/egggggg/detrista-concept/utils"
	"github.com/gorilla/mux"
	"nhooyr.io/websocket"
)

type Player struct {
	Board [10][20]int
	Conn  *websocket.Conn
	Game  *Game
	ID    string
	Piece int
}

type Game struct {
	Current  int
	Pieces   []int
	PieceMap []string
	Players  []*Player
	Slug     string
}

type Hub struct {
	Count     int
	Games     map[string]*Game
	GamesMu   sync.Mutex
	Players   map[string]*Player
	PlayersMu sync.Mutex
}

func (game *Game) AddPlayer(player *Player) {
	game.Players = append(game.Players, player)
	player.Game = game
}

func main() {
	router := mux.NewRouter()
	hub := Hub{
		Games: make(map[string]*Game),
		Players: make(map[string]*Player),
	}

	rand.Seed(time.Now().UnixMicro())

	router.HandleFunc("*", ServeIndex)
	router.PathPrefix("/static/").
		Handler(
			http.StripPrefix("/static/", 
				http.FileServer(http.Dir("./static/detrista/build/static")),
			),
		)
	router.HandleFunc("/create", CreateGame(&hub))
	router.HandleFunc("/api/games/{slug}", HookGame(&hub))

	server := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Running on :8080")
	log.Fatal(server.ListenAndServe())
}

func ServeIndex(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./static/detrista/build/index.html")
}

func CreateGame(hub *Hub) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		slug := utils.RandomString(10)
		counter := 0

		hub.GamesMu.Lock()
		defer hub.GamesMu.Unlock()

		for hub.Games[slug] != nil {
			if counter > 10 {
				http.Error(res, "failed to create game, please try again", http.StatusInternalServerError)
				return
			}

			slug = utils.RandomString(10)
			counter++
		}

		game := &Game{
			PieceMap: []string{"i", "l", "o", "s", "t", "j", "z"},
			Players: make([]*Player, 0),
			Slug:     slug,
		}

		hub.Games[slug] = game

		http.Redirect(res, req, "/play/"+slug, http.StatusTemporaryRedirect)
	}
}

func HookGame(hub *Hub) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		slug := mux.Vars(req)["slug"]
		if slug == "" {
			http.Error(res, "page not found", http.StatusNotFound)
			return
		}

		hub.GamesMu.Lock()
		defer hub.GamesMu.Unlock()

		game := hub.Games[slug]
		if game == nil {
			http.Error(res, "page not found", http.StatusNotFound)
			return
		}

		id := utils.RandomString(20)

		conn, err := websocket.Accept(res, req, nil)
		if err != nil {
			http.Error(res, "couldn't upgrade connection", http.StatusBadRequest)
			return
		}

		player := &Player{
			Conn: conn,
			ID: id,
		}

		game.AddPlayer(player)
		
		utils.WriteTimeout(req.Context(), time.Second*5, conn, []byte("connected"))
	}
}