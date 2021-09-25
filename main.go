package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/chan157/nomadcoin/blockchain"
)

const port string = ":4000"

type HomeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := HomeData{"HOME", blockchain.GetBlockChain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
