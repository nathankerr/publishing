package main

import (
	"path/filepath"
	"http"
	"io/ioutil"
	"log"
	"template"
)

const DRAFTS_DIR = "drafts"
const TEMPLATES_DIR = "templates"

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

func IndexServer(w http.ResponseWriter, req *http.Request) {
	templates := templateWithLayout("index.html")
	
	err := templates.Execute(w, "layout.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Loads the specified template from the TEMPLATES_DIR
// along with the layouts.
// Use the template by set.Execute(io.Writer, "layout.html", data)
func templateWithLayout(filename string) (*template.Set) {
	layout := filepath.Join(TEMPLATES_DIR, "layout.html")
	templates, err := template.ParseTemplateFiles(layout)
	if err != nil {
		log.Fatal(err)
	}

	contentFilename := filepath.Join(TEMPLATES_DIR, filename)
	content := template.New("Content")
	_, err = content.ParseFile(contentFilename)
	if err != nil {
		log.Fatal(err)
	}
	templates.Add(content)

	return templates
}

func main() {
	http.HandleFunc("/drafts", DraftsServer)
	http.HandleFunc("/pdf/", PDFServer)
	http.HandleFunc("/", IndexServer)

	log.Println("Starting Server")
	err := http.ListenAndServe("127.0.0.1:3000", nil)
	if err != nil {
		log.Fatal("Listen and Serve: ", err.String())
	}
}
