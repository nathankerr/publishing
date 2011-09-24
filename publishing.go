package main

import (
	"http"
	"io"
	"io/ioutil"
	"log"
)

const DRAFTS_DIR = "drafts"

func DraftsServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Drafts\n")

	draftsInfo, err := ioutil.ReadDir(DRAFTS_DIR)
	if err != nil {
		io.WriteString(w, err.String())
	}

	for _, v := range draftsInfo {
		io.WriteString(w, v.Name)
	}
}

func main() {
	http.HandleFunc("/drafts", DraftsServer)

	log.Println("Starting Server")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Listen and Serve: ", err.String())
	}
}
