package main

import (
	"http"
	//"io"
	"io/ioutil"
	"log"
	"template"
)

const DRAFTS_DIR = "drafts"

func DraftsServer(w http.ResponseWriter, req *http.Request) {
	draftsInfo, err := ioutil.ReadDir(DRAFTS_DIR)
	if err != nil {
		log.Fatal(err)
	}

	template, err := template.ParseFile("templates/drafts.html")
	if err != nil {
		log.Fatal(err)
	}

	err = template.Execute(w, draftsInfo)
	if err != nil {
		log.Fatal(err)
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
