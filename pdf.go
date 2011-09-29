package main

import (
	"bytes"
	"exec"
	"path/filepath"
	"http"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"url"
)

const PANDOC_DIR = "pandoc"

func PDFServer(w http.ResponseWriter, req *http.Request) {
	filename := draftFilename(req)
	//iconv := exec.Command("iconv", "-f", "utf-16", "-t", "utf-8", filename)
	// try automatic reencoding
	// Writer for mac defaults to UTF-8 encoding, while
	// Writer for iPad only uses UTF-16
	iconv := exec.Command("iconv", "-t", "utf-8", filename)
	// I will make sure and force mac to UTF-16

	iconv_out, err := iconv.Output()
	if err != nil {
		log.Fatal("pdfserver iconv_out:", err)
	}

	tmpdir, err := ioutil.TempDir("tmp/", "")
	if err != nil {
		log.Fatal("pdfserver tmpdir:", err)
	}
	defer os.RemoveAll(tmpdir)

	template, err := filepath.Abs(PANDOC_DIR)
	if err != nil {
		log.Fatal("pdfserver template:", err)
	}
	template = filepath.Join(template, "template.tex")

	//pdf_out := filepath.Join(tmpdir, "output.pdf")
	//pandoc := exec.Command("markdown2pdf",
	//	"-o", pdf_out,
	//	"--xetex",
	//	"--template", template,
	//)

	tex_out := filepath.Join(tmpdir, "output.tex")
	pandoc := exec.Command("pandoc",
		"-o", tex_out,
		"--xetex",
		"--smart",
		"--template", template,
	)

	pandoc.Stdin = bytes.NewBuffer(iconv_out)
	pandoc.Stdout = os.Stdout
	pandoc.Stderr = os.Stderr

	if pandoc.Run() != nil {
		log.Fatal("pdfserver pandoc.Run:", err)
	}

	pdf_out := filepath.Join(tmpdir, "output.pdf")
	xelatex := exec.Command("xelatex",
		//"-o", pdf_out,
		"-output-directory=" + tmpdir,
		tex_out,
	)
	xelatex.Stdout = os.Stdout
	xelatex.Stderr = os.Stderr
	if xelatex.Run() != nil {
		log.Fatal("pdfserver xelatex.Run:", err)
	}

	content, err := ioutil.ReadFile(pdf_out)
	if err != nil {
		log.Fatal("pdfserver read content:", err)
	}
	w.Write(content)
}

func draftFilename(req *http.Request) string {
	rawPath := strings.SplitN(req.URL.String(), "/", 3)
	rawFilename, err := url.QueryUnescape(rawPath[2])
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(DRAFTS_DIR, rawFilename)
}
