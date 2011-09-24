package main

import (
	"path/filepath"
	"http"
	"io"
	"log"
	"strings"
	"url"
)

func PDFServer(w http.ResponseWriter, req *http.Request) {
	filename := draftFilename(req)
	io.WriteString(w, filename)

	

	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//w.Write(contents)
}

func draftFilename(req *http.Request) string {
	rawPath := strings.SplitN(req.URL.String(), "/", 3)
	rawFilename, err := url.QueryUnescape(rawPath[2])
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(DRAFTS_DIR, rawFilename)
}
