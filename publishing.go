package main

import (
	"http"
	"io"
	"log"
)

func DraftsServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Drafts")
}

func main() {
	http.HandleFunc("/drafts", DraftsServer)

	log.Println("Starting Server")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Listern and Serve: ", err.String())
	}
}
