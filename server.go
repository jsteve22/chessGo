package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jsteve22/chessGo/board"
)

type Response struct {
	Data []board.JsonMove
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	data := make([]board.JsonMove, 0)
	length := r.ContentLength
	fmt.Printf("body length: %v\n", length)

	if length > 0 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		respBytes := buf.String()
		fen := string(respBytes)
		fmt.Printf("body content: %s\n", fen)
		// data = strings.ToLower(data)
		game := board.LoadBoard(fen)
		moves := board.GenerateMoves(game)
		board.PrintMoves(game, moves)
		data = board.GetJsonMoves(moves)
	}
	fmt.Printf("\n")

	header := w.Header()
	header.Set("Access-Control-Allow-Origin", "*")

	response := Response{Data: data}
	marshalled, err := json.Marshal(response)
	// write response and send back to client
	_, err = w.Write(marshalled)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./chessfrontend/dist"))
	http.Handle("/", fs)
	http.HandleFunc("/FEN", homePageHandler)

	fmt.Printf("Starting server on port 12345\n")
	log.Panic(
		http.ListenAndServe(":12345", nil),
	)
}